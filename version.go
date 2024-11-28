package ronin

import "time"

// These should be set via go build -ldflags -X 'vvv'.
var (
	// Ver is the SemVer version of the binary.
	Ver = "v0.0.0"

	// GitCommit is commit hash of version control.
	GitCommit string

	// BuildTimeStr is the time at which the application was built, in RFC3339 format.
	BuildTimeStr = "1970-01-01T00:00:00Z"

	// GoVersion is go build version.
	GoVersion string
)

// BuildTime in RFC3339 format
func BuildTime() (t time.Time, err error) {
	return time.Parse(time.RFC3339, BuildTimeStr)
}
