package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var binName = "scaffold-gen"
var srcName = "./cmd/scaffold"

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot buildtool %s: %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(result)
}

func TestScaffoldCLI(t *testing.T) {

	expected_name := "Hello"
	expected_directory := "World"
	expected_repo := "dragons.com"
	//expected_static := true

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("TestSetupFlags", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-n", expected_name, "-d", expected_directory, "-r", expected_repo, "-s")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected_output := fmt.Sprintf("Generating scaffold for project %s in %s\n", expected_name, expected_directory)

		assert.Equal(t, expected_output, string(out))

	})

	t.Run("TestRepoValidations", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-n", expected_name, "-d", expected_directory)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected_output := fmt.Sprintln("Project repository URL cannot be empty")

		assert.Equal(t, expected_output, string(out))
	})
	t.Run("TestDirValidations", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-n", expected_name, "-r", expected_repo)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected_output := fmt.Sprintln("Project path cannot be empty")

		assert.Equal(t, expected_output, string(out))
	})
	t.Run("TesNameValidations", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-d", expected_directory, "-r", expected_repo)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected_output := fmt.Sprintln("Project name cannot be empty")

		assert.Equal(t, expected_output, string(out))
	})
	//os.Args = []string{"scaffold_gen", "-n", expected_name, "-d", expected_directory, "-r", expected_repo, "-s"}
	//main()
	//app := application{}
	//app.initOptions()
	// assert.Equal(t, expected_name, app.ProjectName)
	// assert.Equal(t, expected_directory, app.Location)
	// assert.Equal(t, expected_repo, app.Repository)
	// assert.Equal(t, expected_static, app.Static)

}
