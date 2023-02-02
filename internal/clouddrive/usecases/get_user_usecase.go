package usecases

import (
	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func GetUser(payload dtos.DTOGetUserReq) (*dtos.DTOGetUserRes, error) {
	coll := db.UsersDatabase.Collection(db.USERS_COLLECTION)
	ctx, cancel := db.Context()
	defer cancel()

	var document entities.User

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
