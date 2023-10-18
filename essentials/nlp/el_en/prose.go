package el_en

import (
	"github.com/watermint/prose/v3"
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"sync"
)

const (
	CacheNamespace = "nlp-en-prose"
)

type Container interface {
	NewDocument(text string) (doc *prose.Document, err error)
}

func NewContainer(cache ec_file.File, logger esl.Logger) Container {
	return &containerImpl{
		cache:  cache,
		logger: logger,
	}
}

type containerImpl struct {
	cache      ec_file.File
	logger     esl.Logger
	model      *prose.Model
	modelMutex sync.Mutex
}

func (z *containerImpl) loadModel(path string) (model *prose.Model, err error) {
	l := z.logger
	l.Debug("Load model", esl.String("path", path))
	//defer func() {
	//	if failure := recover(); failure != nil {
	//		l.Debug("Unable to load model", esl.Any("failure", failure))
	//		err = fmt.Errorf("unable to load model: %v", failure)
	//	}
	//}()
	model = prose.ModelFromDisk(path)
	return
}

func (z *containerImpl) loadData() (model *prose.Model, err error) {
	z.modelMutex.Lock()
	defer z.modelMutex.Unlock()
	if z.model != nil {
		return z.model, nil
	}

	l := z.logger
	resources := []string{
		"AveragedPerceptron/tags.gob",
		"AveragedPerceptron/weights.gob",
		"AveragedPerceptron/classes.gob",
		"Maxent/words.gob",
		"Maxent/mapping.gob",
		"Maxent/labels.gob",
		"Maxent/weights.gob",
	}
	urlBase := "https://raw.githubusercontent.com/watermint/prose/master/model/"

	for _, r := range resources {
		l.Debug("Load resource", esl.String("resource", r))
		if _, err := z.cache.Get(CacheNamespace, r, urlBase+r); err != nil {
			l.Debug("Unable to retrieve cache", esl.Error(err), esl.String("resource", r))
			return nil, err
		}
	}
	z.model, err = z.loadModel(z.cache.Path(CacheNamespace))
	return z.model, err
}

func (z *containerImpl) NewDocument(text string) (doc *prose.Document, err error) {
	model, err := z.loadData()
	if err != nil {
		return nil, err
	}
	doc, err = prose.NewDocument(text,
		prose.WithTagging(true),
		prose.WithExtraction(true),
		prose.WithSegmentation(true),
		prose.UsingModel(model),
	)
	if err != nil {
		return nil, err
	}
	return
}
