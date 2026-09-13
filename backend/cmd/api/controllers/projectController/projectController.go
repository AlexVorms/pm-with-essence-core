package projectController

import (
	"fmt"
	"log"
	"net/http"
	projectDTO "pm-with-essence/cmd/api/model/projectDTO"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectController struct {
	projectService ProjectService
}
type ProjectService interface {
	CreateProject(model projectDTO.ProjectDTO) *app_errors.HttpError
	DeleteProject(projectId uuid.UUID) *app_errors.HttpError
	GetProject(projectId uuid.UUID) (*pm_entity.Project, *app_errors.HttpError)
	UpdateProject(projectID uuid.UUID, model projectDTO.ProjectDTO) *app_errors.HttpError
}

func NewProjectController(projectService ProjectService) *ProjectController {
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
// @Failure 401 {object} responseDTO.ErrorResponse
// @Failure 500 {object} responseDTO.ErrorResponse
// @Security BearerAuth
// @Router       /project [post]
func (ps *ProjectController) CreateProject(c *gin.Context) {
	userID := c.GetString("userID")
	log.Println("userID:", userID)

	var newProjectDTO projectDTO.ProjectDTO

	if err := c.ShouldBindJSON(&newProjectDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ps.projectService.CreateProject(newProjectDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// DeleteProject
// @Summary delete project
// @Tags Project
// @Description delete project
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200
// @Router       /project/{projectId} [delete]
func (ps *ProjectController) DeleteProject(c *gin.Context) {
	projectId := c.Param("projectId")
	parsedUUID, err1 := uuid.Parse(projectId)
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
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// UpdateProject
// @Summary update project
// @Tags Project
// @Description update project
// @Produce json
// @Param projectId path string true "Project ID"
// @Param newProjectDTO body projectDTO.ProjectDTO true "newProjectDTO"
// @Success 200
// @Router       /project/{projectId} [put]
func (ps *ProjectController) UpdateProject(c *gin.Context) {
	projectId := c.Param("projectId")
	parsedUUID, err := uuid.Parse(projectId)
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
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

func (ps *ProjectController) GetAllProjects(c *gin.Context) {

}

// GetProject
// @Summary get project
// @Tags Project
// @Description update project
// @Produce json
// @Param projectId path string true "Project ID"
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
	c.JSON(http.StatusOK, project)
}
