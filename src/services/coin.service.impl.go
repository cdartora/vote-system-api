package services

import (
	"context"
	"errors"
	"fmt"

	"example.com/vote-system-api/src/models"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CoinServiceImpl struct {
	coincollection *mongo.Collection
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewCoinService(coincollection *mongo.Collection, usercollection *mongo.Collection, ctx context.Context) CoinService {
	return &CoinServiceImpl{
		coincollection: coincollection,
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (c *CoinServiceImpl) resetVotes() ([]models.Vote, error) {
	var votes []models.Vote

	cursor, err := c.usercollection.Find(c.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next((c.ctx)) {
		var user models.User
		var vote models.Vote
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		vote.UserId = user.Id
		vote.Vote = 0
		votes = append(votes, vote)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(votes) == 0 {
		return nil, errors.New(("users not found"))
	}

	return votes, err
}

func (c *CoinServiceImpl) CreateCoin(coin *models.Coin) error {
	newId := uuid.NewV4()
	coin.Id = newId
	votes, err := c.resetVotes()
	if err != nil {
		return err
	}
	coin.Votes = votes
	_, err = c.coincollection.InsertOne(c.ctx, coin)
	return err
}

// func (c *CoinServiceImpl) GetVotes(*uuid.UUID) (*models.CoinVotes, error) {

// }

func (c *CoinServiceImpl) GetAllVotes() ([]*models.CoinVotes, error) {
	var coins []*models.CoinVotes

	cursor, err := c.coincollection.Find(c.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next((c.ctx)) {
		var coin models.Coin
		var coinVotes models.CoinVotes
		err := cursor.Decode(&coin)
		if err != nil {
			return nil, err
		}
		fmt.Println("coinId: ", coin.Id)
		numberOfVotes, err := c.getNumberOfVotes(coin.Id)
		if err != nil {
			fmt.Println("error: ", err)
			return nil, err
		}
		coinVotes.Votes = *numberOfVotes
		coins = append(coins, &coinVotes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(coins) == 0 {
		return nil, errors.New(("documents not found"))
	}

	return coins, nil
}

// func (c *CoinServiceImpl) DeleteCoin(*uuid.UUID) error

func (c *CoinServiceImpl) getNumberOfVotes(coinId uuid.UUID) (*int, error) {
	var votes = 0
	var coin models.Coin

	err := c.coincollection.FindOne(c.ctx, bson.D{bson.E{"_id", coinId}}).Decode(coin)
	if err != nil {
		return nil, err
	}

	for _, s := range coin.Votes {
		votes += s.Vote
	}
	fmt.Println("votes: ", votes)
	return &votes, err
}

func (c *CoinServiceImpl) HandleVote(coinId uuid.UUID, userId string, vote int) (*int, error) {
	filter := bson.D{
		bson.E{"_id", coinId},
		bson.E{"votes", bson.D{
			bson.E{
				"$elemMatch", bson.E{
					"user_id", userId,
				},
			},
		}},
	}
	query := bson.D{
		bson.E{"$set", bson.D{
			bson.E{"votes.$.vote", 1},
		}},
	}
	result, _ := c.coincollection.UpdateOne(c.ctx, filter, query)

	if result.MatchedCount != 1 {
		return nil, errors.New("no matched document found for update")
	}

	numberOfVotes, err := c.getNumberOfVotes(coinId)

	if err != nil {
		return nil, err
	}

	return numberOfVotes, nil
}
