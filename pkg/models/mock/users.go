package mock

import (
	"time"

	"github.com/abassGarane/snippet/pkg/models"
)

var MockUser = &models.User{ 
  ID: 1, 
  Name: "Abass",
  Email: "abass@abass.com",
  Created: time.Now(),        
}

type UserModel struct{}

func (m *UserModel)Insert(name, email, password string)error  {
  switch email {
  case "dupe@email.com":
    return models.ErrorDuplicateEmail
  default:
    return nil
  }
}

func (m *UserModel)Authenticate( email, password string)(int,error)  {
  switch email {
  case "abass@abass.com":
    return 1,nil
  default:
    return 0, models.ErrorInvalidCredentials
  }
}

func (m *UserModel)Get(id int)(*models.User, error)  {
  switch id {
  case 1:
    return MockUser, nil
  default:
    return nil, models.ErrorNoRecord
  }
}

