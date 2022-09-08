package main

import (
	"flag"
	"pi-wegrzyn/eeprom-generator/opers"
)

func main() {

	var scenarioFilename = flag.String("scenario", "scenario.yaml", "Location of scenario yaml config")
	var modulesFilename = flag.String("modules", "modules.yaml", "Location of modules yaml config")

	flag.Parse()

	scenarioConfig := opers.GetScenarioConfig(*scenarioFilename)
	modulesConfig := opers.GetModulesConfig(*modulesFilename)

	data := [][]byte{opers.CreateTimelapse(modulesConfig.Modules[0], scenarioConfig.ScenarioModules[0], scenarioConfig.Duration)}

	opers.EepromToFiles("./", "tmp", data)

}
