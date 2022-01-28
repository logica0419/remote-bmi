package benchmark

import (
	"fmt"
	"os/exec"

	"github.com/gofrs/uuid"
	"github.com/logica0419/remote-bmi/server/repository"
	shellword "github.com/mattn/go-shellwords"
)

type Benchmarker struct {
	repo    *repository.Repository
	command string
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
	server, err := b.repo.SelectServerByUserIDAndServerNumber(userID, serverNumber)
	if err != nil {
		return uuid.Nil, err
	}

	command := fmt.Sprintf(b.command, server.Address)
	args, err := shellword.Parse(command)
	if err != nil {
		return uuid.Nil, err
	}

	stdoutBinary, err := exec.Command(args[0], args[1:]...).Output()
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
