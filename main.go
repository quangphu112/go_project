package main

import (
	"fmt"
	"go_project/databases"
	"go_project/routes"
	"log"

	"go_project/initializers"
)

func init() { 
	initializers.LoadEnvVariables()
}
func main() {
	db, err := databases.ConnectDB()
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}
	fmt.Println("Connected to database successfully", db)

	r := routes.UserRouter(db)
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
