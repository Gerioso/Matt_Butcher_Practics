package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var conf Configuration
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
