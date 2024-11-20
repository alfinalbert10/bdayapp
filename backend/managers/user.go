package managers

import (
	"bdayappbknd/common"
	"bdayappbknd/database"
	"bdayappbknd/models"
	"errors"
)

type UserManager struct {
	//db
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(userData *common.UserCreationInput) (*models.User, error) {
	newUser := &models.User{Username: userData.Username, Birthdate: userData.Birthdate}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("User creation failed")
	}

	return newUser, nil
}

func (userMgr *UserManager) List() ([]models.User, error) {
	users := []models.User{}
	database.DB.Find(&users)
	return users, nil
}

func (userMgr *UserManager) Get(id string) (models.User, error) {
	user := models.User{}
	database.DB.First(&user, id)
	return user, nil
}

func (userMgr *UserManager) Delete(UserId string) error {
	user := models.User{}
	database.DB.First(&user, UserId)
	database.DB.Delete(&user)
	return nil
}

func (userMgr *UserManager) Update(userId string, userData *common.UserUpdateInput) (*models.User, error) {

	user := models.User{}

	database.DB.First(&user, userId)
	//user.Username = userData.Username
	//user.Birthdate = userData.Birthdate
	//database.DB.Save(&user)
	database.DB.Model(&user).Updates(models.User{Username: userData.Username, Birthdate: userData.Birthdate})
	return &user, nil
}
