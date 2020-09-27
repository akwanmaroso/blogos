package auto

import "github.com/akwanmaroso/blogos/api/models"

var users = []models.User{
	models.User{
		Nickname: "Jhon Doe",
		Email:    "jhondoe@email.com",
		Password: "12345689",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title",
		Content: "Content",
	},
}
