package option

import "github.com/rojack96/gocker/helpers"

const (
	dryRun = "--dry-run"
)

// DryRun - Execute command in dry run mode
func DryRun() string {
	return helpers.Option(dryRun)
}
