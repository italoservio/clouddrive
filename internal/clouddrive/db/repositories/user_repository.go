package repositories

import (
	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/dtos"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository() *UserRepository {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	return &UserRepository{
		coll: coll,
	}
}

func (usr *UserRepository) UserByEmail(email string) (*entities.User, error) {
	ctx, cancel := db.Context()
	defer cancel()

	var document *entities.User
	err := usr.coll.FindOne(
		ctx,
		bson.D{{Key: "email", Value: email}},
	).Decode(&document)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &entities.User{}, nil
		}
		return nil, err
	}

	return document, nil
}

func (usr *UserRepository) UserById(id string) (*entities.User, error) {
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = usr.coll.FindOne(
		ctx,
		bson.D{{Key: "_id", Value: objectId}},
	).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &entities.User{}, nil
		}
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

func (usr *UserRepository) CreateUser(
	payload dtos.DTOCreateUserReq,
) (*entities.User, error) {
	ctx, cancel := db.Context()
	defer cancel()

	res, err := usr.coll.InsertOne(ctx, payload)
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
