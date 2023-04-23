package model

type Employee struct {
	Id       int
	Position string
	User
}

func (e *Employee) SetId(id int) {
	e.Id = id
}

func (e *Employee) SetPosition(position string) {
	e.Position = position
}
