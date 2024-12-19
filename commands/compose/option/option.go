package option

import "github.com/rojack96/gocker/helpers"

const (
	dryRun      = "--dry-run"
	noColor     = "--no-color"
	noLogPrefix = "--no-log-prefix"
	timestamps  = "--timestamps"
)

// DryRun - Execute command in dry run mode
func DryRun() string {
	return helpers.Option(dryRun)
}

// NoColor - Produce monochrome output
func NoColor() string {
	return helpers.Option(noColor)
}

// NoLogPrefix - Don't print prefix in logs
func NoLogPrefix() string {
	return helpers.Option(noLogPrefix)
}

// Timestamps - Show timestamps
func Timestamps() string {
	return helpers.Option(timestamps)
}
