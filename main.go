package main

import (
	"fmt"
	"ftp/database"
	"ftp/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	initEnv()
	database.Init()
	port := os.Getenv("PORT")
	fmt.Print(port)
	handler := router.Init()
	_ =handler.Run(port)
	fmt.Println(os.Getenv("PORT"))
}

func initEnv(){
	err:=godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error: .env file cannot load")
	}
}
