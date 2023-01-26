package main

import (
	"fmt"

	"example.com/m/pkg/config"
)

func main(){
	test := config.NewConfig()
	fmt.Println(test)
}