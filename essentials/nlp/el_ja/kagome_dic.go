package el_ja

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
)

const (
	DictionaryIpa = "ipa"
	DictionaryUni = "uni"

	ModeNormal = "normal"
	ModeSearch = "search"
	ModeExtend = "extend"

	CacheNamespace = "nlp-ja-kagome"
)

type Container interface {
	// Load loads dictionary by name.
	Load(name string) (d *dict.Dict, err error)

	// LoadIpa loads IPA dictionary.
	LoadIpa() (d *dict.Dict, err error)

	// LoadUni loads UniDic dictionary.
	LoadUni() (d *dict.Dict, err error)

	// NewTokenizer creates a tokenizer.
	NewTokenizer(name string, omitBosEos bool) (t *tokenizer.Tokenizer, err error)

	// NewIpaTokenizer creates a tokenizer with IPA dictionary.
	NewIpaTokenizer(omitBosEos bool) (t *tokenizer.Tokenizer, err error)

	// NewUniTokenizer creates a tokenizer with UniDic dictionary.
	NewUniTokenizer(omitBosEos bool) (t *tokenizer.Tokenizer, err error)
}

func NewContainer(cache ec_file.File, logger esl.Logger) Container {
	return &containerImpl{
		cache:  cache,
		logger: logger,
	}
}

type containerImpl struct {
	cache  ec_file.File
	logger esl.Logger
}

func (z containerImpl) NewTokenizer(name string, omitBosEos bool) (t *tokenizer.Tokenizer, err error) {
	dic, err := z.Load(name)
	if err != nil {
		return nil, err
	}
	options := make([]tokenizer.Option, 0)
	if omitBosEos {
		options = append(options, tokenizer.OmitBosEos())
	}
	return tokenizer.New(dic, options...)
}

func (z containerImpl) NewIpaTokenizer(omitBosEos bool) (t *tokenizer.Tokenizer, err error) {
	return z.NewTokenizer(DictionaryIpa, omitBosEos)
}

func (z containerImpl) NewUniTokenizer(omitBosEos bool) (t *tokenizer.Tokenizer, err error) {
	return z.NewTokenizer(DictionaryUni, omitBosEos)
}

func (z containerImpl) LoadIpa() (d *dict.Dict, err error) {
	return z.Load(DictionaryIpa)
}

func (z containerImpl) LoadUni() (d *dict.Dict, err error) {
	return z.Load(DictionaryUni)
}

func (z containerImpl) Load(name string) (d *dict.Dict, err error) {
	l := z.logger.With(esl.String("name", name))
	var url string
	switch name {
	case DictionaryIpa:
		url = "https://raw.githubusercontent.com/watermint/kagome-dict/main/ipa/ipa.dict"
	case DictionaryUni:
		url = "https://raw.githubusercontent.com/watermint/kagome-dict/main/uni/uni.dict"
	default:
		return nil, fmt.Errorf("unknown dictionary: %s", name)
	}
	path, err := z.cache.Get(CacheNamespace, name, url)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		l.Debug("Unable to read file", esl.Error(err))
		return nil, err
	}
	zr, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		l.Debug("Unable to open zip", esl.Error(err))
		return nil, err
	}
	d, err = dict.Load(zr, true)
	if err != nil {
		l.Debug("Unable to load dictionary", esl.Error(err))
		return nil, err
	}
	return d, nil
}
