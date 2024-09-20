package goexec

import (
	"context"
	"os/exec"
)

func initCmd(ctx context.Context, command string, opts ...Option) (*exec.Cmd, error) {
	opt := loadOptions(opts...)
	cmd := new(exec.Cmd)
	if ctx != nil {
		cmd = exec.CommandContext(ctx, command)
	} else {
		cmd = exec.Command(command)
	}
	err := with(cmd, opt)
	return cmd, err
}

// Run executes the given command and wait for it to complete.
func Run(command string, opts ...Option) error {
	return RunCtx(nil, command, opts...)
}

func RunCtx(ctx context.Context, command string, opts ...Option) error {
	cmd, err := initCmd(ctx, command, opts...)
	if err != nil {
		return err
	}
	return cmd.Run()
}

// RunCtxOutput executes the given command and wait for it to complete.
func RunCtxOutput(ctx context.Context, command string, opts ...Option) ([]byte, error) {
	cmd, err := initCmd(ctx, command, opts...)
	if err != nil {
		return nil, err
	}
	return cmd.Output()
}

// RunOutput executes the given command and wait for it to complete.
func RunOutput(command string, opts ...Option) ([]byte, error) {
	return RunCtxOutput(nil, command, opts...)
}

func RunCtxWait(ctx context.Context, command string, opts ...Option) error {
	cmd, err := initCmd(ctx, command, opts...)
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

func RunWait(command string, opts ...Option) error {
	return RunCtxWait(nil, command, opts...)
}
