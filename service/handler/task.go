package handler

import (
	"errors"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.New("Bad Request")})
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func (p *Handler) GetTask(c *gin.Context) {
	// id := c.Param("id")
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
		return
	}
	// t, err := getTaskById(id)
	t, err := p.ambiez.Storage.GetTask(c, queryID)
	if err != nil {
		fmt.Println("ERR ", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, t)
}

func (p *Handler) AddTask(c *gin.Context) {
	var newTask m.TaskRequest
	fmt.Println("context ", c)

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong JSON Format"})
		return
	}

	t, err := p.ambiez.AddTask(c, newTask)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, t)
}

func (p *Handler) UpdateTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
		return
	}
	var newTask m.TaskRequest
	//TODO: handle case param outside taskrequest params
	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Wrong JSON Format"})
		return
	}
	t, err := p.ambiez.UpdateTask(c, queryID, newTask)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Successfully update task", "id": t.ID})

}

func (p *Handler) ToggleTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Println("[TaskHandler][GetTask] bad request, err: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
		return
	}

	err = p.ambiez.ToggleTask(c, queryID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

}

func (p *Handler) RemoveTask(c *gin.Context) {
	queryID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "ID must be an integer"})
		return
	}
	_, err = p.ambiez.RemoveTask(c, queryID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

}
