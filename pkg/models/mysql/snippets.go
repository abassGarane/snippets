package mysql

import (
	"database/sql"

	"github.com/abassGarane/snippet/pkg/models"
)

type SnippetModel struct{
  DB *sql.DB
}

func (m *SnippetModel) Insert(title,content,expires string)(int, error)  {
  tx, _ := m.DB.Begin()
  stmt := `INSERT INTO snippets (title,content,created,expires)
  values(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP,INTERVAL ? DAY))`
  res, err := m.DB.Exec(stmt,title,content,expires)
  if err != nil{
    tx.Rollback()
    return 0, nil
  }
  id, err := res.LastInsertId()
  if err != nil{
    return 0, err
  }
  _=tx.Commit()
  return int(id), nil
}

func (m *SnippetModel)Get(id int)(*models.Snippet, error)  {
  stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() AND id=?`
  row := m.DB.QueryRow(stmt,id)
  s := &models.Snippet{}
  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires )
  if err == sql.ErrNoRows{
    return nil, models.ErrorNoRecord
  }else if err != nil{
    return nil, err
  }
  return s, nil
}

func (m *SnippetModel) Latest()([]*models.Snippet,error)  {

  stmt := `SELECT id, title, content, created, expires FROM snippets
  WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  snippets := []*models.Snippet{}
  for rows.Next(){
    snippet := *&models.Snippet{}
    err := rows.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
    if err != nil{
      return nil, err
    }
    snippets = append(snippets, &snippet)
  }

  if err = rows.Err(); err != nil{
    return nil, err
  }

  return snippets,nil 
}
