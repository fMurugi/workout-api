package main

import (
	"workout-api/data"
	"workout-api/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	data.InitDB()
	
	r := gin.Default()
	routes.RegisterWorkoutRoutes(r)

	r.Run(":8080")
}
