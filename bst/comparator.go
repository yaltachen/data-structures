package bst

type Comparator interface {
	Compare(interface{}, interface{}) int
}
