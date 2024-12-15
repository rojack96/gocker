package option

import "gocker/helpers"

const (
	dryRun = "--dry-run"
)

func DryRun() string {
	return helpers.Option(dryRun)
}
