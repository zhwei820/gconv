package gconv

import (
	"encoding/json"
	"os"
)

func WriteToFile(v interface{}, fn string) {
	os.Create(fn)
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	SetExportExpand(true)

	_, err = file.Write([]byte(Export(v)))
	if err != nil {
		panic(err)
	}
}

func LoadJsonFromFile(v interface{}, fn string) error {
	file, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
