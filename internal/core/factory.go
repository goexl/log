package core

type Factory interface {
	New() (Executor, error)
}
