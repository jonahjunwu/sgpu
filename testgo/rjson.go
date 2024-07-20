package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJsonFile(jfile string) ([]string, error) {

	type Inputlog struct {
		Data  []string `json: "data"`
		Image string   `json: "image"`
	}

	//file, _ := ioutil.ReadFile("input.json")
	jsonFile, err := os.Open(jfile)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var wFile Inputlog
	json.Unmarshal([]byte(byteValue), &wFile)
	fmt.Println(wFile)
	fmt.Println(wFile.Data)

	return wFile.Data, nil
}

func main() {
	whisperFiles, _ := readJsonFile("input.json")
	fmt.Println(whisperFiles)
	for _, data := range whisperFiles {
		fmt.Println(data)
	}
}
