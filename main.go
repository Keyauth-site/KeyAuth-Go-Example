package main

import (
	EpicAuthApp "EpicAuth/EpicAuth"
	"fmt"
	"os"
	"time"
)

func Input(message string) string {
	fmt.Print(message)

	var input string
	fmt.Scanln(&input)
	return input
}

func main() {
	EpicAuthApp.Api(
		"EpicAuth",    // -- Application Name
		"mpgOizljNW", // -- Owner ID
		"1.1",        // -- Application Version
		"",           // -- Token Path (PUT NULL OR LEAVE BLANK IF YOU DON'T WANT TO USE TOKEN SYSTEM)
	)

	fmt.Println("[1] Login")
	fmt.Println("[2] Register")
	fmt.Println("[3] Upgrade")
	fmt.Println("[4] License Only Login")

	ans := Input("\nChoose your option: ")

	if ans == "1" {
		username := Input("Input username: ")
		password := Input("Input password: ")

		EpicAuthApp.Login(username, password)
	} else if ans == "2" {
		username := Input("Input username: ")
		password := Input("Input password: ")
		license := Input("Input license: ")

		EpicAuthApp.Register(username, password, license)
	} else if ans == "3" {
		username := Input("Input username: ")
		license := Input("Input license: ")

		EpicAuthApp.Upgrade(username, license)
	} else if ans == "4" {
		license := Input("Input license: ")

		EpicAuthApp.License(license)
	} else {
		fmt.Println("Invalid option")
		time.Sleep(2 * time.Second)
		main()
	}

	fmt.Println("\nUser Data:")
	fmt.Println("   Username: ", EpicAuthApp.Username)
	fmt.Println("   IP Address: ", EpicAuthApp.IP)
	fmt.Println("   HWID: ", EpicAuthApp.HWID)
	fmt.Println("   Created At: ", EpicAuthApp.CreatedDate)
	fmt.Println("   Last Login At: ", EpicAuthApp.LastLogin)
	fmt.Println("   Subscription: ", EpicAuthApp.Subscription)

	fmt.Println("\nExiting application in 10 seconds...")
	time.Sleep(10 * time.Second)
	os.Exit(0)
}
