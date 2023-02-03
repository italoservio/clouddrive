package usecases

import (
	"log"

	"github.com/italoservio/clouddrive/internal/clouddrive/repositories"
	"github.com/italoservio/clouddrive/internal/clouddrive/usecases/dtos"
)

func GetUser(payload dtos.DTOGetUserReq) (*dtos.DTOGetUserRes, error) {
	var document *dtos.DTOGetUserRes

	document, err := repositories.GetUserRepo(payload)
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
