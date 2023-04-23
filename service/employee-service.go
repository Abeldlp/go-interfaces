package service

import "github.com/Abeldlp/interfaces/model"

type EmployeeService struct {
	Employees []model.Employee
}

func NewEmployeeService() Service[model.Employee] {
	return &EmployeeService{}
}

func (e *EmployeeService) GetAll() []model.Employee {
	return e.Employees
}

func (e *EmployeeService) Create(employee model.Employee) model.Employee {
	employee.Id = len(e.Employees) + 1
	e.Employees = append(e.Employees, employee)
	return employee
}
