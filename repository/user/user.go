package user

import (
	"vidlearn-final-projcect/business/user"
	"vidlearn-final-projcect/util"

	"gorm.io/gorm"
)

func UserRepository(dbCon *util.DatabaseConnection) user.Repository {
	var userRepository user.Repository
	userRepository = CreateMySQlRepository(dbCon.MySQlDB)

	return userRepository
}

type MySQLRepository struct {
	db *gorm.DB
}

func CreateMySQlRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repository *MySQLRepository) GetAllUser() (users []*user.User, err error) {
	err = repository.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *MySQLRepository) GetUserByID(ID int) (user *user.User, err error) {
	err = repository.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) GetUserByEmail(email string) (user *user.User, err error) {
	err = repository.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) GetUserByVerifyCode(verifyCode string) (user *user.User, err error) {
	err = repository.db.Where("verify_code=?", verifyCode).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) CreateUser(user *user.User) (*user.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *MySQLRepository) UpdateUser(userCurrent *user.User, IDCurrent int) (*user.User, error) {
	updateUser := map[string]interface{}{
		"ID":         IDCurrent,
		"Name":       userCurrent.Name,
		"Email":      userCurrent.Email,
		"Password":   userCurrent.Password,
		"Role":       userCurrent.Role,
		"IsReset":    userCurrent.IsReset,
		"VerifyCode": userCurrent.VerifyCode,
		"VerifiedAt": userCurrent.VerifiedAt,
		"CreatedAt":  userCurrent.CreatedAt,
		"UpdatedAt":  userCurrent.UpdatedAt,
	}

	err := repository.db.Model(&userCurrent).Where("id=?", IDCurrent).Updates(updateUser).Error
	if err != nil {
		return nil, err
	}

	return userCurrent, nil
}

func (repository *MySQLRepository) DeleteUser(ID int) (user *user.User, err error) {
	err = repository.db.First(&user, ID).Error
	if err != nil {
		return nil, err
	}

	err = repository.db.Delete(&user, ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
