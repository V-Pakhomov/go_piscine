package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
)

type DataBase interface {
	read(filename string) []Cake
	write()
}

type Ingredient struct {
	Name string					`json:"ingredient_name" xml:"itemname"`
	Count string				`json:"ingredient_count" xml:"itemcount"`
	Unit string					`json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name string					`json:"name" xml:"name"`
	Time string					`json:"time" xml:"stovetime"`
	Ingredients []Ingredient	`json:"ingredients" xml:"ingredients>item"`
}

type JSON struct {
	Cakes []Cake				`json:"cake"`
}

type XML struct {
	XMLName xml.Name			`xml:"recipes"`
	Cakes []Cake				`xml:"cake"`
}

func (x *XML) read(filename string) []Cake {
	data := fileToBytes(filename)
	xml.Unmarshal(data, &x)
	return x.Cakes
}

func (x *XML) write() {
	prettyXML, err := xml.MarshalIndent(x, "", "    ")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n", prettyXML)
}

func (j *JSON) write() {
	prettyJSON, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n", prettyJSON)
}

func (j *JSON) read(filename string) []Cake {
	data := fileToBytes(filename)
	json.Unmarshal(data, &j)
	return j.Cakes
}

func fileToBytes(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}
