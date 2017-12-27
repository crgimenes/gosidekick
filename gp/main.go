package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/crgimenes/goconfig"
)

type config struct {
	Path       string `cfg:"p"`
	Submodules bool   `cfg:"sm"`
}

var cfg config

func folderExists(path string) (fs bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println(err)
		return
	}
	fs = true
	return
}

func execHelper(path, name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

func gitExec(path string) (err error) {
	err = execHelper(path, "git", "stash")
	if err != nil {
		return
	}
	err = execHelper(path, "git", "checkout", "master")
	if err != nil {
		return
	}
	err = execHelper(path, "git", "pull")
	if err != nil {
		return
	}
	return
}

func visit(path string, f os.FileInfo, perr error) error {
	if perr != nil {
		return perr
	}
	if !f.IsDir() {
		return nil
	}
	if f.Name() == "vendor" {
		return filepath.SkipDir
	}
	fs := folderExists(path + "/.git")
	if !fs {
		return nil
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	fmt.Println(path)

	err = gitExec(path)
	if err != nil {
		return err
	}

	if cfg.Submodules {
		return nil
	}
	return filepath.SkipDir
}

func run() (err error) {
	goconfig.PrefixEnv = `GIT_PULL`
	if err = goconfig.Parse(&cfg); err != nil {
		return
	}
	if cfg.Path == "" {
		lastPar := flag.NArg() - 1
		cfg.Path = flag.Arg(lastPar)
		if cfg.Path == "" {
			cfg.Path = "./"
		}
	}
	err = filepath.Walk(cfg.Path, visit)
	if err == io.EOF {
		err = nil
	}
	return
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
