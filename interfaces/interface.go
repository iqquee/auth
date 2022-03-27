package interfaces

import (
	"github.com/hisyntax/auth/database"
	"github.com/hisyntax/auth/models"
)

//create a new user into the mysql datase
func MysqlCreateUser(user *models.User) (err error) {
	//check if the user email is already taken
	if err := database.MysqlDB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//gets a user infomation from the mysql database by the user email
func GetUserByEmail(user *models.User, email string) (*models.User, error) {
	if err := database.MysqlDB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
