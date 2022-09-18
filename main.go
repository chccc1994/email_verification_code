package main

import (
	"email_code/models"
	"email_code/routers"
	"fmt"
)

func main() {
	fmt.Println("hello golang")
	models.InitDb()
	routers.InitRouter()
}