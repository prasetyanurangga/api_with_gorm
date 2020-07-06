package repository

import (
	"api_with_gorm/model"
	"api_with_gorm/utils"
	"log"
)

// Output struct
type Output struct{
	Result interface{}
	Error error
}

// GetAll Func
func GetAll() Output{
	db, err := utils.GetConnect()

	if err != nil{
		log.Fatal(err)
	}
	var users model.Users

	er := db.Find(&users).Error
	if er != nil{
		return Output{
			Error : er,
		}

	}

	return Output{
		Result : users,
	}
}


// Save Func
func Save(user model.User) Output{
	db, err := utils.GetConnect()

	if err != nil{
		log.Fatal(err)
	}
	er := db.Create(&user).Error
	if er != nil{
		return Output{
			Error : er,
		}

	}

	return Output{
		Result : user,
	}
}

// Delete Func
func Delete(user model.User) Output{
	db, err := utils.GetConnect()

	if err != nil{
		log.Fatal(err)
	}
	er := db.Delete(&user).Error
	if er != nil{
		return Output{
			Error : er,
		}

	}

	return Output{
		Result : user,
	}
}

// Update Func
func Update(user model.User) Output{
	db, err := utils.GetConnect()

	if err != nil{
		log.Fatal(err)
	}
	er := db.Model(&user).Updates(&user).Error
	if er != nil{
		return Output{
			Error : er,
		}

	}

	return Output{
		Result : user,
	}
}
