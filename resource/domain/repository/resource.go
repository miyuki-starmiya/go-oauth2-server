package repository

import (
	"github.com/miyuki-starmiya/go-oauth2-server/resource/domain/entity"
)

func GetResource() (*entity.Resource, error) {
	var resource = &entity.Resource{
		Name:        "resource",
		Description: "This is a resource",
	}

	return resource, nil
}
