package main

import (
	database "hugdev/ambiez-go/database"
	taskHandler "hugdev/ambiez-go/handler"
	taskmodule "hugdev/ambiez-go/taskmodule"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}
	dbConfig := database.Config{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     port,
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}

	log.Println("Initializing DB Connection")
	db := database.GetDatabaseConnection(dbConfig)

	log.Println("Initializing Usecase")
	am := taskmodule.NewTaskModule(db)

	log.Println("Initializing Handler")
	ts := taskHandler.NewAmbiezHandler(am)

	router := gin.Default()
	router.GET("/tasks", ts.GetTasks)
	router.GET("/tasks/todo", ts.GetTodoTasks)
	router.GET("/tasks/completed", ts.GetCompletedTasks)
	router.GET("/tasks/:id", ts.GetTask)
	router.POST("/tasks", ts.AddTask)
	router.PATCH("/tasks/:id", ts.UpdateTask)
	router.PATCH("/tasks/toggle/:id", ts.ToggleTask)
	router.DELETE("/tasks/:id", ts.RemoveTask)
	router.Run("0.0.0.0:9000")
}
