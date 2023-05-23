package main

import (
	"fmt"
	"log"
	"myproject/db"
)

func main() {
	fmt.Print("hello")
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	res, erro := db.getEmployeeByID(1)
	log.Print(res)
	defer db.CloseDB()
}
