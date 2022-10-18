package mock

import (
	"time"

	"github.com/abassGarane/snippet/pkg/models"
)

var MockSnippet = &models.Snippet{
  ID: 1, 
  Title: "An elephant",
  Content: "A shattered visage lies...",
  Created: time.Now(),
  Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel)Insert(title, content, expires string)(int, error)  {
  return 2,nil
}

func (m *SnippetModel)Get(id int)(*models.Snippet, error)  {
  switch id  {
  case 1:
    return MockSnippet, nil 
  default:
    return nil, models.ErrorNoRecord
  }
}

func (m *SnippetModel)Latest()([]*models.Snippet, error)  {
  return []*models.Snippet{MockSnippet}, nil
}
