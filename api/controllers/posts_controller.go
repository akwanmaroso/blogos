package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/akwanmaroso/blogos/api/databases"
	"github.com/akwanmaroso/blogos/api/helpers/responses"
	"github.com/akwanmaroso/blogos/api/models"
	"github.com/akwanmaroso/blogos/api/repository"
	"github.com/akwanmaroso/blogos/api/repository/crud"
	"io/ioutil"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	db, err := databases.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostCRUD(db)
	func(postsRepository repository.PostsRepository) {
		posts, err := postsRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, posts)
	}(repo)

}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := databases.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostCRUD(db)
	func(postsRepository repository.PostsRepository) {
		post, err := postsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)

}
