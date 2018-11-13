package rbt

type Comparator interface {
	Compare(interface{}, interface{}) int
}
