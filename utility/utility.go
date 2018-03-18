package utility
import (
	"github.com/richardsang2008/startup_project/model"
	"os"
	"encoding/json"
)
var (
	MLog Log
	MCache Cache
)
type Utility struct {

}
func LoadConfiguration(file string) model.Config {
	var config model.Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		panic("Can not LoadConfiguration: "+err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}