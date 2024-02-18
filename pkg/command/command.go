package command

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	// Cmd is a combination of the program and provided args
	Cmd    string
	DryRun bool
}

func NewCommand(cmd string, dryRun bool) *Command {
	return &Command{
		Cmd:    cmd,
		DryRun: dryRun,
	}
}

func (c *Command) Execute() (string, error) {
	return c.ExecuteString()
}

func (c *Command) ExecuteString() (string, error) {
	var out strings.Builder
	if err := c.ExecuteStream(&out); err != nil {
		return "", err
	}
	return out.String(), nil
}

func (c *Command) ExecuteStream(writer io.Writer) error {
	// enable capability of using variables
	cmd := exec.Command(DefaultShell, DefaultShellArg, c.Cmd)
	if c.DryRun {
		fmt.Println(cmd.String())
		return nil
	}
	cmd.Stdout = writer
	cmd.Stderr = writer
	return cmd.Run()
}
