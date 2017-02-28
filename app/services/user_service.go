package services

import (
	"github.com/revel/revel"
)

type UserService struct {
    CommonService
}

func (c UserService) AddUser(jsonStr []byte, session revel.Session) map[string]interface{} {
    return c.CallBackend("POST", "user", jsonStr, true, session) 
}

func (c UserService) Authenticate(jsonStr []byte, session revel.Session) map[string]interface{} {
    return c.CallBackend("POST", "authenticate", jsonStr, false, session) 
}

