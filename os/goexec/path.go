package goexec

import "os/exec"

type LP struct {
	isExec bool
	name   string
	p      string
	e      error
}

func NewLookPath(name string) *LP {
	return &LP{name: name}
}

func (lp *LP) LookPath() {
	s, err := exec.LookPath(lp.name)
	lp.isExec = true
	lp.p = s
	lp.e = err
}

func (lp *LP) Path() string {
	return lp.p
}

func (lp *LP) Error() error {
	return lp.e
}

func (lp *LP) Exists() bool {
	if !lp.isExec {
		return false
	}
	return lp.Error() == nil
}
