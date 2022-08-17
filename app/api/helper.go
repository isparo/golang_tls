package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Help() {
	content, err := ioutil.ReadFile("./dist/thermopylae.txt")

	fmt.Println("error : ", err)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func Help2() {

	myd, err := os.Getwd()

	fmt.Println("err : ", err)
	fmt.Println("myd : ", myd)

	Help()
}
