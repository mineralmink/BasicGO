package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[int]*Todo{
	1: &Todo{ID: 1, Title: "pay check", Status: "active"},
	2: &Todo{ID: 2, Title: "pay credit", Status: "completed"},
	3: &Todo{ID: 3, Title: "homework", Status: "inactive"},
	4: &Todo{ID: 4, Title: "buy new shoes", Status: "completed"},
	//key 1 with value
}

func getTodoHandler(c *gin.Context) {
	status := c.Query("status")
	items := []*Todo{}
	for _, item := range todos {
		if status != "" {
			if item.Status == status {
				items = append(items, item)
				//e.g. http://localhost:1234/todos?status=completed
			}
		} else {
			items = append(items, item)
		}

	}
	c.JSON(http.StatusOK, items)
}
func createTodoHandler(c *gin.Context) {
	t := Todo{}
	//r.body and read body -> bind json and send to &t
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := len(todos)
	id++
	t.ID = id
	todos[t.ID] = &t
	c.JSON(http.StatusCreated, todos) //status created -> 201 created
}

func deleteTodosHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	delete(todos, id)
	c.JSON(http.StatusOK, todos)
}

func updateTodosHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t := todos[id]
	if err := c.ShouldBindJSON(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, t)
}

func getTodoByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) //atomic to integer
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, ok := todos[id]

	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodoHandler)
	r.GET("/todos/:id", getTodoByIdHandler)
	r.POST("/todos", createTodoHandler)
	r.PUT("/todos/:id", updateTodosHandler)
	r.DELETE("/todos/:id", deleteTodosHandler)
	r.Run(":1234")
}
