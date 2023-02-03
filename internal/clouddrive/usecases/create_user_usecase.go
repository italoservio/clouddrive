package usecases

import (
	"log"

	"github.com/italoservio/clouddrive/internal/clouddrive/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
)

func CreateUser(payload dtos.DTOCreateUserReq) (*dtos.DTOCreateUserRes, error) {

	var document *dtos.DTOCreateUserRes

	document, err := repositories.CreateUser(payload)
	if err != nil {
		log.Fatal(err)
	}

	return &dtos.DTOCreateUserRes{
		Id:        document.Id,
		FirstName: document.FirstName,
		LastName:  document.LastName,
		Email:     document.Email,
		Pass:      document.Pass,
	}, nil
}
