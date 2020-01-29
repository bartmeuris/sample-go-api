package api

import (
	"github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
	"github.com/jinzhu/gorm"
)

// APIVersion is the API version constant
const APIVersion = "v1"

// Service is an interface to define an internal API resource
type Service interface {
	Register(prefix string, container *restful.Container)
}

func prepPrefix(pre string) string {
	return pre
}

// RegisterAll registers all api endpoints
func RegisterAll(db *gorm.DB) (*restful.Container, error) {
	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})

	//team := TeamAPI{Db: db}
	//team.Register("/api/teams", container)
	NewTeamAPI(db).Register("/api/teams", container)

	config := openapi.Config{
		WebServices: container.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/api/apidocs.json",
	}
	container.Add(openapi.NewOpenAPIService(config))

	return container, nil
}
