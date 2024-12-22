package option

import (
	"github.com/rojack96/gocker/helpers"
)

const (
	all           = "--all"
	build         = "--build"
	dryRun        = "--dry-run"
	detach        = "--detach"
	forceRecreate = "--force-recreate"
	noBuild       = "--no-build"
	noRecreate    = "--no-recreate"
	noColor       = "--no-color"
	noLogPrefix   = "--no-log-prefix"
	noDeps        = "--no-deps"
	noTty         = "--no-tty"
	noTrunc       = "--no-trunc"
	timestamps    = "--timestamps"
	quiet         = "--quiet"
	quietPull     = "--quiet-pull"
	removeOrphans = "--remove-orphans"
	volumes       = "--volumes"
	services      = "--services"
	includeDeps   = "--include-deps"
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

func Services() string {
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
