package mysql

import (
	"database/sql"
	"strings"

	"github.com/abassGarane/snippet/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{
  DB *sql.DB
}

func (m *UserModel)Insert(name, email, password string)error  {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  if err != nil{
    return err
  }
  stmt := `INSERT INTO users(name, email, hashed_password, created) 
  VALUES (?, ?, ?, UTC_TIMESTAMP())`
  _, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
  if err != nil{
    if mysqlError, ok:= err.(*mysql.MySQLError); ok{
      if mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users_uc_email"){
        return models.ErrorDuplicateEmail
      }
    }
  }
  return err
}
func (m *UserModel)Authenticate(email, password string)(int,error)  {
  return 0, nil
}
func (m *UserModel)Get(id int)(*models.User,error)  {
  return nil, nil
}

