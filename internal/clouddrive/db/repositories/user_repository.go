package repositories

import (
	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserByEmail(email string) (*entities.User, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var user *entities.User
	err := coll.FindOne(ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserById(id string) (*entities.User, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = coll.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&document)
	if err != nil {
		return nil, err
	}

	return &entities.User{
		Id:        document.Id,
		FirstName: document.FirstName,
		LastName:  document.LastName,
		Email:     document.Email,
		Pass:      document.Pass,
	}, nil
}

func CreateUser(payload dtos.DTOCreateUserReq) (*entities.User, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	res, err := coll.InsertOne(ctx, payload)
	if err != nil {
		return nil, err
	}

	stringified_id := res.InsertedID.(primitive.ObjectID).Hex()
	return &entities.User{
		Id:        stringified_id,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Pass:      payload.Pass,
	}, nil
}
