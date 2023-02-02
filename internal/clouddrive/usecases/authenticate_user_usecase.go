package usecases

import (
	"github.com/italoservio/clouddrive/internal/clouddrive/db"
	"github.com/italoservio/clouddrive/internal/clouddrive/entities"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func AuthenticateUser(payload dtos.DTOAuthenticateReq) (*dtos.DTOAuthenticateUserRes, error) {
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
