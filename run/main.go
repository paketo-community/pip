package main

import (
	"os"

	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/cargo"
	"github.com/paketo-buildpacks/packit/chronos"
	"github.com/paketo-buildpacks/packit/draft"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/postal"
	"github.com/paketo-buildpacks/packit/scribe"
	"github.com/paketo-community/pip"
)

func main() {
	entries := draft.NewPlanner()
	dependencies := postal.NewService(cargo.NewTransport())
	planRefinery := pip.NewPlanRefinery()
	logs := scribe.NewEmitter(os.Stdout)
	installProcess := pip.NewPipInstallProcess(pexec.NewExecutable("python"))

	packit.Run(
		pip.Detect(),
		pip.Build(installProcess, entries, dependencies, planRefinery, logs, chronos.DefaultClock),
	)
}