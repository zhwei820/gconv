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
	sByte, _ := base64.StdEncoding.DecodeString(string(data))
	s, _ := aes[0].Decrypt(string(sByte))
	err = json.Unmarshal([]byte(s), v)
	if err != nil {
		return err
	}
	return nil
}

func WritePlainToFile(v interface{}, fn string) {
	os.Create(fn)
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(Export(v)))
	if err != nil {
		panic(err)
	}
}

func LoadPlainJsonFromFile(v interface{}, fn string) error {
	data, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
