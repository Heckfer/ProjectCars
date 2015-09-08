package handler
import (
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"log"

	"github.com/drborges/appx"
	"appengine/datastore"
)

func ListPostsHandler(r render.Render, appx *appx.Datastore) {

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: "",
		Data: &[]*model.Post{},
	}

	err := appx.Query(model.Posts.All()).Results(response.Data)

	if err != nil && err != datastore.Done {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = err.Error()
	}

	r.JSON(response.ErrorCode, response)
}
