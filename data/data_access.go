package data

import (
	"github.com/jinzhu/gorm"
	"github.com/richardsang2008/startup_project/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/richardsang2008/startup_project/utility"
	"fmt"
)

var (
	err        error
	DataBase *gorm.DB
)

type DataAccess struct{

}
func (s *DataAccess) New() {
	DataBase,err =gorm.Open("mysql", "mapuser:password@tcp(127.0.0.1:3306)/pgpool?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		utility.MLog.Panic("Error creating connection pool: " + err.Error())
	}
	utility.MLog.New("config/appconfig.json")
	DataBase.AutoMigrate(&model.Person{})
}
func(s *DataAccess) Close() {
	utility.MLog.Info("Closing database")
	DataBase.Close()
}
func (s *DataAccess) AddPerson(firstname,lastname,city string) (string,error){
	utility.MLog.Debug("DataAccess AddPerson starting "+firstname)
	person:=model.Person{FirstName:firstname,LastName:lastname,City:city}
	DataBase.Create(&person)
	utility.MLog.Debug("DataAccess AddPerson ending")
	ret := fmt.Sprint(person.ID)
	return ret,nil
}
func (s*DataAccess) GetPerson(id uint) (*model.Person,error) {
	utility.MLog.Debug("DataAccess GetPerson starting " + fmt.Sprint(id))
	var person model.Person
	if err:= DataBase.Where("id=?",id).First(&person).Error;err!=nil {
		utility.MLog.Error("DataAccess GetPerson failed "+err.Error())
		return nil,err
	}else{
		utility.MLog.Debug("DataAccess GetPerson success " +fmt.Sprint(id))
		return &person,nil
	}
}
