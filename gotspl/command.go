package gotspl

type TSPLCommand interface {
	GetMessage() ([]byte, error)
}
