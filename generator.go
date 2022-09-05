package main

import (
	"flag"
	"fmt"
	"pi-wegrzyn/eeprom-generator/opers"
)

func main() {

	var scenarioFilename = flag.String("scenario", "scenario.yaml", "Location of scenario yaml config")
	var modulesFilename = flag.String("modules", "modules.yaml", "Location of modules yaml config")

	flag.Parse()

	fmt.Println(opers.GetScenarioConfig(*scenarioFilename))
	fmt.Println(opers.GetModulesConfig(*modulesFilename))
}
