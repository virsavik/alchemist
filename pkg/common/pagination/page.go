package pagination

type Page[T any] struct {
	Items []T
	Index int64
	Size  int64
	Total int64
}
