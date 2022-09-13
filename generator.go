package main

import (
	"flag"
	"fmt"
	"os"
	"pi-wegrzyn/eeprom-generator/opers"
)

const version string = "1.0"

func main() {

	var scenarioFilename = flag.String("scenario", "scenario.yaml", "Location of scenario yaml config")
	var modulesFilename = flag.String("modules", "modules.yaml", "Location of modules yaml config")
	var outputPath = flag.String("outputs", ".", "Output of EEPROM files")
	var info = flag.Bool("version", false, "Print version")

	flag.Parse()

	if *info {
		fmt.Printf("Current version: %s\n", version)
		os.Exit(0)
	}

	scenarioConfig := opers.ScenarioConfig{}
	modulesConfig := opers.ModulesConfig{}

	opers.GetConfig(*scenarioFilename, &scenarioConfig)
	opers.GetConfig(*modulesFilename, &modulesConfig)

	opers.EepromToFiles(*outputPath, "module1", opers.CreateTimelapse(modulesConfig.Modules[0], scenarioConfig.ScenarioModules[0], scenarioConfig.Duration))
	opers.EepromToFiles(*outputPath, "module2", opers.CreateTimelapse(modulesConfig.Modules[1], scenarioConfig.ScenarioModules[1], scenarioConfig.Duration))

}
