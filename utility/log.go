package utility

import (
	log "github.com/sirupsen/logrus"
	"os"

)

type Log struct{

}
func (l *Log) New(configFilePath string) {
	log.SetLevel(log.DebugLevel)
	f, err := os.OpenFile("mylog.log", os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	//f,err :=os.OpenFile()

	//LoadConfiguration(configFilePath)
}
func (l *Log) Debug(args ...interface{}){
	log.Debug(args)
}
func (l *Log) Panic(args ...interface{}) {
	log.Panic(args)
}
func (l *Log) Info(args ...interface{}) {
	log.Info(args)
}
func (l *Log) Error(args ...interface{}){
	log.Error(args)
}
func (l *Log) Warning(args ...interface{}){
	log.Warning(args)
}