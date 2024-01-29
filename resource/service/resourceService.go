package service

import (
	"net/http"

	"go-oauth2-server/resource/domain/entity"
	"go-oauth2-server/resource/util"
)

var resource = &entity.Resource{
    Name:        "resource",
    Description: "This is a resource",
}

func GetResource(r *http.Request) (*entity.Resource, error) {
    _, err := util.GetAccessToken(r)
    if err != nil {
        return nil, err
    }

    return resource, nil
}
