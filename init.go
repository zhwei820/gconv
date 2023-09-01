package gconv

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

var key = "XefXzFKWgIVkvWTf"

func WriteToFile(v interface{}, fn string, aes ...*AES) {
	if aes == nil {
		aes = append(aes, &AES{
			Key: []byte(key),
		})
	}
	os.Create(fn)
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	SetExportExpand(true)
	s, _ := aes[0].Encrypt(Export(v))

	s = base64.StdEncoding.EncodeToString([]byte(s))
	_, err = file.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}

func LoadJsonFromFile(v interface{}, fn string, aes ...*AES) error {
	if aes == nil {
		aes = append(aes, &AES{
			Key: []byte(key),
		})
	}
	data, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	// Convert []byte to string
	s, _ := aes[0].Decrypt(string(data))
	sByte, _ := base64.StdEncoding.DecodeString(s)
	err = json.Unmarshal([]byte(sByte), v)
	if err != nil {
		return err
	}
	return nil
}
