package db

type Repository interface {
	Insert()
	Select()
	Update()
	Delete()
}
