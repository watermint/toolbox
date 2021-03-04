package qrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/watermint/toolbox/essentials/log/esl"
	"image/png"
	"os"
)

const (
	qrCodeErrorCorrectionLevelL = "l"
	qrCodeErrorCorrectionLevelM = "m"
	qrCodeErrorCorrectionLevelQ = "q"
	qrCodeErrorCorrectionLevelH = "h"
	qrCodeEncodeAuto            = "auto"
	qrCodeEncodeNumeric         = "numeric"
	qrCodeEncodeAlphaNumeric    = "alpha_numeric"
	qrCodeEncodeUnicode         = "unicode"
)

var (
	qrCodeErrorCorrectionLevels = []string{
		qrCodeErrorCorrectionLevelL,
		qrCodeErrorCorrectionLevelM,
		qrCodeErrorCorrectionLevelQ,
		qrCodeErrorCorrectionLevelH,
	}
	qrCodeEncodes = []string{
		qrCodeEncodeAuto,
		qrCodeEncodeNumeric,
		qrCodeEncodeAlphaNumeric,
		qrCodeEncodeUnicode,
	}
)

func fromLevelString(lv string) qr.ErrorCorrectionLevel {
	switch lv {
	case qrCodeErrorCorrectionLevelL:
		return qr.L
	case qrCodeErrorCorrectionLevelM:
		return qr.M
	case qrCodeErrorCorrectionLevelQ:
		return qr.Q
	case qrCodeErrorCorrectionLevelH:
		return qr.H
	}
	return qr.M
}

func fromEncodeString(en string) qr.Encoding {
	var encode qr.Encoding = qr.Auto
	switch en {
	case "auto":
		encode = qr.Auto
	case "numeric":
		encode = qr.Numeric
	case "alpha_numeric":
		encode = qr.AlphaNumeric
	case "unicode":
		encode = qr.Unicode
	}
	return encode
}

func createQrCodeImage(l esl.Logger, path, text string, size int, errorCorrection, encode string) error {
	qrc, err := qr.Encode(text, fromLevelString(errorCorrection), fromEncodeString(encode))
	if err != nil {
		l.Debug("Unable to encode", esl.Error(err))
		return err
	}
	scaled, err := barcode.Scale(qrc, size, size)
	if err != nil {
		l.Debug("Unable to scale", esl.Error(err))
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		l.Debug("Unable to create output file", esl.Error(err))
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if err = png.Encode(out, scaled); err != nil {
		l.Debug("Unable to encode to the image", esl.Error(err))
		return err
	}
	return nil
}
