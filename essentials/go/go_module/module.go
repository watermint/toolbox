package go_module

import "runtime/debug"

type Module interface {
	// Path to the module.
	Path() string

	// Version of the module.
	Version() string

	// Sum of the module.
	Sum() string

	// Licenses of the module.
	Licenses() []License
}

func NewModule(m *debug.Module, licenses []License) Module {
	return moduleImpl{
		module:   m,
		licenses: licenses,
	}
}

type moduleImpl struct {
	module   *debug.Module
	licenses []License
}

func (z moduleImpl) Path() string {
	return z.module.Path
}

func (z moduleImpl) Version() string {
	return z.module.Version
}

func (z moduleImpl) Sum() string {
	return z.module.Sum
}

func (z moduleImpl) Licenses() []License {
	return z.Licenses()
}
