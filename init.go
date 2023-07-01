package gconv

import "os"

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
