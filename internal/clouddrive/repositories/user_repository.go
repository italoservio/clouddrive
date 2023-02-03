package repositories

import (
	"log"

	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthenticateRepo(payload dtos.DTOAuthenticateReq) (*dtos.DTOAuthenticateUserRes, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"email", payload.Email}},
				bson.D{{"pass", payload.Pass}},
			},
		},
	}

	err := coll.FindOne(ctx, filter).Decode(&document)
	if err != nil {
		log.Fatal(err)
	}

	return &dtos.DTOAuthenticateUserRes{
		Id:        document.Id,
		FirstName: document.FirstName,
		LastName:  document.LastName,
		Email:     document.Email,
		Pass:      document.Pass,
	}, nil
}

func GetUserRepo(payload dtos.DTOGetUserReq) (*dtos.DTOGetUserRes, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User

	log.Printf(payload.Email)
	err := coll.FindOne(ctx, bson.D{{"email", payload.Email}}).Decode(&document)
	if err != nil {
		log.Fatal(err)
	}

	return &dtos.DTOGetUserRes{
		Id:        document.Id,
		FirstName: document.FirstName,
		LastName:  document.LastName,
		Email:     document.Email,
		Pass:      document.Pass,
	}, nil
}

func CreateUser(payload dtos.DTOCreateUserReq) (*dtos.DTOCreateUserRes, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User

	err := coll.FindOne(ctx, bson.D{{"email", payload.Email}}).Decode(&document)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// erro de conflito
		}
		// erro de falha ao chamar db
	}

	res, err := coll.InsertOne(ctx, payload)
	if err != nil {
		// erro de falha ao chamar db
	}

	return &dtos.DTOCreateUserRes{
		Id:        res.InsertedID.(primitive.ObjectID),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Pass:      payload.Pass,
	}, nil
}
