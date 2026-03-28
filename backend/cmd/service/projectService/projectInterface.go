package projectService

import "github.com/google/uuid"

type IProjectService interface {
	IsProjectExist(uuid uuid.UUID) (bool, error)
}
