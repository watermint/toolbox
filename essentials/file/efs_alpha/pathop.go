package efs_alpha

type PathOps interface {
	Exist() (bool, error)
	Move(dest Path) error
	Copy(dest Path) error
}
