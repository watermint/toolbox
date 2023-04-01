package sv_file

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/figma/api/fg_client"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"regexp"
)

type MsgFile struct {
	VerifyOk                       app_msg.Message
	VerifyFileKeyTooLong           app_msg.Message
	VerifyFileKeyInvalidCharacters app_msg.Message
}

var (
	MFile = app_msg.Apply(&MsgFile{}).(*MsgFile)
)

type VerifyResult int

const (
	VerifyFileKeyLooksOkay VerifyResult = iota
	VerifyFileKeyTooLong
	VerifyFileKeyInvalidChar
)

const (
	// FileKeyMaxLength maximum length of the file key.
	// There is no clear definition in the API doc as of implementation.
	// But keep it smaller for safe.
	// Current file key length is ~24 alpha-numeric characters.
	// This threshold includes buffer for future changes in Figma API side.
	FileKeyMaxLength = 36
)

var (
	// FileKeyRegex is the allowed Team Id pattern.
	// There is no clear definition in the API doc as of implementation.
	// Current file key looks like alpha-numeric.
	FileKeyRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)
)

func VerifyFileKey(fileKey string) (VerifyResult, app_msg.Message) {
	if FileKeyMaxLength < len(fileKey) {
		return VerifyFileKeyTooLong, MFile.VerifyFileKeyTooLong
	}
	if !FileKeyRegex.MatchString(fileKey) {
		return VerifyFileKeyInvalidChar, MFile.VerifyFileKeyInvalidCharacters
	}
	return VerifyFileKeyLooksOkay, MFile.VerifyOk
}

type ImageOpts struct {
	Ids    string  `url:"ids"`
	Scale  *string `url:"scale,omitempty"`
	Format *string `url:"format,omitempty"`
}

var (
	ImageFormats = []string{
		"jpg", "png", "svg", "pdf",
	}
)

const (
	ImageScaleMin     = 1
	ImageScaleMax     = 400
	ImageScaleDefault = 100
)

type File interface {
	Info(key string) (doc *mo_file.Document, err error)
	Image(key string, ids string, scale int, format string) (urls map[string]string, err error)
}

func New(client fg_client.Client) File {
	return &fileImpl{
		client: client,
	}
}

type fileImpl struct {
	client fg_client.Client
}

func (z fileImpl) Image(key string, ids string, scale int, format string) (urls map[string]string, err error) {
	if r, m := VerifyFileKey(key); r != VerifyFileKeyLooksOkay {
		return nil, errors.New(m.Key())
	}

	opts := &ImageOpts{
		Ids: ids,
	}
	if scale != ImageScaleDefault {
		scaleStr := fmt.Sprintf("%0.00f", float32(scale)/100.0)
		opts.Scale = &scaleStr
	}
	if format != "" {
		opts.Format = &format
	}
	res := z.client.Get("images/"+key, api_request.Query(opts))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	j := res.Success().Json()
	if resErr, found := j.Find("err"); found {
		if !resErr.IsNull() {
			return nil, errors.New(resErr.RawString())
		}
	}
	images, found := j.FindObject("images")
	if !found {
		return nil, errors.New("unexpect response")
	}
	urls = make(map[string]string)

	for k, v := range images {
		if url, found := v.String(); !found {
			return nil, errors.New("no download url in the response")
		} else {
			urls[k] = url
		}
	}
	return urls, nil
}

func (z fileImpl) Info(key string) (doc *mo_file.Document, err error) {
	if r, m := VerifyFileKey(key); r != VerifyFileKeyLooksOkay {
		return nil, errors.New(m.Key())
	}

	res := z.client.Get("files/" + key)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	doc = &mo_file.Document{}
	j := res.Success().Json().RawString()
	err = json.Unmarshal([]byte(j), doc)
	return doc, err
}
