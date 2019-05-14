package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
)

func newVlt() vlt {
	return vlt{bin: "vlt"}
}

type vlt struct {
	bin string
}

func (v *vlt) execute(params ...string) (string, error) {
	if !v.checkBin() {
		err := errors.New("could not find vlt executable")
		return "", err
	}

	strWriter := stringWriter{}
	p := newProjectStructure()
	cmd := exec.Command(v.bin, params...)
	cmd.Dir = p.getWorkDir()
	cmd.Stdout = io.MultiWriter(os.Stdout, &strWriter)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return "", err
}

func (v *vlt) checkBin() bool {
	_, err := exec.LookPath("vlt")
	return err == nil
}
