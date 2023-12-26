package dd_screenshot

import (
	"bytes"
	"fmt"
	"image"
	"text/template"
	"time"
)

func SnapName(nameTmpl *template.Template, displayId int, seq int64, bounds image.Rectangle) (string, error) {
	now := time.Now()
	nowUtc := now.UTC()
	nameValues := map[string]string{
		"Date":          now.Format("2006-01-02"),
		"DateUTC":       nowUtc.Format("2006-01-02"),
		"DisplayHeight": fmt.Sprintf("%d", bounds.Size().Y),
		"DisplayId":     fmt.Sprintf("%d", displayId),
		"DisplayWidth":  fmt.Sprintf("%d", bounds.Size().X),
		"DisplayX":      fmt.Sprintf("%d", bounds.Min.X),
		"DisplayY":      fmt.Sprintf("%d", bounds.Min.Y),
		"Sequence":      fmt.Sprintf("%05d", seq),
		"Time":          now.Format("15-04-05"),
		"TimeUTC":       nowUtc.Format("15-04-05"),
		"Timestamp":     now.Format("20060102-150405"),
		"TimestampUTC":  now.Format("20060102-150405Z"),
	}
	var nameBuf bytes.Buffer
	err := nameTmpl.Execute(&nameBuf, nameValues)
	if err != nil {
		return "", err
	}
	return nameBuf.String(), nil
}
