package main

import (
	"log"

	app "nostos/app"
)

// GitBranch is set by the CI build process to the name of the branch
//nolint:gochecknoglobals // This is filled in by the build system
var GitBranch = "local"

// GitCommit is set by the CI build process to the commit hash
//nolint:gochecknoglobals // This is filled in by the build system
var GitCommit = "build"

func main() {
	log.SetFlags(log.Lshortfile)

	instance := app.Create(GitBranch, GitCommit)

	if err := instance.Run(); err != nil {
		return
	}
}



































