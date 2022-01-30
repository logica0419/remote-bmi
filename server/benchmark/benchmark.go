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
	repo *repository.Repository
	cmd  command
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
		cmd:  command,
		repo: repo,
	}, nil
}

func (b *Benchmarker) Run(userID uuid.UUID, serverNumber int) (uuid.UUID, error) {
	b.Lock()
	defer b.Unlock()

	servers, err := b.repo.SelectServersByUserID(userID)
	if err != nil {
		return uuid.Nil, err
	}

	var target *repository.Server
	for _, server := range servers {
		if server.ServerNumber == serverNumber {
			target = server
		}
	}
	if target == nil {
		return uuid.Nil, fmt.Errorf("target server not found")
	}

	command, err := b.cmd.createCmd(servers, serverNumber)
	if err != nil {
		return uuid.Nil, err
	}
	args, err := shellword.Parse(command)
	if err != nil {
		return uuid.Nil, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = b.cmd.workDir
	stdoutBinary, err := cmd.Output()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}
	stdout := string(stdoutBinary)
	log := &repository.Log{
		ID:       id,
		UserID:   userID,
		ServerID: target.ID,
		StdOut:   stdout,
	}

	err = b.repo.InsertLog(log)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
