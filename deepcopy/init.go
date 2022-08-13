package deepcopy

import "encoding/json"

func CopyExported(dst interface{}, src interface{}) error { // only for exported member
	if res, err := json.Marshal(src); err != nil {
		return err
	} else {
		err = json.Unmarshal(res, dst)
		return err
	}
}
