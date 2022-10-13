package models

import (
	"errors"
  "time"
)

var ErrorNoRecord = errors.New("models: no matching record found")

type Snippet struct{
  ID int
  Title string
  Content string
  Created time.Time
  expires time.Time
}
