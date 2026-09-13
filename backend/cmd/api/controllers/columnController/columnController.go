package columnController

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/columnDTO"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/internal/api/app_errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ColumnController struct {
	columnService ColumnService
}
type ColumnService interface {
	CreateColumn(boardID uuid.UUID, model columnDTO.ColumnDTO) *app_errors.HttpError
	DeleteColumn(columnID uuid.UUID) *app_errors.HttpError
	UpdateColumn(columnID uuid.UUID, model columnDTO.ColumnUpdateDTO) *app_errors.HttpError
	ChangeColumnOrder(model columnDTO.UpdateColumnOrderDTO) *app_errors.HttpError
}

func NewColumnController(columnService ColumnService) *ColumnController {
	return &ColumnController{
		columnService: columnService,
	}
}

// CreateColumn
// @Summary create new column
// @Tags ColumnData
// @Description create new board
// @Produce json
// @Param boardID path string true "boardID"
// @Param ColumnDTO body columnDTO.ColumnDTO true "ColumnDTO"
// @Success 200
// @Router       /column/{boardID} [post]
func (p *ColumnController) CreateColumn(c *gin.Context) {
	boardID := c.Param("boardID")
	parsedUUID, err1 := uuid.Parse(boardID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	var ColumnDTO columnDTO.ColumnDTO
	if err := c.ShouldBindJSON(&ColumnDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p.columnService.CreateColumn(parsedUUID, ColumnDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// UpdateColumn
// @Summary create new column
// @Tags ColumnData
// @Description create new board
// @Produce json
// @Param columnID path string true "columnID"
// @Param ColumnUpdateDTO body columnDTO.ColumnUpdateDTO true "ColumnUpdateDTO"
// @Success 200
// @Router       /column/{columnID} [put]
func (p *ColumnController) UpdateColumn(c *gin.Context) {
	columnID := c.Param("columnID")
	parsedUUID, err1 := uuid.Parse(columnID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	var ColumnUpdateDTO columnDTO.ColumnUpdateDTO
	if err := c.ShouldBindJSON(&ColumnUpdateDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p.columnService.UpdateColumn(parsedUUID, ColumnUpdateDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// DeleteColumn
// @Summary delete column
// @Tags ColumnData
// @Description delete column
// @Produce json
// @Param columnID path string true "columnID"
// @Success 200
// @Router       /column/{columnID} [delete]
func (p *ColumnController) DeleteColumn(c *gin.Context) {
	columnID := c.Param("columnID")
	parsedUUID, err1 := uuid.Parse(columnID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := p.columnService.DeleteColumn(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}

// UpdateColumnOrder
// @Summary update column order
// @Tags ColumnData
// @Description update column order
// @Produce json
// @Param updateColumnOrderDTO body columnDTO.UpdateColumnOrderDTO true "UpdateColumnOrderDTO"
// @Success 200
// @Router       /column [post]
func (p *ColumnController) UpdateColumnOrder(c *gin.Context) {
	var updateColumnOrderDTO columnDTO.UpdateColumnOrderDTO
	if err := c.ShouldBindJSON(&updateColumnOrderDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := p.columnService.ChangeColumnOrder(updateColumnOrderDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "success",
	})
}
