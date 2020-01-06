package api

import (
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
	"github.com/jinzhu/gorm"
)

// TeamResource onject
type TeamResource struct {
	Db *gorm.DB
}

// NewTeamService creates a new team service
func NewTeamService(prefix string) *restful.WebService {
	prefix = prepPrefix(prefix)
	return nil
}

// Register registers the API on the given container the API
func (r TeamResource) Register(prefix string, container *restful.Container) {
	r.Db.AutoMigrate(&TeamDb{})
	ws := new(restful.WebService)
	ws.
		Path(prefix).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion(APIVersion)

	tags := []string{"teams"}
	ws.Route(ws.GET("").To(r.getAll).
		Doc("Get all teams").
		Metadata(openapi.KeyOpenAPITags, tags).
		Writes([]Team{}).
		Returns(200, "OK", []Team{}),
	)
	ws.Route(ws.GET("/{id}").To(r.get).
		Doc("Get a specific team").
		Param(ws.PathParameter("id", "team id").DataType("string")).
		Metadata(openapi.KeyOpenAPITags, tags).
		Writes(Team{}).
		Returns(200, "OK", Team{}).
		Returns(404, "Not Found", nil),
	)
	ws.Route(ws.POST("/").To(r.add).
		Doc("Create a team").
		Metadata(openapi.KeyOpenAPITags, tags).
		Reads(Team{}).
		Returns(201, "Created", Team{}),
	)
	ws.Route(ws.DELETE("/{id}").To(r.delete).
		Doc("Delete a team").
		Param(ws.PathParameter("id", "team id").DataType("string")).
		Metadata(openapi.KeyOpenAPITags, tags).
		Returns(200, "OK", nil).
		Returns(404, "Not Found", nil),
	)

	container.Add(ws)
}

func (r TeamResource) get(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var ms TeamDb
	if err := r.Db.Where("ID = ?", id).First(&ms).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, fmt.Sprintf("Team '%s' not found", id))
	} else {
		response.WriteEntity(ms.Team)
	}
}

func (r TeamResource) getAll(request *restful.Request, response *restful.Response) {
	var msdbo []TeamDb

	if err := r.Db.Find(&msdbo).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, "Could not fetch teams")
	} else {
		mso := make([]Team, len(msdbo))
		for i, ob := range msdbo {
			mso[i] = ob.Team
		}
		response.WriteEntity(mso)
	}
}

func (r TeamResource) add(request *restful.Request, response *restful.Response) {
	team := TeamDb{}
	err := request.ReadEntity(&(team.Team))
	if err == nil {
		r.Db.Create(&team)
		response.WriteHeaderAndEntity(http.StatusCreated, team.Team)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (r TeamResource) delete(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var ms TeamDb
	if err := r.Db.Where("ID = ?", id).First(&ms).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, fmt.Sprintf("Team '%s' not found", id))
	} else if err := r.Db.Delete(&ms).Error; err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

}
