package main

import (
	"fmt"
	"github.com/olebedev/config"
	"io/ioutil"
	"path/filepath"
)
//https://godoc.org/github.com/olebedev/config
func main()  {

	currPath,err:=filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	fmt.Println(currPath)



	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	yamlString := string(file)

	cfg, err := config.ParseYaml(yamlString)
	//cfg, err := config.ParseJson(jsonString)

	// "localhost"
	host, err := cfg.String("development.database.host")
	if err != nil {
		panic(err)
	}
	fmt.Println(host)




}
