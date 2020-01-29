package api

import (
	"fmt"
	"net/http"

	"github.com/bartmeuris/sample-go-api/models"
	"github.com/emicklei/go-restful"
	openapi "github.com/emicklei/go-restful-openapi"
	"github.com/jinzhu/gorm"
)

// TeamAPI onject
type TeamAPI struct {
	Db *gorm.DB
}

// NewTeamAPI returns a new TeamAPI object
func NewTeamAPI(db *gorm.DB) *TeamAPI {
	return &TeamAPI{Db: db}
}

// Register registers the API on the given container the API
func (r TeamAPI) Register(prefix string, container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path(prepPrefix(prefix)).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		ApiVersion(APIVersion)

	tags := []string{"teams"}
	ws.Route(ws.GET("").To(r.getAll).
		Doc("Get all teams").
		Metadata(openapi.KeyOpenAPITags, tags).
		Writes([]models.Team{}).
		Returns(200, "OK", []models.Team{}),
	)
	ws.Route(ws.GET("/{id}").To(r.get).
		Doc("Get a specific team").
		Param(ws.PathParameter("id", "team id").DataType("string")).
		Metadata(openapi.KeyOpenAPITags, tags).
		Writes(models.Team{}).
		Returns(200, "OK", models.Team{}).
		Returns(404, "Not Found", nil),
	)
	ws.Route(ws.POST("/").To(r.add).
		Doc("Create a team").
		Metadata(openapi.KeyOpenAPITags, tags).
		Reads(models.Team{}).
		Returns(201, "Created", models.Team{}),
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

func (r TeamAPI) get(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var ms models.TeamDb
	if err := r.Db.Where("ID = ?", id).First(&ms).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, fmt.Sprintf("Team '%s' not found", id))
	} else {
		response.WriteEntity(ms.Team)
	}
}

func (r TeamAPI) getAll(request *restful.Request, response *restful.Response) {
	var msdbo []models.TeamDb

	if err := r.Db.Find(&msdbo).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, "Could not fetch teams")
	} else {
		mso := make([]models.Team, len(msdbo))
		for i, ob := range msdbo {
			mso[i] = ob.Team
		}
		response.WriteEntity(mso)
	}
}

func (r TeamAPI) add(request *restful.Request, response *restful.Response) {
	team := models.TeamDb{}
	err := request.ReadEntity(&(team.Team))
	if err == nil {
		r.Db.Create(&team)
		response.WriteHeaderAndEntity(http.StatusCreated, team.Team)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (r TeamAPI) delete(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var ms models.TeamDb
	if err := r.Db.Where("ID = ?", id).First(&ms).Error; err != nil {
		response.WriteErrorString(http.StatusNotFound, fmt.Sprintf("Team '%s' not found", id))
	} else if err := r.Db.Delete(&ms).Error; err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

}
