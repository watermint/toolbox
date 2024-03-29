package es_resource

type Bundle interface {
	Build() Resource
	Data() Resource
	Images() Resource
	Keys() Resource
	Messages() Resource
	Release() Resource
	Templates() Resource
	Web() Resource
}

func New(tpl, msg, web, key, img, dat, bld, rel Resource) Bundle {
	return &bundleImpl{
		tpl: tpl,
		msg: msg,
		web: web,
		key: key,
		img: img,
		dat: dat,
		bld: bld,
		rel: rel,
	}
}

// NewChainBundle creates a new bundle by merging multiple bundles.
// The order of bundles is important. The first bundle has the highest priority.
func NewChainBundle(langCodes []string, bundles ...Bundle) Bundle {
	resTmp := make([]Resource, 0)
	resWeb := make([]Resource, 0)
	resKey := make([]Resource, 0)
	resImg := make([]Resource, 0)
	resDat := make([]Resource, 0)
	resBld := make([]Resource, 0)
	resRel := make([]Resource, 0)
	for _, b := range bundles {
		resTmp = append(resTmp, b.Templates())
		resWeb = append(resWeb, b.Web())
		resKey = append(resKey, b.Keys())
		resImg = append(resImg, b.Images())
		resDat = append(resDat, b.Data())
		resBld = append(resBld, b.Build())
		resRel = append(resRel, b.Release())
	}
	resMsg := MergedMessageResource(bundles, langCodes)
	return &bundleImpl{
		tpl: NewMergedResource(resTmp...),
		msg: NewMergedResource(resMsg),
		web: NewMergedResource(resWeb...),
		key: NewMergedResource(resKey...),
		img: NewMergedResource(resImg...),
		dat: NewMergedResource(resDat...),
		bld: NewMergedResource(resBld...),
		rel: NewMergedResource(resRel...),
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
		bld: EmptyResource(),
		rel: EmptyResource(),
	}
}

type bundleImpl struct {
	bld Resource
	dat Resource
	img Resource
	key Resource
	msg Resource
	rel Resource
	tpl Resource
	web Resource
}

func (z bundleImpl) Release() Resource {
	return z.rel
}

func (z bundleImpl) Build() Resource {
	return z.bld
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
