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

	data1 := []uint8{0x12, 0x99, 0x88}
	data2 := []uint8{0x13, 0x80, 0x99}

	data := [][]byte{data1, data2}

	opers.EepromToFiles("./", "tmp", data)

}
