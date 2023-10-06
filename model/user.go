package model

import "log"

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) error {
	sqlInsertStatement := `insert into users(name, email, password) values($1, $2, $3);`

	_, err := db.Exec(sqlInsertStatement, user.Name, user.Email, user.Password)
	return err
}

func GetUser(id string) (User, error) {
	var user User

	sqlSelectStatement := `select * from users where id=$1;`
	rows, err := db.Query(sqlSelectStatement, id)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func CheckEmail(email string, user *User) bool {
	statement := `select id, name email from users where email=$1`
	rows, err := db.Query(statement, email)
	if err != nil {
		log.Print(err)
		return false
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return false
		}
	}

	return true
}
