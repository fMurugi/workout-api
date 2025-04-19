package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"workout-api/data"
	"workout-api/models"
)

//register workout routes
func RegisterWorkoutRoutes(r *gin.Engine){
	r.POST("/workouts",createWorkout)
	r.GET("/workouts",listWorkouts)
	r.DELETE("/workouts",deleteWorkout)
}
func createWorkout(c *gin.Context){
	var w models.Workout
	if err := c.ShouldBindBodyWithJSON(&w); err != nil{
		c.JSON(400, gin.H{"error":err.Error()})
		return
	}
	
	w.ID = uuid.New().String()
	result := data.DB.Create(&w)
	if result.Error != nil{
		erroMessage := result.Error.Error()
		c.JSON(500,gin.H{"error": erroMessage})
		return
	}

	c.JSON(201,w)
}

func listWorkouts(c *gin.Context){
	workouts := []models.Workout{}
	
	if err := data.DB.Find(&workouts).Error; err!=nil{
		c.JSON(500,gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,workouts)
}

func deleteWorkout(c *gin.Context){
	id := c.Param("id")
	if err:= data.DB.Delete(&models.Workout{}, "id = ?", id).Error; err != nil{
		c.JSON(500,gin.H{"Error": err.Error()})

	}
	c.JSON(200,gin.H{"message":"Deleted"})
}