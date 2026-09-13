package taskController

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/cmd/api/model/taskDTO"
	"pm-with-essence/internal/api/app_errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskController struct {
	taskService TaskService
}
type TaskService interface {
	CreateTask(columnID uuid.UUID, model taskDTO.TaskDTO) *app_errors.HttpError
	CompleteTask(taskID uuid.UUID) *app_errors.HttpError
	UpdateTask(taskID uuid.UUID, model taskDTO.TaskDTO) *app_errors.HttpError
	DeleteTask(taskID uuid.UUID) *app_errors.HttpError
	GetTask(taskID uuid.UUID) (*taskDTO.GetTaskDTO, *app_errors.HttpError)
	ChangeTaskColumn(taskID uuid.UUID, columnID uuid.UUID) *app_errors.HttpError
}

func NewTaskController(taskService TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

// CreateTask
// @Summary create new task
// @Tags Task
// @Description create new task
// @Produce json
// @Param columnId path string true "ColumnData ID"
// @Param newTaskDTO body taskDTO.TaskDTO true "newTaskDTO"
// @Success 200
// @Router       /task/{columnId} [post]
func (tc *TaskController) CreateTask(c *gin.Context) {
	columnId := c.Param("columnId")
	columnUUID, err1 := uuid.Parse(columnId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	var newTaskDTO taskDTO.TaskDTO
	if err := c.ShouldBindJSON(&newTaskDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.taskService.CreateTask(columnUUID, newTaskDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// DeleteTask
// @Summary delete task
// @Tags Task
// @Description delete task
// @Produce json
// @Param taskId path string true "Task ID"
// @Success 200
// @Router       /task/{taskId} [delete]
func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")
	taskUUID, err1 := uuid.Parse(taskId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := tc.taskService.DeleteTask(taskUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// GetTask
// @Summary get task
// @Tags Task
// @Description get task
// @Produce json
// @Param taskId path string true "Task ID"
// @Success 200
// @Router        /task/{taskId} [get]
func (tc *TaskController) GetTask(c *gin.Context) {
	taskId := c.Param("taskId")
	taskUUID, err1 := uuid.Parse(taskId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	task, err := tc.taskService.GetTask(taskUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, task)
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
// @Param taskId path string true "Task ID"
// @Param newTaskDTO body taskDTO.TaskDTO true "newTaskDTO"
// @Success 200
// @Router       /task/{taskId} [put]
func (tc *TaskController) UpdateTask(c *gin.Context) {
	taskId := c.Param("taskId")
	taskUUID, err1 := uuid.Parse(taskId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	var newTaskDTO taskDTO.TaskDTO
	if err := c.ShouldBindJSON(&newTaskDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := tc.taskService.UpdateTask(taskUUID, newTaskDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// ChangeTaskColumn
// @Summary update task
// @Tags Task
// @Description update task
// @Produce json
// @Param columnId path string true "ColumnData ID"
// @Param taskId path string true "Task ID"
// @Success 200
// @Router       /task/column/{columnId}/task/{taskId} [put]
func (tc *TaskController) ChangeTaskColumn(c *gin.Context) {
	taskId := c.Param("taskId")
	taskUUID, err1 := uuid.Parse(taskId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	columnId := c.Param("columnId")
	columnUUID, err1 := uuid.Parse(columnId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := tc.taskService.ChangeTaskColumn(taskUUID, columnUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// FinishTask
// @Summary update task
// @Tags Task
// @Description update task
// @Produce json
// @Param taskId path string true "Task ID"
// @Success 200
// @Router       /task/{taskId}/complete [put]
func (tc *TaskController) FinishTask(c *gin.Context) {
	taskId := c.Param("taskId")
	taskUUID, err1 := uuid.Parse(taskId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := tc.taskService.CompleteTask(taskUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}
