package opers

import (
	"fmt"
)

func generateSteps(steps []Step, duration int) (listS []float64) {
	for i := 0; i < len(steps); i++ {
		if i == 0 {
			for j := 0; j < steps[i].Duration; j++ {
				listS = append(listS, steps[i].Endval)
			}
		} else {
			f := func(x float64) float64 {
				return (steps[i].Endval-steps[i-1].Endval)*(x+1)/float64(steps[i].Duration) + steps[i-1].Endval
			}
			for j := 0; j < steps[i].Duration; j++ {
				listS = append(listS, f(float64(j)))
			}
		}
	}

	if len(listS) != duration {
		fmt.Printf("Error: mismatch between ScenarioDuration (%d) and steps durations (%d) near %f endval\n", duration, len(listS), listS[len(listS)-1])
		panic(1)
	}

	return
}

func CreateTimelapse(module Module, scenario ScenarioModule, duration int) (timelapse [][]byte) {

	listVcc := generateSteps(scenario.Voltage, duration)
	listTemp := generateSteps(scenario.Temperature, duration)
	listTxPower := generateSteps(scenario.TxPower, duration)
	listRxPower := generateSteps(scenario.RxPower, duration)
	listOsnr := generateSteps(scenario.Osnr, duration)

	for i := 0; i < duration; i++ {
		timelapseStep := make([]byte, 0)
		timelapseStep = append(timelapseStep, GeneratePageLow(module, listTemp[i], listVcc[i])...)
		timelapseStep = append(timelapseStep, GeneratePage00h(module)...)
		timelapseStep = append(timelapseStep, GeneratePage01h(module)...)
		timelapseStep = append(timelapseStep, GeneratePage02h(module)...)
		timelapseStep = append(timelapseStep, GeneratePage04h(module)...)
		timelapseStep = append(timelapseStep, GeneratePage11h(module, listTxPower[i], listRxPower[i])...)
		timelapseStep = append(timelapseStep, GeneratePage12h(module)...)
		timelapseStep = append(timelapseStep, GeneratePage25h(listOsnr[i], listTemp[i])...)
		timelapse = append(timelapse, timelapseStep)
	}

	return
}
