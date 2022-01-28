package benchmark

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/logica0419/remote-bmi/server/repository"
	shellword "github.com/mattn/go-shellwords"
)

type Benchmarker struct {
	repo    *repository.Repository
	command command
	sync.Mutex
}

type Config struct {
	Version string
}

func NewBenchmarker(c *Config, repo *repository.Repository) (*Benchmarker, error) {
	command, ok := commands[c.Version]
	if !ok {
		return nil, fmt.Errorf("unknown version: %s", c.Version)
	}

	return &Benchmarker{
		command: command,
		repo:    repo,
	}, nil
}

func (b *Benchmarker) Run(userID uuid.UUID, serverNumber int) (uuid.UUID, error) {
	b.Lock()
	defer b.Unlock()

	server, err := b.repo.SelectServerByUserIDAndServerNumber(userID, serverNumber)
	if err != nil {
		return uuid.Nil, err
	}

	commandStr := fmt.Sprintf(b.command.command, server.Address)
	args, err := shellword.Parse(commandStr)
	if err != nil {
		return uuid.Nil, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = b.command.workDir
	stdoutBinary, err := cmd.Output()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}
	stdout := string(stdoutBinary)
	log := repository.Log{
		ID:       id,
		UserID:   userID,
		ServerID: server.ID,
		StdOut:   stdout,
	}

	err = b.repo.InsertLog(log)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
