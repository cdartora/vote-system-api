package services

import (
	"example.com/vote-system-api/models"
	uuid "github.com/satori/go.uuid"
)

type CoinService interface {
	CreateCoin(*models.Coin) error
	// GetVotes(*uuid.UUID) (*models.CoinVotes, error)
	GetAllVotes() ([]*models.CoinVotes, error)
	// DeleteCoin(*uuid.UUID) error
	HandleVote(coinId uuid.UUID, userId string, vote int) (*int, error)
}
