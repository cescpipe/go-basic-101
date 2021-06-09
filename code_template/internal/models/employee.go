package models

type Employee struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	EmpId     string `json:"emp_id" bson:"emp_id"`
}
