package goexec

import (
	"errors"
	"io"
	"os/exec"
)

type options struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	dir    string
	env    []string
	args   []string
}

type Option func(*options)

func WithEnv(key, value string) Option {
	return func(opt *options) {
		if opt.env == nil {
			opt.env = make([]string, 0)
		}
		opt.env = append(opt.env, key+"="+value)
	}
}

func WithArg(arg string) Option {
	return func(o *options) {
		if o.args == nil {
			o.args = make([]string, 0)
		}
		o.args = append(o.args, arg)
	}
}

func WithArgs(args ...string) Option {
	return func(opt *options) {
		if opt.args == nil {
			opt.args = make([]string, 0)
		}
		opt.args = append(opt.args, args...)
	}
}

func WithStdin(stdin io.Reader) Option {
	return func(opt *options) {
		opt.stdin = stdin
	}
}

func WithStdout(stdout io.Writer) Option {
	return func(opt *options) {
		opt.stdout = stdout
	}
}

func WithStderr(stderr io.Writer) Option {
	return func(opt *options) {
		opt.stderr = stderr
	}
}

func loadOptions(opts ...Option) *options {
	opt := &options{}
	for i := range opts {
		opts[i](opt)
	}
	return opt
}

func with(cmd *exec.Cmd, opt *options) error {
	if cmd == nil {
		return errors.New("cmd is nil")
	}
	if opt == nil {
		return errors.New("opt is nil")
	}
	if len(opt.env) > 0 {
		cmd.Env = append(cmd.Env, opt.env...)
	}
	cmd.Args = append(cmd.Args, opt.args...)
	if opt.stdin != nil {
		cmd.Stdin = opt.stdin
	}
	if opt.stdout != nil {
		cmd.Stdout = opt.stdout
	}
	if opt.stderr != nil {
		cmd.Stderr = opt.stderr
	}
	if opt.dir != "" {
		cmd.Dir = opt.dir
	}
	return nil
}
