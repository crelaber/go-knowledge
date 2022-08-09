package version

import (
	"fmt"
	"runtime"
)

var (
	buildVersion    = "unknown"
	buildGitVersion = "unknown"
	buildUser       = "unknown"
	buildHost       = "unknown"
	buildState      = "unknown"
	buildTime       = "unknown"
)

type BuildInfo struct {
	Version       string `json:"version"`
	GitRevision   string `json:"revision"`
	User          string `json:"user"`
	Host          string `json:"host"`
	GolangVersion string `json:"golang_version"`
	BuildState    string `json:"state"`
	BuildTime     string `json:"time"`
}

var (
	Info BuildInfo
)

func init() {
	Info = BuildInfo{
		Version:       buildVersion,
		GitRevision:   buildGitVersion,
		User:          buildUser,
		Host:          buildHost,
		GolangVersion: runtime.Version(),
		BuildState:    buildState,
		BuildTime:     buildTime,
	}
}

func (b BuildInfo) String() string {
	return fmt.Sprintf("%s@%s-%s-%s-%s-%s", b.User, b.Host, b.Version, b.GitRevision, b.BuildState, b.BuildTime)
}

func (b BuildInfo) LongForm() string {
	return fmt.Sprintf(`Version: %v
GitRevision: %v
User: %v@%v
GolangVersion: %v
BuildStatus: %v
BuildTime: %v
`,
		b.Version,
		b.GitRevision,
		b.User,
		b.Host,
		b.GolangVersion,
		b.BuildState,
		b.BuildTime)
}
