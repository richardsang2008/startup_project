package services

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/richardsang2008/startup_project/model"
	"github.com/richardsang2008/startup_project/controller"
	"github.com/richardsang2008/startup_project/utility"
	"net/http/httputil"
)

func CreatePerson(c *gin.Context) {
	var person model.Person
	c.BindJSON(&person)
	requestDump, err := httputil.DumpRequest(c.Request,true)
	if err !=nil {
		utility.MLog.Panic(err.Error())
	}
	utility.MLog.Debug("Webservice CreatePerson starting "+string(requestDump))
	newid, err := controller.CreatePerson(person)
	if err != nil {
		utility.MLog.Error("Webservice CreatePerson error "+ err.Error())
		c.JSON(500,gin.H{"message":err.Error()})
	} else {
		utility.MLog.Debug("Webservice CreatePerson end "+string(requestDump))
		c.JSON(200, gin.H{"id": newid})
	}
}
func GetPerson(c *gin.Context) {

	id := c.Params.ByName("id")
	utility.MLog.Error("Webservice GetPerson starting "+ id)
	var person *model.Person
	var err error
	u, _ := strconv.ParseUint(id, 10, 32)
	person, err = controller.GetPerson(uint(u))
	if err != nil {
		utility.MLog.Error("Webservice GetPerson error "+ err.Error())
		c.JSON(404,gin.H{"message":err.Error()})
	} else {
		utility.MLog.Error("Webservice GetPerson end "+ id)
		c.JSON(200, *person)
	}

}
