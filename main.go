package main

import(
	"github.com/gin-gonic/gin"
	"workout-api/routes"

)

func main(){
	r := gin.Default()
	routes.RegisterWorkoutRoutes(r)

	r.Run(":8080")
}
