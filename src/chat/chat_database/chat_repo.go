package chat_database

import "goChat/src/db/db_domain"

const tableName = "chat"

type TodoRepoImpl struct {
	db *db_domain.DB
}
