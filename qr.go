package main

import (
	"github.com/skip2/go-qrcode"
)

// GenerateQRCode menghasilkan QR Code dari data yang diberikan
func GenerateQRCode(data string) string {
	qrCode, _ := qrcode.New(data, qrcode.Low)
	qrCodeData := qrCode.ToSmallString(false)
	return qrCodeData
}
