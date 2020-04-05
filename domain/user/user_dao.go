package user

import (
	"fmt"
	"tripplanner/datasources/mysql/users_db"
	"tripplanner/utils/error"

	"github.com/go-sql-driver/mysql"
)

var (
	userDB = make(map[int64]*User)
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, phone, created_at, password, status) VALUES(?,?,?,?,?,?,?)"
	querySelectUser  = "SELECT id, first_name, last_name, email, phone, created_at, status FROM users_db.users WHERE id=?"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=?, phone=? WHERE id=?"
	queryDeleteUser  = "DELETE FROM users WHERE id=?"
	queryFindUser    = "Select id, first_name, last_name, email, created_at, phone, status FROM users_db.users WHERE status=?"
)

func (user *User) Save() *error.AppErr {
	insrtStmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return error.NewInternalServerError(err.Error())
	}
	defer insrtStmt.Close()

	insrtRslt, err := insrtStmt.Exec(user.FirstName, user.LastName, user.Email, user.Phone, user.CreatedAt, user.Password, user.Status)

	if err != nil {
		sqlErr, err := err.(*mysql.MySQLError)
		if err {
			return error.NewInternalServerError("Error saving new user")
		}
		if sqlErr != nil {
			switch sqlErr.Number {
			case 1062:
				return error.NewBadRequestError(fmt.Sprintf("Email id %s already exists", user.Email))
			default:
				return error.NewInternalServerError(fmt.Sprintf("Error saving new user: %s", sqlErr.Error()))
			}
		}
	}

	insrtId, err := insrtRslt.LastInsertId()
	if err != nil {
		return error.NewInternalServerError(fmt.Sprintf("Error geting last insert id %s", err.Error()))
	}

	user.Id = insrtId
	return nil
}

func (user *User) Get() *error.AppErr {
	slctStmnt, err := users_db.Client.Prepare(querySelectUser)

	if err != nil {
		return error.NewInternalServerError(err.Error())
	}

	defer slctStmnt.Close()

	rslt := slctStmnt.QueryRow(user.Id)
	if err := rslt.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.CreatedAt, &user.Status); err != nil {
		return error.NewNotFoundError(err.Error())
	}
	return nil
}

func (user *User) UpdateUser() *error.AppErr {
	updateStmnt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		return error.NewInternalServerError(err.Error())
	}
	defer updateStmnt.Close()

	if _, err = updateStmnt.Exec(user.FirstName, user.LastName, user.Email, user.Phone, user.Id); err != nil {
		sqlErr, isSQLErr := err.(*mysql.MySQLError)
		if !isSQLErr {
			return error.NewInternalServerError(err.Error())
		}
		if sqlErr != nil {
			switch sqlErr.Number {
			case 1062:
				return error.NewInternalServerError(sqlErr.Message)
			default:
				return error.NewInternalServerError(sqlErr.Message)
			}
		}
	}
	return nil
}

func (user *User) Delete() *error.AppErr {

	delStmnt, err := users_db.Client.Prepare(queryDeleteUser)

	if err != nil {
		return error.NewInternalServerError(err.Error())
	}
	defer delStmnt.Close()

	if _, err = delStmnt.Exec(user.Id); err != nil {
		sqlErr, isSQLErr := err.(*mysql.MySQLError)
		if !isSQLErr {
			return error.NewInternalServerError(err.Error())
		}
		if sqlErr != nil {
			switch sqlErr.Number {
			case 1062:
				return error.NewInternalServerError(sqlErr.Message)
			default:
				return error.NewInternalServerError(sqlErr.Message)
			}
		}
	}
	return nil
}

func (user *User) FindByStatus() (Users, *error.AppErr) {
	findUsrStmnt, err := users_db.Client.Prepare(queryFindUser)

	if err != nil {
		return nil, error.NewInternalServerError(err.Error())
	}
	defer findUsrStmnt.Close()

	rows, err := findUsrStmnt.Query(user.Status)

	if err != nil {
		sqlErr, isSQLErr := err.(*mysql.MySQLError)
		if !isSQLErr {
			return nil, error.NewInternalServerError(err.Error())
		}
		if sqlErr != nil {
			switch sqlErr.Number {
			case 1062:
				return nil, error.NewInternalServerError(sqlErr.Message)
			default:
				return nil, error.NewInternalServerError(sqlErr.Message)
			}
		}
	}
	defer rows.Close()
	users := make([]User, 0)

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.Phone, &user.Status); err != nil {
			sqlErr, isSQLErr := err.(*mysql.MySQLError)
			if !isSQLErr {
				return nil, error.NewInternalServerError(err.Error())
			}
			if sqlErr != nil {
				switch sqlErr.Number {
				case 1062:
					return nil, error.NewInternalServerError(sqlErr.Message)
				default:
					return nil, error.NewInternalServerError(sqlErr.Message)
				}
			}
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, error.NewNotFoundError(fmt.Sprintf("No user matchin status %s", user.Status))
	}
	return users, nil
}
