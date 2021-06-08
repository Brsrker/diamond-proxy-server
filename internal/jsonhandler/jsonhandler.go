package jsonhandler

import (
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// JSONHandler is ...
type JSONHandler struct {
	file    string
	content []byte
}

// New is asd...
func New(file string) JSONHandler {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return JSONHandler{file, content}
}

//Refresh ...
func (jh *JSONHandler) Refresh() error {
	content, err := ioutil.ReadFile(jh.file)
	if err != nil {
		return err
	}
	jh.content = content
	return nil

}

//Set ...
func (jh *JSONHandler) Set(key string, value interface{}) error {
	json, err := sjson.SetBytes(jh.content, key, value)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(jh.file, json, 0644)
	if err != nil {
		return err
	}

	jh.Refresh()
	return nil
}

//Get ...
func (jh *JSONHandler) Get(key string) gjson.Result {
	return gjson.GetBytes(jh.content, key)
}

//GetJSON ...
func (jh *JSONHandler) GetJSON() []byte {
	return jh.content
}
