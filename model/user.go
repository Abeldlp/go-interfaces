package model

type User struct {
	Id   int
	Name string
}

func (u *User) SetId(id int) {
	u.Id = id
}

func (u *User) SetName(name string) {
	u.Name = name
}
