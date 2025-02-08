package options

import (
	"github.com/rojack96/gocker/helpers"
)

const (
	all           = "--all"
	Attach        = "--attach"
	build         = "--build"
	detach        = "--detach"
	dryRun        = "--dry-run"
	forceRecreate = "--force-recreate"
	includeDeps   = "--include-deps"
	noBuild       = "--no-build"
	noColor       = "--no-color"
	noDeps        = "--no-deps"
	noLogPrefix   = "--no-log-prefix"
	noRecreate    = "--no-recreate"
	noTty         = "--no-tty"
	noTrunc       = "--no-trunc"
	quiet         = "--quiet"
	quietPull     = "--quiet-pull"
	removeOrphans = "--remove-orphans"
	services      = "--services"
	timestamps    = "--timestamps"
	volumes       = "--volumes"
)

func All() string {
	return helpers.Option(all)
}

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

func ServicesOption() string {
	return helpers.Option(services)
}

func IncludeDeps() string {
	return helpers.Option(includeDeps)
}

func NoDeps() string {
	return helpers.Option(noDeps)
}

func NoTty() string {
	return helpers.Option(noTty)
}

func NoTrunc() string {
	return helpers.Option(noTrunc)
}
