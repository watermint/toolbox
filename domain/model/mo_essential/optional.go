package mo_essential

type Optional interface {
	Ok() bool
}

type OptionalMutable interface {
	// mark value as unset
	Unset()
}
