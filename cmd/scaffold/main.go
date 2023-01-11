package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type application struct {
	ProjectName *string
	Location    *string
	Repository  *string
	Static      *bool
}

func (app *application) initOptions() {
	static_files := flag.Bool("s", false, "Will the project support static files")
	project_name := flag.String("n", "", "Project Name for the project")
	loc := flag.String("d", "", "Directory to create the project")
	repo := flag.String("r", "", "Repository that will store the project")

	app.Static = static_files
	app.ProjectName = project_name
	app.Location = loc
	app.Repository = repo
	flag.Parse()
}

func (app *application) generateProject(w io.Writer) {
	fmt.Fprintf(w, "Generating scaffold for project %s in %s\n", *app.ProjectName, *app.Location)
}

func (app *application) validate(w io.Writer) bool {
	var result bool = true
	repo := app.Repository
	if *repo == "" {
		fmt.Fprintln(w, "Project repository URL cannot be empty")
		result = false
	}
	dir := app.Location
	if *dir == "" {
		fmt.Fprintln(w, "Project path cannot be empty")
		result = false
	}
	name := app.ProjectName
	if *name == "" {
		fmt.Fprintln(w, "Project name cannot be empty")
		result = false
	}
	return result
}

func main() {
	app := application{}
	app.initOptions()
	if app.validate(os.Stdout) {
		app.generateProject(os.Stdout)
	}
}
