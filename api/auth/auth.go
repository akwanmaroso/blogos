package auth

import (
	"github.com/akwanmaroso/blogos/api/databases"
	"github.com/akwanmaroso/blogos/api/models"
	"github.com/akwanmaroso/blogos/api/security"
	"github.com/akwanmaroso/blogos/utils/channels"
	"github.com/jinzhu/gorm"
)

func SigIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan <-bool) {
		defer close(ch)
		db, err = databases.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
