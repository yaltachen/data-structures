package heap

type Comparator interface {
	Compare(interface{}, interface{}) int
}
