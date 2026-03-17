package controllers

import "pm-with-essence/cmd/service/projectService"

type ProjectController struct {
	projectService *projectService.ProjectService
}

func NewProjectController(projectService *projectService.ProjectService) *ProjectController {
	return &ProjectController{
		projectService: projectService,
	}
}
