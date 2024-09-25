package storages

import "shell-backend/pkg/models"

type UserStorage struct {
	Storage *SqlLiteStorage
}

func (s *UserStorage) CreateUser(user models.UserModel) error {
	return nil
}

func (s *UserStorage) RemoveUser(userId string) error {
	return nil
}

func (s *UserStorage) GetUser(userId string) (models.UserModel, error) {
	var user models.UserModel

	return user, nil
}

func (s *UserStorage) GetUsers() ([]models.UserModel, error) {
	return nil, nil
}

func (s *UserStorage) CreateTable() error {
	createSql := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL,
		"password" TEXT NOT NULL,
		"role" TEXT NOT NULL,
		"created_at" INTEGER
	  );`

	statement, err := s.Storage.DB.Prepare(createSql)

	if err != nil {
		return err
	}

	_, err = statement.Exec()

	if err != nil {
		return err
	}

	return nil
}
