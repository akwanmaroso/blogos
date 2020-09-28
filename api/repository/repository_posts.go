package repository

import "github.com/akwanmaroso/blogos/api/models"

type PostsRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint64) (models.Post, error)
	Update(uint64, models.Post) (int64, error)
	Delete(postId uint64, userId uint32) (int64, error)
}
