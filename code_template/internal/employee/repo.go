package employee

import (
	"context"
	"go-tutorial/code_template/internal"
	"go-tutorial/code_template/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	cv   *internal.Configs
	repo *mongo.Database
}

func NewRepo(cv *internal.Configs) *Repo {
	return &Repo{
		cv:   cv,
		repo: cv.MongoDB.Client.Database(cv.MongoDB.Database),
	}
}

func (r Repo) GetEmployeeById(ctx context.Context, employeeId string) ([]models.Employee, error) {

	filter := bson.M{
		"emp_id": employeeId,
	}

	coll := r.repo.Collection("sample_employee")
	cur, err := coll.Find(ctx, filter, options.Find().SetLimit(10))
	if err != nil {
		return nil, err
	}

	var rtn []models.Employee
	if err := cur.All(ctx, &rtn); err != nil {
		return nil, err
	}
	return rtn, nil
}
