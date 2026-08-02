package models

type PageResult[T any] struct {
	Records []T
	Total   int64
}
