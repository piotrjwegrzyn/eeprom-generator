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

func GeneratePageLow(module Module) (page []byte) {
	page = append(page, byte(module.SFF8024Identifier))
	for i := 0; i < 127; i++ {
		page = append(page, byte(i))
	}
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

func GeneratePage02h() (page []byte) {
	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
	return
}

func GeneratePage04h() (page []byte) {
	for i := 0; i < 128; i++ {
		page = append(page, byte(i))
	}
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

func GeneratePage24h() (page []byte) {
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
