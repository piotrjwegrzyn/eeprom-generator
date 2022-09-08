package opers

func CreateTimelapse(module Module, scenario ScenarioModule, duration int) (timelapse []byte) {

	timelapse = append(timelapse, GeneratePageLow()...)
	timelapse = append(timelapse, GeneratePage00h()...)
	timelapse = append(timelapse, GeneratePage01h()...)
	timelapse = append(timelapse, GeneratePage02h()...)
	timelapse = append(timelapse, GeneratePage04h()...)
	timelapse = append(timelapse, GeneratePage10h()...)
	timelapse = append(timelapse, GeneratePage11h()...)
	timelapse = append(timelapse, GeneratePage12h()...)
	timelapse = append(timelapse, GeneratePage24h()...)
	timelapse = append(timelapse, GeneratePage25h()...)

	return
}
