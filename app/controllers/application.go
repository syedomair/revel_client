package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_client/app/services"
	"github.com/syedomair/revel_client/app/routes"
        b64 "encoding/base64"
)

type Application struct {
     *revel.Controller
}

func (c Application) Index() revel.Result {

    return c.Redirect(routes.Application.About())
}

func (c Application) About() revel.Result {

    user_id := c.Session["user_id"]
    user_name := c.Session["user_name"]
    return c.Render(user_id, user_name)
}

func (c Application) GetJWTToken(email string, password string ) revel.Result {

    password = b64.StdEncoding.EncodeToString([]byte(password )) 

    jwtStr := services.CommonService{}.CreateJWTToken(email, password)

    return c.RenderJson(jwtStr)
}
