package database

import (
	_const "github.com/bilalkocoglu/file-service/pkg/const"
	"github.com/pkg/errors"
	"time"
)

type ApplicationUser struct {
	ID        				uint       `gorm:"primarykey" json:"id"`
	CreatedAt 				time.Time  `json:"createdAt"`
	UpdatedAt 				time.Time  `json:"updatedAt"`
	Username  				string     `json:"username"`
	Password  				string     `json:"password"`
	CompanyName      		string     `json:"companyName"`
	Email     				string     `json:"email"`
}

func CreateDefaultUser() {
	var appUser ApplicationUser

	err := GetUserByUsername(&appUser, _const.Username)

	if err != nil {
		panic(err)
	}

	if appUser.ID == 0 {
		appUser = ApplicationUser{
			Username: 			_const.Username,
			Password: 			_const.Password,
			CompanyName:     	_const.CompanyName,
			Email:    			_const.Email,
		}

		err := SaveUser(&appUser)

		if err != nil {
			errors.Wrap(err, "Can not create default appUser")
		}
	}
}

func SaveUser(appUser *ApplicationUser) (err error) {
	if err = DB.Create(appUser).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(appUser *ApplicationUser, username string) (err error) {
	DB.Where("username = ?", username).First(appUser)

	return nil
}

func GetUserById(appUser *ApplicationUser, id uint64) (err error) {
	DB.Where("id = ?", id).First(appUser)

	return nil
}

func GetAllAppUsers(appUsers *[]ApplicationUser) {
	DB.Find(appUsers)
}
