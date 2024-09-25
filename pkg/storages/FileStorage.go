package storages

import (
	"fmt"
	"shell-backend/pkg/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

const (
	FILELIST_TABLE = "filelist"
)

type FileStorage struct {
	Storage *SqlLiteStorage
}

func (s *FileStorage) GetAllByUserId(userId string) ([]models.FileModel, error) {
	rows, err := s.Storage.DB.Query("SELECT * FROM filelist WHERE user_id=?", userId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []models.FileModel{}

	for rows.Next() {
		var item models.FileModel

		err := rows.Scan(
			&item.Id,
			&item.Table,
			&item.Url,
			&item.UserId,
			&item.Records,
			&item.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		result = append(result, item)
	}

	return result, nil
}

func (s *FileStorage) CreateFile(item models.FileModel) error {
	sql := `INSERT INTO filelist('table', 'url', 'user_id', 'records', 'created_at') VALUES(?,?,?,?,?)`

	statement, err := s.Storage.DB.Prepare(sql)

	if err != nil {
		return errors.Wrap(err, "prepare sql")
	}

	_, err = statement.Exec(item.Table, item.Url, item.UserId, item.Records, item.CreatedAt)

	if err != nil {
		return errors.Wrap(err, "exec sql")
	}

	return nil
}

func (s *FileStorage) RemoveFile(itemId int) error {
	sql := `DELETE FROM filelist WHERE id=?`

	statement, err := s.Storage.DB.Prepare(sql)

	if err != nil {
		return errors.Wrap(err, "prepare sql")
	}

	_, err = statement.Exec(itemId)

	if err != nil {
		return errors.Wrap(err, "exec sql")
	}

	return nil
}

func (s *FileStorage) GetFiles() ([]models.FileModel, error) {
	rows, err := s.Storage.DB.Query("SELECT * FROM filelist")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := []models.FileModel{}

	for rows.Next() {
		var item models.FileModel

		err := rows.Scan(
			&item.Id,
			&item.Table,
			&item.Url,
			&item.UserId,
			&item.Records,
			&item.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		result = append(result, item)
	}

	return result, nil
}

func (s *FileStorage) CreateTable() error {
	createSql := `CREATE TABLE IF NOT EXISTS filelist (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"table" TEXT NOT NULL UNIQUE,
		"url" TEXT NOT NULL,
		"user_id" TEXT NOT NULL,
		"records" INTEGER,
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
