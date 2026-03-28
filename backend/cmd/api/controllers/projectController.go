package controllers

import (
	"fmt"
	"log"
	"net/http"
	projectDTO "pm-with-essence/cmd/api/model/projectDTO"
	"pm-with-essence/cmd/service/projectService"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	projectService *projectService.ProjectService
}

func NewProjectController(projectService *projectService.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}

// CreateProject
// @Summary create new project
// @Tags Project
// @Description create new project
// @Produce json
// @Param newProjectDTO body projectDTO.ProjectDTO true "newProjectDTO"
// @Success 200
// @Router       /project [post]
func (ps *ProjectController) CreateProject(c *gin.Context) {
	var newProjectDTO projectDTO.ProjectDTO
	if err := c.ShouldBindJSON(&newProjectDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ps.projectService.CreateProject(newProjectDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
	return
}

// DeleteProject
// @Summary delete project
// @Tags Project
// @Description delete project
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200
// @Router       /project/{projectId} [delete]
func (ps *ProjectController) DeleteProject(c *gin.Context) {
	project_id := c.Param("project_id")
	parsedUUID, err1 := uuid.Parse(project_id)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := ps.projectService.DeleteProject(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// UpdateProject
// @Summary update project
// @Tags Project
// @Description update project
// @Produce json
// @Param project_id path string true "Project ID"
// @Param newProjectDTO body projectDTO.ProjectDTO true "newProjectDTO"
// @Success 200
// @Router       /project/{projectId} [update]
func (ps *ProjectController) UpdateProject(c *gin.Context) {
	project_id := c.Param("project_id")
	parsedUUID, err := uuid.Parse(project_id)
	if err != nil {
		fmt.Println("Ошибка парсинга UUID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var newProjectDTO projectDTO.ProjectDTO
	if err = c.ShouldBindJSON(&newProjectDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err1 := ps.projectService.UpdateProject(parsedUUID, newProjectDTO)
	if err1 != nil {
		fmt.Println("error occurred: " + err1.Error())
		c.JSON(err1.Code, err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (ps *ProjectController) GetAllProjects(c *gin.Context) {

}

// GetProject
// @Summary get project
// @Tags Project
// @Description update project
// @Produce json
// @Param project_id path string true "Project ID"
// @Success 200
// @Router       /project/{projectId} [get]
func (ps *ProjectController) GetProject(c *gin.Context) {
	projectId := c.Param("projectId")
	parsedUUID, err := uuid.Parse(projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err1 := ps.projectService.GetProject(parsedUUID)
	if err1 != nil {
		fmt.Println("error occurred: " + err1.Error())
		c.JSON(err1.Code, err1)
		return
	}
	c.JSON(http.StatusOK, gin.H{"project": project})
}
