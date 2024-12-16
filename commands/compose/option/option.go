package option

import "github.com/rojack96/gocker/helpers"

const (
	dryRun = "--dry-run"
)

func DryRun() string {
	return helpers.Option(dryRun)
}
