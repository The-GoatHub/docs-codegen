package api

import (
	"context"

	"github.com/google/uuid"
)

var defaultGoat = Goat{
	Id:   uuid.MustParse("3267b999-a3bc-482e-aabd-cb2b2b618cb5"),
	Name: "lovely goat",
	Age:  2,
}

func (s *Server) GetGoatByID(_ context.Context, req GetGoatByIDRequestObject) (GetGoatByIDResponseObject, error) {
	if req.Id == defaultGoat.Id {
		return GetGoatByID200JSONResponse{
			Id:   defaultGoat.Id,
			Name: defaultGoat.Name,
			Age:  defaultGoat.Age,
		}, nil
	}

	return GetGoatByID404JSONResponse{N404JSONResponse{Message: "Goat not found"}}, nil
}
