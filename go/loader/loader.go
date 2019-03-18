package loader

import (
	"encoding/json"
	"io/ioutil"
)

//LoadJSON Load json file and return your object
func LoadJSON(pathFile string, object interface{}) error {
	file, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return err
	}
	json.Unmarshal(file, object)
	return nil
}
