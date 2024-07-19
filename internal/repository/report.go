package repository

import (
	"context"
	"fmt"
	"go_mongo_db/internal/entities"
	"go_mongo_db/internal/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportRepository struct{}

type ReportParams struct {
	Id          primitive.ObjectID `bson:"_id"`
	Responsible string             `bson:"Responsible"`
}

const reportCollection string = "Report"

type IReportRepository interface {
	Create(report entities.Report) error
	Find(filter *ReportParams) (*[]entities.Report, error)
	Update(id primitive.ObjectID, value entities.Report) error
	Delete(id primitive.ObjectID) error
	handleFilter(params *ReportParams) bson.M
	handleSet(params *entities.Report) bson.M
}

func (r *ReportRepository) Create(report entities.Report) error {
	database := &services.MongoOperations{}
	ctx := context.TODO()

	report.Id = primitive.NewObjectID()

	_, err := database.Insert(ctx, reportCollection, report)

	if err != nil {
		fmt.Println("INSERT ERROR")
	}

	return err
}

func (r *ReportRepository) Find(params *ReportParams) (*[]entities.Report, error) {
	database := &services.MongoOperations{}
	ctx := context.TODO()

	values := []entities.Report{}

	filter := r.handleFilter(params)

	err := database.Find(ctx, reportCollection, filter, &values)

	if err != nil {
		fmt.Println("FIND ERROR")
	}

	return &values, err
}

func (r *ReportRepository) Update(id primitive.ObjectID, value entities.Report) error {
	database := &services.MongoOperations{}
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	set := r.handleSet(&value)

	_, err := database.Update(ctx, reportCollection, filter, set)

	if err != nil {
		fmt.Println("UPDATE ERROR: ", err)
	}

	return err
}

func (r *ReportRepository) Delete(id primitive.ObjectID) error {
	database := &services.MongoOperations{}
	ctx := context.TODO()

	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	_, err := database.Delete(ctx, reportCollection, filter)

	if err != nil {
		fmt.Println("UPDATE ERROR")
	}

	return err
}

func (r *ReportRepository) handleFilter(params *ReportParams) bson.M {
	filter := bson.M{}

	if params != nil {

		if params.Id != primitive.NilObjectID {
			filter["_id"] = params.Id
		}

		if params.Responsible != "" {
			filter["Responsible"] = params.Responsible
		}
	}

	return filter
}

func (r *ReportRepository) handleSet(params *entities.Report) bson.M {
	update := bson.M{}

	if params != nil {

		set := bson.M{}

		if params.Responsible != "" {
			set["Responsible"] = params.Responsible
		}

		if params.Title != "" {
			set["Title"] = params.Title
		}

		if params.Content != "" {
			set["Content"] = params.Content
		}

		if len(set) > 0 {
			update["$set"] = set
		}
	}

	return update
}
