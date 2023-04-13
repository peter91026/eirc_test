package main

import (
	"eirc/bpm"
	"eirc/structure"
	"encoding/json"
	"fmt"
)

func main() {
	post()
}

func get() {

	read := []structure.Read{}
	err := json.Unmarshal(bpm.Get_test(), &read)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(read)
	fmt.Println(bpm.Get_test())
}

func post() {
	read := structure.Read{
		Title:  "hello world",
		UserId: 5,
	}
	marshal, err := json.Marshal(read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bpm.Post_test(marshal))
}
