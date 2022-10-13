package mysql

import (
	"database/sql"

	"github.com/abassGarane/snippet/pkg/models"
)

type SnippetModel struct{
  DB *sql.DB
}

func (m *SnippetModel) Insert(title,content,expires string)(int, error)  {
  stmt := `INSERT INTO snippets (title,content,created,expires)
  values(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP,INTERVAL ? DAY))`
  res, err := m.DB.Exec(stmt,title,content,expires)
  if err != nil{
    return 0, nil
  }
  id, err := res.LastInsertId()
  if err != nil{
    return 0, err
  }
  return int(id), nil
}

func (m *SnippetModel)Get(id int)(*models.Snippet, error)  {
  return nil, nil
}

func (m *SnippetModel) Latest()([]*models.Snippet,error)  {
  return nil,nil 
}
