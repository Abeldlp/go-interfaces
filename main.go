package main

import (
	"fmt"
	"github.com/Abeldlp/interfaces/model"
	"github.com/Abeldlp/interfaces/service"
)

func main() {
	usrSrv := service.NewUserService()
	empSrv := service.NewEmployeeService()

	usrSrv.Create(model.User{
		Name: "Abel",
	})

	rob := usrSrv.Create(model.User{
		Name: "Robin",
	})

	empSrv.Create(model.Employee{
		Position: "Developer",
		User:     rob,
	})

	users := usrSrv.GetAll()
	employees := empSrv.GetAll()

	fmt.Println(users)
	fmt.Println(employees[0].Name)
}
