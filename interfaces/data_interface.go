package interfaces

import "github.com/richardsang2008/startup_project/model"

type DataInterface interface{
	DeletePerson(id uint) (string,error)
	AddPerson(firstname,lastname,city string) (string,error)
	GetPerson(id uint) (*model.Person,error)
	UpdatePerson(id uint,firstname,lastname,city string) (string,error)
}
