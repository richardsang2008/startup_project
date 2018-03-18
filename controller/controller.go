package controller

import (
	"github.com/richardsang2008/startup_project/data"
	"github.com/richardsang2008/startup_project/model"
	"github.com/richardsang2008/startup_project/utility"
	"fmt"
)

var (
	Data data.DataAccess
)
func CreatePerson(person model.Person) (string, error) {
	utility.MLog.Debug("Controller CreatePerson starting "+ person.FirstName)
	newid, err := Data.AddPerson(person.FirstName, person.LastName, person.City)
	if err != nil {
		utility.MLog.Error("Controller CreatePerson error firstname ",person.FirstName,err.Error())
		return "-1", err
	} else {
		utility.MLog.Debug("Controller CreatePerson ends "+ person.FirstName)
		return newid,nil
	}
}
func GetPerson(id uint) (*model.Person,error){
	utility.MLog.Debug("Controller GetPerson starting "+ fmt.Sprint(id))
	person,err:=Data.GetPerson(id)
	if err!= nil {
		utility.MLog.Error("Controller GetPerson error id ",fmt.Sprint(id),err.Error())
		return nil,err
	} else {
		utility.MLog.Debug("Controller GetPerson end" + fmt.Sprint(id))
		return person,nil
	}
}