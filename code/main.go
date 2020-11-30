package main

import (
	"fmt"

	"github.com/osamakhanedu/webservice/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "ABC",
		LastName:  "EFG",
	}

	fmt.Println(u)
}
