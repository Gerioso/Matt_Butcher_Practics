package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	conf, err := yaml.ReadFile("./cong.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf.Get("path"))
	fmt.Println(conf.GetBool("enabled"))
}
