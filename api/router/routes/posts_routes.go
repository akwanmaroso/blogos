package routes

import (
	"github.com/akwanmaroso/blogos/api/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		Uri: "/posts",
		Method: http.MethodGet,
		Handler: controllers.GetPosts,
	},
	{
		Uri: "/posts",
		Method: http.MethodPost,
		Handler: controllers.CreatePost,
	},
}
