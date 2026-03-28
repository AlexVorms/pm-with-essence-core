package controllers

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/boardDTO"
	"pm-with-essence/cmd/service/boardService"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BoardController struct {
	boardService *boardService.BoardService
}

func NewBoardController(boardService *boardService.BoardService) *BoardController {
	return &BoardController{
		boardService: boardService,
	}
}

// CreateBoard
// @Summary create new board
// @Tags Board
// @Description create new board
// @Produce json
// @Param projectId path string true "Project ID"
// @Param BoardDTO body boardDTO.BoardDTO true "BoardDTO"
// @Success 200
// @Router       /board/{projectID} [post]
func (b *BoardController) CreateBoard(c *gin.Context) {
	projectId := c.Param("project_id")
	parsedUUID, err1 := uuid.Parse(projectId)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	var BoardDTO boardDTO.BoardDTO
	if err := c.ShouldBindJSON(&BoardDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := b.boardService.CreateBoard(parsedUUID, BoardDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// DeleteBoard
// @Summary create new board
// @Tags Board
// @Description create new board
// @Produce json
// @Param boardID path string true "Board ID"
// @Success 200
// @Router       /board/{boardID} [delete]
func (b *BoardController) DeleteBoard(c *gin.Context) {
	boardID := c.Param("boardID")
	parsedUUID, err1 := uuid.Parse(boardID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	err := b.boardService.DeleteBoard(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

// GetBoard
// @Summary get board
// @Tags Board
// @Description get board
// @Produce json
// @Param boardID path string true "Board ID"
// @Success 200
// @Router       /board/{boardID} [get]
func (b *BoardController) GetBoard(c *gin.Context) {
	boardID := c.Param("boardID")
	parsedUUID, err1 := uuid.Parse(boardID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	board, err := b.boardService.GetBoard(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"board": board})
}

// GetBoards
// @Summary get board
// @Tags Board
// @Description get board
// @Produce json
// @Param boardID path string true "Board ID"
// @Success 200
// @Router       /board/{boardID} [put]
func (b *BoardController) GetBoards(c *gin.Context) {
	boardID := c.Param("boardID")
	parsedUUID, err1 := uuid.Parse(boardID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	boards, err := b.boardService.GetAllBoards(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"boards": boards})
}

// UpdateBoard
// @Summary update board
// @Tags Board
// @Description update board
// @Produce json
// @Param projectId path string true "Project ID"
// @Param BoardDTO body boardDTO.BoardDTO true "BoardDTO"
// @Success 200
// @Router       /board/{projectID} [post]
func (b *BoardController) UpdateBoard(c *gin.Context) {
	boardID := c.Param("boardID")
	parsedUUID, err1 := uuid.Parse(boardID)
	if err1 != nil {
		fmt.Println("Ошибка парсинга UUID: ", err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	var BoardDTO boardDTO.BoardDTO
	if err := c.ShouldBindJSON(&BoardDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := b.boardService.UpdateBoard(parsedUUID, BoardDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
