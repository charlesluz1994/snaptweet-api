package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

/*
Create a repository of users.
Recebe um banco que foi aberto pelo controller, a func pega o banco e joga neste struct.
E dentro dessa instancia de struct, vao ter os métodos que vao se conectar diretamente com o banco de dados.
*/
func NewUsersRepository(db *sql.DB) *users {
	return &users{db: db}
}

func (repository users) Create(user models.User) (uint64, error) {

	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastInsertedId, err := result.LastInsertId()

	return uint64(lastInsertedId), nil
}
