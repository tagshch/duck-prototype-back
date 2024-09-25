package models

type FileModel struct {
	Id        int    `json:"id"`
	Table     string `json:"table"`
	Url       string `json:"url"`
	UserId    string `json:"user_id"`
	Records   int    `json:"records"`
	CreatedAt int    `json:"time"`
}
