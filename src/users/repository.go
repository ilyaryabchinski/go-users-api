package users

import (
	"ilyaryabchinski/gotask/src/db"
	"ilyaryabchinski/gotask/src/utils"
	"strconv"
)

func GetAll() ([]User, error) {
	database := db.GetDatabaseInstance()
	iter := database.NewIterator(nil, nil)

	result := []User{}
	for iter.Next() {
		user := User{}
		err := utils.DecodeFromJsonString(iter.Value(), &user)
		if err != nil {
			return result, err
		}
		result = append(result, user)
	}

	return result, nil

}

func GetOne(personalCode uint64) (User, error) {
	database := db.GetDatabaseInstance()

	user := User{}

	key := strconv.FormatUint(personalCode, 10)
	value, err := database.Get([]byte(key), nil)

	if err != nil {
		return user, err
	}
	error := utils.DecodeFromJsonString(value, &user)
	return user, error

}

func Create(user User) error {
	database := db.GetDatabaseInstance()

	user.PersonalCode = utils.GetRandomUint64()

	encoded, err := utils.EncodeToJsonString(user)
	if err != nil {
		return err
	}

	key := strconv.FormatUint(user.PersonalCode, 10)
	error := database.Put([]byte(key), encoded, nil)

	return error

}

func Edit(personalCode uint64, user User) error {
	database := db.GetDatabaseInstance()

	encoded, err := utils.EncodeToJsonString(user)
	if err != nil {
		return err
	}

	key := strconv.FormatUint(personalCode, 10)
	error := database.Put([]byte(key), encoded, nil)

	return error
}

func Delete(personalCode uint64) error {
	database := db.GetDatabaseInstance()

	key := strconv.FormatUint(personalCode, 10)
	error := database.Delete([]byte(key), nil)

	return error
}
