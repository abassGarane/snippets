package models

import (
	"errors"
  "time"
)

var(
  ErrorNoRecord = errors.New("models: no matching record found")
  ErrorInvalidCredentials = errors.New("models: invalid credentials") 
  ErrorDuplicateEmail = errors.New("models: Duplicate email")
)

type Snippet struct{
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}

type User struct{
  ID int
  Name string
  Email string
  HashedPassword []byte
  Created time.Time
}
