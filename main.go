package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v8"
)

func main(){
	client := gocloak.NewClient("http://localhost:8080")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "admin", "admin", "master")
	if err != nil {
		panic("Something wrong with the credentials or url")
	}
	params := gocloak.GetUsersParams{}
	userArr, userErr := client.GetUsers(ctx, token.AccessToken, "master", params)
	if userErr != nil {
		panic("Something wrong with the credentials or url")
	}
	fmt.Println(userArr)
}
