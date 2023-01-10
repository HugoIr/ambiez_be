package handler

import (
	"fmt"
	m "hugdev/ambiez-go/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// *gin.Context gives us the requests data
func (p *Handler) GetTasks(c *gin.Context) {

	res, err := p.ambiez.GetTaskAll(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Error aja"})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

// func getTaskById(id string) (*m.Task, error) {
// 	for i, t := range tasks {
// 		if t.ID == id {
// 			return &tasks[i], nil
// 		}
// 	}
// 	return nil, errors.New("Task is not found")
// }

func (p *Handler) GetTask(c *gin.Context) {
	// id := c.Param("id")
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		return
	}
	// t, err := getTaskById(id)
	t, err := p.ambiez.Storage.GetTask(c, queryID)
	if err != nil {
		fmt.Println("ERR ", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error(), "status": http.StatusNotFound})
		return
	}
	c.IndentedJSON(http.StatusOK, t)
}

func (p *Handler) AddTask(c *gin.Context) {
	var newTask m.TaskRequest
	fmt.Println("context ", c)

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Wrong JSON Format"})
		return
	}

	t, err := p.ambiez.AddTask(c, newTask)
	if err != nil {
		log.Println("[TaskHandler][AddTask] bad request, err: ", err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, t)
}

func (p *Handler) UpdateTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		return
	}
	var newTask m.TaskRequest
	//TODO: handle case param outside taskrequest params
	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Wrong JSON Format"})
		return
	}
	t, err := p.ambiez.UpdateTask(c, queryID, newTask)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Successfully update task", "id": t.ID})

}

func (p *Handler) ToggleTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		return
	}

	err = p.ambiez.ToggleTask(c, queryID)

	if err != nil {
		//TODO: pikirin error yang lebih baik
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": err.Error()})
		return
	}

}

func (p *Handler) RemoveTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		return
	}
	p.ambiez.RemoveTask(c, queryID)

}
