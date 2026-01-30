package buildinfo

// These are set via -ldflags at build/release time.
// Defaults are fine for local dev.
var (
	Version = "0.1.0"
	Commit  = "dev"
	Date    = "dev"
)
