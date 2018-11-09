package avl

type Comparator interface {
	Compare(interface{}, interface{}) int
}
