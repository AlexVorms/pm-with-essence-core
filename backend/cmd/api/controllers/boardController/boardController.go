package boardController

import (
	"fmt"
	"log"
	"net/http"
	"pm-with-essence/cmd/api/model/boardDTO"
	"pm-with-essence/cmd/api/model/responseDTO"
	"pm-with-essence/internal/api/app_errors"
	pm_entity "pm-with-essence/internal/domain/entity/pm-entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BoardController struct {
	boardService BoardService
}
type BoardService interface {
	CreateBoard(projectID uuid.UUID, model boardDTO.BoardDTO) *app_errors.HttpError
	DeleteBoard(boardID uuid.UUID) *app_errors.HttpError
	GetBoard(boardID uuid.UUID) (*boardDTO.GetBoardDTO, *app_errors.HttpError)
	GetAllBoards(projectID uuid.UUID) ([]*pm_entity.Board, *app_errors.HttpError)
	UpdateBoard(boardID uuid.UUID, model boardDTO.BoardDTO) *app_errors.HttpError
}

func NewBoardController(boardService BoardService) *BoardController {
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
// @Router       /board/{projectId} [post]
func (b *BoardController) CreateBoard(c *gin.Context) {
	projectId := c.Param("projectId")
	parsedUUID, err1 := uuid.Parse(projectId)
	if err1 != nil {
		fmt.Println("parsing UUID error: ", err1)
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Parsing UUID error",
		})
		return
	}
	var BoardDTO boardDTO.BoardDTO
	if err := c.ShouldBindJSON(&BoardDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}
	err := b.boardService.CreateBoard(parsedUUID, BoardDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "Success",
	})
}

// DeleteBoard
// @Summary delete board
// @Tags Board
// @Description delete board
// @Produce json
// @Param boardId path string true "Board ID"
// @Success 200
// @Router       /board/{boardId} [delete]
func (b *BoardController) DeleteBoard(c *gin.Context) {
	boardId := c.Param("boardId")
	parsedUUID, err1 := uuid.Parse(boardId)
	if err1 != nil {
		fmt.Println("parsing UUID error: ", err1)
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Parsing UUID error",
		})
		return
	}
	err := b.boardService.DeleteBoard(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "Success",
	})
}

// GetBoard
// @Summary get board
// @Tags Board
// @Description get board
// @Produce json
// @Param boardId path string true "Board ID"
// @Success 200
// @Router       /board/{boardId} [get]
func (b *BoardController) GetBoard(c *gin.Context) {
	boardId := c.Param("boardId")
	parsedUUID, err1 := uuid.Parse(boardId)
	if err1 != nil {
		fmt.Println("parsing UUID error: ", err1)
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Parsing UUID error",
		})
		return
	}
	board, err := b.boardService.GetBoard(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, board)
}

// GetBoards
// @Summary get all boards
// @Tags Board
// @Description get all boards
// @Produce json
// @Param projectId path string true "Project ID"
// @Success 200
// @Router       /board/all-boards/{projectId} [get]
func (b *BoardController) GetBoards(c *gin.Context) {
	projectId := c.Param("projectId")
	parsedUUID, err1 := uuid.Parse(projectId)
	if err1 != nil {
		fmt.Println("parsing UUID error: ", err1)
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Parsing UUID error",
		})
		return
	}
	boards, err := b.boardService.GetAllBoards(parsedUUID)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, boards)
}

// UpdateBoard
// @Summary update board
// @Tags Board
// @Description update board
// @Produce json
// @Param boardId path string true "Board ID"
// @Param BoardDTO body boardDTO.BoardDTO true "BoardDTO"
// @Success 200
// @Router       /board/{boardId} [post]
func (b *BoardController) UpdateBoard(c *gin.Context) {
	boardId := c.Param("boardId")
	parsedUUID, err1 := uuid.Parse(boardId)
	if err1 != nil {
		fmt.Println("parsing UUID error: ", err1)
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Parsing UUID error",
		})
		return
	}
	var BoardDTO boardDTO.BoardDTO
	if err := c.ShouldBindJSON(&BoardDTO); err != nil {
		log.Printf("error parsing json: " + err.Error())
		c.JSON(http.StatusBadRequest, responseDTO.ErrorResponse{
			Message: "Invalid request body",
		})
		return
	}
	err := b.boardService.UpdateBoard(parsedUUID, BoardDTO)
	if err != nil {
		fmt.Println("error occurred: " + err.Error())
		c.JSON(err.Code, responseDTO.ErrorResponse{
			Message: err.Message,
		})
		return
	}
	c.JSON(http.StatusOK, responseDTO.MessageResponse{
		Message: "Success",
	})
}
