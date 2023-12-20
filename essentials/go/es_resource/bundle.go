package es_resource

type Bundle interface {
	Templates() Resource
	Messages() Resource
	Web() Resource
	Keys() Resource
	Images() Resource
	Data() Resource
}

func New(tpl, msg, web, key, img, dat Resource) Bundle {
	return &bundleImpl{
		tpl: tpl,
		msg: msg,
		web: web,
		key: key,
		img: img,
		dat: dat,
	}
}

// NewChainBundle creates a new bundle by merging multiple bundles.
// The order of bundles is important. The first bundle has the highest priority.
func NewChainBundle(bundles ...Bundle) Bundle {
	resTmp := make([]Resource, 0)
	resMsg := make([]Resource, 0)
	resWeb := make([]Resource, 0)
	resKey := make([]Resource, 0)
	resImg := make([]Resource, 0)
	resDat := make([]Resource, 0)
	for _, b := range bundles {
		resTmp = append(resTmp, b.Templates())
		resMsg = append(resMsg, b.Messages())
		resWeb = append(resWeb, b.Web())
		resKey = append(resKey, b.Keys())
		resImg = append(resImg, b.Images())
		resDat = append(resDat, b.Data())
	}
	return &bundleImpl{
		tpl: NewMergedResource(resTmp...),
		msg: NewMergedResource(resMsg...),
		web: NewMergedResource(resWeb...),
		key: NewMergedResource(resKey...),
		img: NewMergedResource(resImg...),
		dat: NewMergedResource(resDat...),
	}
}

func EmptyBundle() Bundle {
	return &bundleImpl{
		tpl: EmptyResource(),
		msg: EmptyResource(),
		web: EmptyResource(),
		key: EmptyResource(),
		img: EmptyResource(),
		dat: EmptyResource(),
	}
}

type bundleImpl struct {
	tpl Resource
	msg Resource
	web Resource
	key Resource
	img Resource
	dat Resource
}

func (z bundleImpl) Templates() Resource {
	return z.tpl
}

func (z bundleImpl) Messages() Resource {
	return z.msg
}

func (z bundleImpl) Web() Resource {
	return z.web
}

func (z bundleImpl) Keys() Resource {
	return z.key
}

func (z bundleImpl) Images() Resource {
	return z.img
}

func (z bundleImpl) Data() Resource {
	return z.dat
}
