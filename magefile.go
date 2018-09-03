// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
)

func init() {
	os.Setenv("GO111MODULE", "on")
}

func Vet() error {
	fmt.Println("Vetting files...")

	return sh.Run("go", "vet", "./...")
}

func Lint() error {
	fmt.Println("Linting files...")

	return sh.Run("golint", "-set_exit_status", "./...")
}

func Fmt() error {
	fmt.Println("Formatting files...")

	return sh.Run("go", "fmt", "./...")
}

func Deps() error {
	var err error

	fmt.Println("Installing dependencies...")

	if err = sh.Run("dep", "ensure"); err != nil {
		return err
	}

	if err = sh.Run("go", "mod", "tidy"); err != nil {
		return err
	}

	return nil
}

func Test() error {
	fmt.Println("Testing...")

	return sh.Run("go", "test", "-v", "-short", "./...")
}

func TestRace() error {
	fmt.Println("Testing with race detector...")

	return sh.Run("go", "test", "-v", "-race", "./...")
}

func TestCoverage() error {
	fmt.Println("Testing with race detector and test coverage...")

	return sh.Run("go", "test", "-v", "-cover", "-race", "-coverprofile=coverage.out", "./...")
}
