package bst

type Comparable interface {
	CompareTo(interface{}) int
}

type Comparator interface {
	Compare(interface{}, interface{}) int
}
