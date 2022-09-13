package main

import (
	"flag"
	"pi-wegrzyn/eeprom-generator/opers"
)

func main() {

	var scenarioFilename = flag.String("scenario", "scenario.yaml", "Location of scenario yaml config")
	var modulesFilename = flag.String("modules", "modules.yaml", "Location of modules yaml config")
	var outputPath = flag.String("outputs", ".", "Output of EEPROM files")

	flag.Parse()

	scenarioConfig := opers.GetScenarioConfig(*scenarioFilename)
	modulesConfig := opers.GetModulesConfig(*modulesFilename)

	opers.EepromToFiles(*outputPath, "module1", opers.CreateTimelapse(modulesConfig.Modules[0], scenarioConfig.ScenarioModules[0], scenarioConfig.Duration))
	opers.EepromToFiles(*outputPath, "module2", opers.CreateTimelapse(modulesConfig.Modules[1], scenarioConfig.ScenarioModules[1], scenarioConfig.Duration))

}
