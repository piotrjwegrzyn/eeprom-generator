package opers

func CreateTimelapse(module Module, scenario ScenarioModule, duration int) (timelapse []byte) {

	currentTemperature := scenario.Temperature[0].Endval
	currentVoltage := scenario.Voltage[0].Endval

	timelapse = append(timelapse, GeneratePageLow(module, float32(currentTemperature/10), uint16(currentVoltage))...)
	timelapse = append(timelapse, GeneratePage00h(module)...)
	timelapse = append(timelapse, GeneratePage01h()...)
	timelapse = append(timelapse, GeneratePage02h(module)...)
	timelapse = append(timelapse, GeneratePage04h()...)
	timelapse = append(timelapse, GeneratePage10h()...)
	timelapse = append(timelapse, GeneratePage11h()...)
	timelapse = append(timelapse, GeneratePage12h()...)
	timelapse = append(timelapse, GeneratePage25h(scenario.Osnr[0].Endval, currentTemperature/10)...)

	return
}
