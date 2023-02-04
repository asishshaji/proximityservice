package main

import (
	"fmt"
	"log"
	"math/rand"
	"proximityservice/models"
	"proximityservice/seed"
	"proximityservice/utils"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	client, err := utils.ConnectDB("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")
	if err != nil {
		log.Fatalln(err)
	}
	businessCollection := client.Database("proximity").Collection("business")
	wg := new(sync.WaitGroup)
	start := time.Now()
	for i := 0; i < 120000; i++ {
		objID := primitive.NewObjectIDFromTimestamp(time.Now())
		r := 0 + rand.Float64()*(3)

		b := models.Business{
			ID:        objID,
			City:      "City-" + fmt.Sprint(i),
			State:     "State-" + fmt.Sprint(i),
			Country:   "Country-" + fmt.Sprint(i),
			Latitude:  8.524139 + r,
			Longitude: 76.936638 + r,
		}
		wg.Add(1)
		go seed.CreateBusiness(b, businessCollection, wg)
	}
	wg.Wait()
	fmt.Println("Time : ", time.Since(start))

}
