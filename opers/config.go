package opers

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Module struct {
	SFF8024Identifier                    int    `yaml:"SFF8024Identifier"`
	CmisRevision                         int    `yaml:"CmisRevision"`
	ModuleRevision                       int    `yaml:"ModuleRevision"`
	MediaType                            int    `yaml:"MediaType"`
	VendorName                           string `yaml:"VendorName"`
	DateCode                             string `yaml:"DateCode"`
	MaxPower                             int    `yaml:"MaxPower"`
	LenghtSMF                            int    `yaml:"LenghtSMF"`
	ModuleTempMax                        int    `yaml:"ModuleTempMax"`
	ModuleTempMin                        int    `yaml:"ModuleTempMin"`
	TempMonHighWarningThreshold          int    `yaml:"TempMonHighWarningThreshold"`
	TempMonLowWarningThreshold           int    `yaml:"TempMonLowWarningThreshold"`
	VccMonHighWarningThreshold           int    `yaml:"VccMonHighWarningThreshold"`
	VccMonLowWarningThreshold            int    `yaml:"VccMonLowWarningThreshold"`
	OpticalPowerTxHighWarningThreshold   int    `yaml:"OpticalPowerTxHighWarningThreshold"`
	OpticalPowerTxLowWarningThreshold    int    `yaml:"OpticalPowerTxLowWarningThreshold"`
	LaserBiasCurrentHighWarningThreshold int    `yaml:"LaserBiasCurrentHighWarningThreshold"`
	LaserBiasCurrentLowWarningThreshold  int    `yaml:"LaserBiasCurrentLowWarningThreshold"`
	OpticalPowerRxHighWarningThreshold   int    `yaml:"OpticalPowerRxHighWarningThreshold"`
	OpticalPowerRxLowWarningThreshold    int    `yaml:"OpticalPowerRxLowWarningThreshold"`
	ProgOutputPowerMin                   int    `yaml:"ProgOutputPowerMin"`
	ProgOutputPowerMax                   int    `yaml:"ProgOutputPowerMax"`
	GridSpacingTxx                       int    `yaml:"GridSpacingTxx"`
	CurrentLaserFrequencyTxx             int    `yaml:"CurrentLaserFrequencyTxx"`
	TargetOutputPowerTxx                 int    `yaml:"TargetOutputPowerTxx"`
}

type ModulesConfig struct {
	Modules []Module `yaml:"Modules"`
}

type Step struct {
	Endval   int `yaml:"endval"`
	Duration int `yaml:"duration"`
}

type ScenarioModule struct {
	Voltage     []Step `yaml:"Voltage"`
	Temperature []Step `yaml:"Temperature"`
	TxPower     []Step `yaml:"TxPower"`
	RxPower     []Step `yaml:"RxPower"`
	Osnr        []Step `yaml:"Osnr"`
}

type ScenarioConfig struct {
	Duration        int              `yaml:"ScenarioDuration"`
	ScenarioModules []ScenarioModule `yaml:"ScenarioModules"`
}

func GetModulesConfig(filename string) ModulesConfig {

	modulesConfig, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error while opening file %s\n", filename)
		panic(1)
	}

	modulesYaml := ModulesConfig{}
	err = yaml.Unmarshal(modulesConfig, &modulesYaml)

	if err != nil {
		fmt.Printf("Error while parsing file %s\n", filename)
		panic(1)
	}

	return modulesYaml
}

func GetScenarioConfig(filename string) ScenarioConfig {

	scenarioConfig, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error while opening file %s\n", filename)
		panic(1)
	}

	scenarioYaml := ScenarioConfig{}
	err = yaml.Unmarshal(scenarioConfig, &scenarioYaml)

	if err != nil {
		fmt.Printf("Error while parsing file %s\n", filename)
		panic(1)
	}

	return scenarioYaml
}
