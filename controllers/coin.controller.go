package controllers

import (
	"fmt"
	"net/http"

	"example.com/vote-system-api/models"
	"example.com/vote-system-api/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CoinController struct {
	CoinService services.CoinService
}

func NewCoinController(coinservice services.CoinService) CoinController {
	return CoinController{
		CoinService: coinservice,
	}
}

func (cc *CoinController) CreateCoin(ctx *gin.Context) {
	var coin models.Coin
	if err := ctx.ShouldBindJSON(&coin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := cc.CoinService.CreateCoin(&coin)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Coin created"})
}

func (cc *CoinController) GetAllVotes(ctx *gin.Context) {
	votes, err := cc.CoinService.GetAllVotes()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, votes)
}

func (cc *CoinController) HandleVote(ctx *gin.Context) {
	idParam := ctx.Param("id")
	coinId, _ := uuid.FromString(idParam)
	var body models.VoteBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	user, _ := ctx.Get("user")
	fmt.Println("user:", user)
	votes, err := cc.CoinService.HandleVote(coinId, "f39ff674-c4fd-4d5a-9514-b46bb2f9c63f", body.Vote)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "something went wrong"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": votes})
}

func (cc *CoinController) RegisterCoinRoutes(rg *gin.RouterGroup) {
	coinroute := rg.Group("/coin")
	coinroute.GET("/", cc.GetAllVotes)
	coinroute.POST("/", cc.CreateCoin)
	coinroute.POST("/vote/:id", cc.HandleVote)
}
