package sort

type Comparable interface {
	Less(Comparable) bool
}
