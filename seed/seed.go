package seed

import (
	"context"
	"proximityservice/models"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateBusiness(b models.Business, bC *mongo.Collection, wg *sync.WaitGroup) error {
	defer wg.Done()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := bC.InsertOne(ctx, b)
	if err != nil {
		return err
	}
	return nil
}
