package main

import (
	"fmt"
	"runtime"
)

var (
	version   = "unknown"
	gitCommit = "unknown" // sha1 from git, output of $(git rev-parse HEAD)
	buildDate = "unknown" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

type info struct {
	Version   string
	GitCommit string
	BuildDate string
	GoVersion string
	Compiler  string
}

func NewInfo() *info {
	// These variables typically come from -ldflags settings to `go build`
	return &info{
		Version:   version,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
	}
}

func (i info) Print() string {
	return fmt.Sprintf("%+v\n", i)
}
