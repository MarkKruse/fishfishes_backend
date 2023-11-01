//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

import "fishfishes_backend/build"

var VERSION = "0.0.1"
var MODULE = ""

func init() {

}

func Build() {
	mg.Deps(
		TestBuild,
	)
}

func Run() error {
	cmd := exec.Command("docker-compose", "up", "-d", "--build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func TestBuild() error {
	return build.Build(MODULE, "ff", VERSION, build.LINUX, build.WINDOWS, build.MAC)
}
