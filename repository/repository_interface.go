package repository

type IRepository[T any] interface {
	FindAll() ([]T, error)
	FindById(id int) (T, error)
	Update(*T) (*T, error)
	Insert(*T) error
	Delete(*T) error
}
