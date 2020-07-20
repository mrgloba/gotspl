package gotspl

type TSPLCommandSequence interface {
	getTSPLCode() ([]byte, error)
}
