package opers

import (
	"crypto/md5"
	"math/rand"
	"time"
)

func Checksum(data []byte) byte {
	checksum := md5.Sum(data)
	return checksum[len(checksum)-1]
}

const TempMonAlarmThreshold int = 10           // in Celsius degrees, added/substracted as High/Low Alarm
const VccMonAlarmThreshold int = 300           // in 0.1 mV, -||-
const OpticalTxRxAlarmThreshold float32 = 1.25 // multiplicable factor for Optics Alarms

func GeneratePageLow(module Module, temperature float32, vcc uint16) (page []byte) {
	page = append(page, byte(module.SFF8024Identifier)) // SFF8024Identifier
	page = append(page, byte(module.CmisRevision))      // CmisRevision
	page = append(page, byte(0x04))                     // MemoryModel + SteppedConfigOnly + MciMaxSpeed
	page = append(page, byte(0b0101))                   // ModuleState
	page = append(page, make([]byte, 5)...)             // FlagsSummary (Banks and others)
	var flagsIndicator byte = 0
	{
		if vcc < uint16(module.VccMonLowWarningThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 7)
		}
		if vcc > uint16(module.VccMonHighWarningThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 6)
		}
		if vcc < uint16(module.VccMonLowWarningThreshold-VccMonAlarmThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 5)
		}
		if vcc > uint16(module.VccMonHighWarningThreshold+VccMonAlarmThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 4)
		}
		if temperature < float32(module.TempMonLowWarningThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 3)
		}
		if temperature > float32(module.TempMonHighWarningThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 2)
		}
		if temperature < float32(module.TempMonLowWarningThreshold-TempMonAlarmThreshold) {
			flagsIndicator = flagsIndicator | (0x1 << 1)
		}
		if temperature > float32(module.TempMonHighWarningThreshold+TempMonAlarmThreshold) {
			flagsIndicator = flagsIndicator | 0x1
		}
	}
	page = append(page, flagsIndicator)     // Latched Flags
	page = append(page, make([]byte, 4)...) // Aux and Custom Flags
	TempMonValue := int16(temperature * 256)
	page = append(page, byte(TempMonValue>>8), byte(TempMonValue&0xFF))                   // TempMonValue
	page = append(page, byte(vcc>>8), byte(vcc&0xFF))                                     // VccMonVoltage
	page = append(page, make([]byte, 14)...)                                              // Aux + Custom + Global Controls
	page = append(page, byte(0xFF))                                                       // Module Level Masks (Vcc + Temp)
	page = append(page, make([]byte, 6)...)                                               // -||- (Aux + Custom) + CDB status
	page = append(page, byte(module.ModuleRevision>>8), byte(module.ModuleRevision&0xFF)) // Module Active Firmware Version
	page = append(page, make([]byte, 44)...)                                              // Fault Information + Reserved + Custom
	page = append(page, byte(module.MediaType))                                           // MediaType
	for i := 0; i < 8; i++ {
		page = append(page, 0xFF, 0x00, 0x00, 0x00) // AppDescriptors
	}
	page = append(page, make([]byte, 10)...) // Password Facilities + Page Mapping

	return
}

func GeneratePage00h(module Module) (page []byte) {
	page = append(page, byte(module.SFF8024Identifier))
	vendorName := (append(make([]byte, 0, 16), module.VendorName...))[0:16]
	page = append(page, vendorName...)                                                // VendorName
	page = append(page, []byte{0xCC, 0xFA, 0xCE}...)                                  // VendorOUI
	page = append(page, append(vendorName[0:4], []byte("xx1234567890")...)...)        // VendorPN
	page = append(page, []byte{0x01, 0x23}...)                                        // VendorRev
	page = append(page, append(vendorName[0:4], []byte("xx1234567890")...)...)        // VendorSN
	page = append(page, append([]byte(module.DateCode)[2:8], []byte{0x0, 0x0}...)...) // DateCode
	page = append(page, []byte("BEST_MEMES")...)                                      // CLEI
	page = append(page, []byte{0b11100000, byte(module.MaxPower)}...)                 // ModulePowerCharacteristics
	page = append(page, []byte{0x00, 0x07}...)                                        // CableAssemblyLinkLength + ConnectorType
	page = append(page, make([]byte, 6)...)                                           // Copper Cable Attenuation
	page = append(page, []byte{0xfe, 0x00, 0x10}...)                                  // MediaLaneInformation + Cable Assembly Information + MediaInterfaceTechnology
	page = append(page, make([]byte, 9)...)                                           // Reserved+Custom
	page = append(page, Checksum(page[0:93]))                                         // PageChecksum
	page = append(page, make([]byte, 33)...)                                          // Custom

	return
}

func GeneratePage01h() (page []byte) {
	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
	return
}

func GeneratePage02h(module Module) (page []byte) {
	// Module-Level Monitor Thresholds (Temp)
	tempTemps := []int16{int16((module.TempMonHighWarningThreshold + TempMonAlarmThreshold) * 256),
		int16((module.TempMonLowWarningThreshold - TempMonAlarmThreshold) * 256),
		int16(module.TempMonHighWarningThreshold * 256),
		int16(module.TempMonLowWarningThreshold * 256),
	}
	for _, v := range tempTemps {
		page = append(page, byte(v>>8), byte(v&0xFF))
	}

	// Module-Level Monitor Thresholds (Vcc)
	tempVccs := []uint16{uint16(module.VccMonHighWarningThreshold + VccMonAlarmThreshold),
		uint16(module.VccMonLowWarningThreshold - VccMonAlarmThreshold),
		uint16(module.VccMonHighWarningThreshold),
		uint16(module.VccMonLowWarningThreshold),
	}
	for _, v := range tempVccs {
		page = append(page, byte(v>>8), byte(v&0xFF))
	}

	page = append(page, make([]byte, 32)...) // Aux + Custom

	// Module-Level Monitor Thresholds (OpticalPowerTx)
	tempOpticalTxs := []uint16{uint16(float32(module.OpticalPowerTxHighWarningThreshold) * OpticalTxRxAlarmThreshold),
		uint16(float32(module.OpticalPowerTxLowWarningThreshold) / OpticalTxRxAlarmThreshold),
		uint16(module.OpticalPowerTxHighWarningThreshold),
		uint16(module.OpticalPowerTxLowWarningThreshold),
	}
	for _, v := range tempOpticalTxs {
		page = append(page, byte(v>>8), byte(v&0xFF))
	}

	page = append(page, make([]byte, 8)...) // LaserBiasCurrent

	// Module-Level Monitor Thresholds (OpticalPowerRx)
	tempOpticalRxs := []uint16{uint16(float32(module.OpticalPowerRxHighWarningThreshold) * OpticalTxRxAlarmThreshold),
		uint16(float32(module.OpticalPowerRxLowWarningThreshold) / OpticalTxRxAlarmThreshold),
		uint16(module.OpticalPowerRxHighWarningThreshold),
		uint16(module.OpticalPowerRxLowWarningThreshold),
	}
	for _, v := range tempOpticalRxs {
		page = append(page, byte(v>>8), byte(v&0xFF))
	}

	page = append(page, make([]byte, 55)...)   // Reserved + Custom
	page = append(page, Checksum(page[0:126])) // Page Checksum
	return
}

func GeneratePage04h(module Module) (page []byte) {
	page = append(page, byte(1<<5))
	page = append(page, make([]byte, 21)...)                                                      // Unsupported Grids
	page = append(page, 0xFF, 0xEE, 0x00, 0x1E)                                                   // GridChannel100GHz
	page = append(page, make([]byte, 44)...)                                                      // Unsupported Grids + FineTuning
	page = append(page, byte(module.ProgOutputPowerMin>>8), byte(module.ProgOutputPowerMin&0xFF)) // ProgOutputPowerMin
	page = append(page, byte(module.ProgOutputPowerMax>>8), byte(module.ProgOutputPowerMax&0xFF)) // ProgOutputPowerMax
	page = append(page, make([]byte, 53)...)                                                      // Reserved
	page = append(page, Checksum(page[0:126]))                                                    // Page Checksum
	return
}

func GeneratePage10h() (page []byte) {
	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
	return
}

func GeneratePage11h() (page []byte) {
	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
	return
}

func GeneratePage12h() (page []byte) {

	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
	return
}

func GeneratePage25h(osnr int, temperature int) (page []byte) {
	page = append(page, make([]byte, 32)...)
	rand.Seed(time.Now().UTC().UnixNano())
	modOsnr := uint16(osnr + (rand.Intn(2*temperature)-temperature)/3)
	page = append(page, []byte{byte(modOsnr >> 8), byte(modOsnr & 0xFF)}...) // VDM real-time OSNR
	page = append(page, make([]byte, 94)...)

	return
}
