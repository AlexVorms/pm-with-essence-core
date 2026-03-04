package controllers

import (
	"net/http"
	"pm-with-essence/cmd/service/taskService"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *taskService.TaskService
}

func NewTaskController(taskService *taskService.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

// CreateTask
// @Summary create new task
// @Tags Task
// @Description create new task
// @Produce json
// @Success 200
// @Router       /task [post]
func (tc *TaskController) CreateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// DeleteTask
// @Summary delete task
// @Tags Task
// @Description delete task
// @Produce json
// @Success 200
// @Router       /task/{taskId} [delete]
func (tc *TaskController) DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// GetTask
// @Summary get task
// @Tags Task
// @Description get task
// @Produce json
// @Success 200
// @Router        /task/{taskId} [get]
func (tc *TaskController) GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// GetAllTasks
// @Summary get all tasks
// @Tags Task
// @Description get all tasks
// @Produce json
// @Success 200
// @Router       /task [get]
func (tc *TaskController) GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// UpdateTask
// @Summary update task
// @Tags Task
// @Description update task
// @Produce json
// @Success 200
// @Router       /task/{taskId} [put]
func (tc *TaskController) UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
