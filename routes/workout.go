package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"workout-api/data"
	"workout-api/models"
)

//register workout routes
func RegisterWorkoutRoutes(r *gin.Engine){
	r.POST("/workouts",createWorkout)
	r.GET("/workouts",listWorkouts)
}
func createWorkout(c *gin.Context){
	var w models.Workout
	if err := c.ShouldBindBodyWithJSON(&w); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	w.ID = uuid.New().String()
	data.Workouts[w.ID] = w
	c.JSON(http.StatusCreated,w)


}

func listWorkouts(c *gin.Context){
	workouts := []models.Workout{}
	for _,v := range data.Workouts{
		workouts = append(workouts,v)
	}
	c.JSON(http.StatusOK,workouts)
}