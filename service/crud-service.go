package service

type Service[T any] interface {
	GetAll() []T
	Create(T) T
}
