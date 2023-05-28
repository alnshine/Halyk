package flatjson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func Unmarshal(data []byte, dst interface{}) error {

	k := reflect.TypeOf(dst).Kind()
	if k != reflect.Ptr {
		fmt.Println("err net ssylka")
		os.Exit(1)
	}

	oldM := make(map[string]interface{})
	err := json.Unmarshal(data, &oldM)
	if err != nil {
		return err
	}
	resM := make(map[string]interface{})
	rec(oldM, resM, "")
	newData, err := json.Marshal(resM)
	if err != nil {
		return err
	}
	err = json.Unmarshal(newData, dst)
	if err != nil {
		return err
	}
	return nil
}
func rec(old, new map[string]interface{}, pres string) {
	for key, value := range old {
		newKey := pres + key
		switch value.(type) {
		case map[string]interface{}:
			rec(value.(map[string]interface{}), new, newKey)
		default:
			new[newKey] = value
		}
	}
}

//func Marshal(src interface{}) ([]byte, error) {
//	// TODO: Write code here
//	return nil, nil
//}
