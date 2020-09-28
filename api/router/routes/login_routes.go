package routes

import (
	"github.com/akwanmaroso/blogos/api/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		Uri: "/login",
		Method: http.MethodPost,
		Handler: controllers.Login,
		AuthRequired: false,
	},
}
