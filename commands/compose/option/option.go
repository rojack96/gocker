package option

import (
	"github.com/rojack96/gocker/helpers"
)

const (
	build         = "--build"
	dryRun        = "--dry-run"
	detach        = "--detach"
	forceRecreate = "--force-recreate"
	noBuild       = "--no-build"
	noRecreate    = "--no-recreate"
	noColor       = "--no-color"
	noLogPrefix   = "--no-log-prefix"
	timestamps    = "--timestamps"
	quiet         = "--quiet"
	quietPull     = "--quiet-pull"
	removeOrphans = "--remove-orphans"
	volumes       = "--volumes"
)

func Build() string {
	return helpers.Option(build)
}

func DryRun() string {
	return helpers.Option(dryRun)
}

func ForceRecreate() string {
	return helpers.Option(forceRecreate)
}

func NoBuild() string {
	return helpers.Option(noBuild)
}

func NoRecreate() string {
	return helpers.Option(noRecreate)
}

func NoColor() string {
	return helpers.Option(noColor)
}

func NoLogPrefix() string {
	return helpers.Option(noLogPrefix)
}

func Timestamps() string {
	return helpers.Option(timestamps)
}

func Quiet() string {
	return helpers.Option(quiet)
}

func QuietPull() string {
	return helpers.Option(quietPull)
}

func RemoveOrphans() string {
	return helpers.Option(removeOrphans)
}

func Volumes() string {
	return helpers.Option(volumes)
}

func Detach() string {
	return helpers.Option(detach)
}
