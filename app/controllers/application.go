package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_client/app/services"
        b64 "encoding/base64"
)

type Application struct {
     *revel.Controller
}

func (c Application) Index() revel.Result {

    resultMap := make(map[string]interface{})
    resultMap = services.BookService{}.GetBooks(c.Session)
    resultMap1 := make(map[string]interface{})
    resultMap1 = resultMap["data"].(map[string]interface{})
    data := resultMap1["list"]

    user_id := c.Session["user_id"]
    user_name := c.Session["user_name"]
    return c.Render(user_id, user_name, data)
}

func (c Application) GetJWTToken(email string, password string ) revel.Result {

    password = b64.StdEncoding.EncodeToString([]byte(password )) 

    jwtStr := services.CommonService{}.CreateJWTToken(email, password)

    return c.RenderJson(jwtStr)
}
