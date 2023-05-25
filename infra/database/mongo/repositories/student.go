package repositories

import (
	"context"

	"github.com/brunoarruda04/golang-clean-architecture/entities"
	"github.com/brunoarruda04/golang-clean-architecture/infra/database/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongo_drive "go.mongodb.org/mongo-driver/mongo"
)

const (
	StudentCollection = "students"
)

type StudentRepository struct {
	Ctx        context.Context
	Collection mongo_drive.Collection
}

func NewStudentRepository(ctx context.Context, client *mongo_drive.Client) *StudentRepository {
	return &StudentRepository{
		Ctx:        ctx,
		Collection: *mongo.GetCollection(ctx, client, StudentCollection),
	}
}

func (sr StudentRepository) Create(student *entities.Student) error {
	_, err := sr.Collection.InsertOne(sr.Ctx, student)
	return err
}

func (sr StudentRepository) List() (students []entities.Student, err error) {
	cur, err := sr.Collection.Find(sr.Ctx, bson.M{})
	if err != nil {
		return students, err
	}

	for cur.Next(sr.Ctx) {
		var student entities.Student
		err = cur.Decode(&student)
		if err != nil {
			return students, err
		}
		students = append(students, student)
	}

	return students, nil
}

func (sr StudentRepository) FindById(id uuid.UUID) (student entities.Student, err error) {
	err = sr.Collection.FindOne(sr.Ctx, bson.M{"_id": id}).Decode(&student)
	return student, err
}

func (sr StudentRepository) Update(student *entities.Student) error {
	_, err := sr.Collection.UpdateOne(sr.Ctx, bson.M{"_id": student.ID}, bson.M{"$set": student})
	return err
}

func (sr StudentRepository) Delete(id uuid.UUID) error {
	_, err := sr.Collection.DeleteOne(sr.Ctx, bson.M{"_id": id})
	return err
}
