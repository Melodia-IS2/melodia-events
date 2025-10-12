package swagger

import (
	"net/http"

	"github.com/Melodia-IS2/melodia-go-utils/pkg/router"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerHandler struct {
}

func (handler *SwaggerHandler) Register(rt *router.Router) {
	rt.Get("/swagger/*", handler.getSwagger)
}

func (handler *SwaggerHandler) getSwagger(w http.ResponseWriter, r *http.Request) error {
	httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")).ServeHTTP(w, r)
	return nil
}
