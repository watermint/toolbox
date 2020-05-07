package dbx_util

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"regexp"
	"strconv"
	"strings"
)

// If you use the Dropbox-API-Arg header, you need to make it "HTTP header safe".
// This means using JSON-style "\uXXXX" escape codes for the character 0x7F and all non-ASCII characters.
// @see https://www.dropbox.com/developers/reference/json-encoding
//
// Returns error if any char that is Unicode plane 1 or above
func HeaderSafeJson(p interface{}) (string, error) {
	l := es_log.Default()

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(p); err != nil {
		l.Debug("Unable to encode", es_log.Error(err))
	}
	sq := buf.String()
	sq1 := strings.Trim(sq, "\n")

	safeQuoted := strconv.QuoteToASCII(sq1)
	safeUnquoted1 := strings.ReplaceAll(safeQuoted, "\\\"", "\"")
	safeUnquoted2 := strings.ReplaceAll(safeUnquoted1, "\\\\", "\\")
	safeUnquoted3 := strings.Trim(safeUnquoted2, "\"")

	extCharPattern := regexp.MustCompile("\\\\U([0-9a-fA-F]{8})")
	if extCharPattern.MatchString(safeUnquoted3) {
		l.Debug("Found extended char",
			es_log.String("Encoded", safeUnquoted3),
			es_log.Strings("ExtChars", extCharPattern.FindAllString(safeUnquoted3, -1)),
		)
		return "", errors.New("does not support unicode extended character")
	}

	return safeUnquoted3, nil
}
