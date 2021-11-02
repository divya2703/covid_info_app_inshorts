package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/divya2703/covid-tracker-rest-api/db"
	"github.com/divya2703/covid-tracker-rest-api/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type mongoRepo struct{}

var collection = db.ConnectDB()

func NewMongoRepository() Repository {

	return &mongoRepo{}
}

func (*mongoRepo) FindByName(stateName string) (*entity.StateReport, error) {
	var stateReport entity.StateReport
	filter := bson.M{"State": stateName}
	err := collection.FindOne(context.TODO(), filter).Decode(&stateReport)

	if err != nil {
		return &stateReport, err
	}
	return &stateReport, err
}

func (*mongoRepo) FindAll() ([]entity.StateReport, error) {

	var stateReports []entity.StateReport
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Println(err)
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var stateReport entity.StateReport

		// & character returns the memory address of the following variable.
		err := cur.Decode(&stateReport) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}
		// add item our array
		stateReports = append(stateReports, stateReport)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return stateReports, err
}
