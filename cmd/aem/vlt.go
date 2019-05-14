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
	return strWriter.String(), err
}

func (v *vlt) checkBin() bool {
	_, err := exec.LookPath(v.bin)
	return err == nil
}

func (v *vlt) getPaths(paramPath string, configPaths []string) []vltPath {
	if len(paramPath) > 0 {
		return []vltPath{v.pathDetail(paramPath)}
	}

	paths := []vltPath{}
	for _, path := range configPaths {
		paths = append(paths, v.pathDetail(path))
	}
	return paths
}

func (v *vlt) pathDetail(path string) vltPath {
	vp := vltPath{recursive: false, path: path}

	if vp.path[0:1] == "!" {
		vp.recursive = true
		vp.path = path[1:]
	}

	return vp
}

type vltPath struct {
	path      string
	recursive bool
}
