package routes

import (
	"net/http"

	"github.com/akwanmaroso/blogos/api/controllers"
)

var postsRoutes = []Route{
	{
		Uri:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},
	{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
	{
		Uri:     "/posts/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetPost,
	},
	{
		Uri:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
	},
	{
		Uri:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePost,
	},
}
