package controllers

import (
	"github.com/revel/revel"
	"github.com/syedomair/revel_client/app/models"
	"github.com/syedomair/revel_client/app/routes"
	"github.com/syedomair/revel_client/app/services"
        "fmt"
        "encoding/json"
        "strconv"
        b64 "encoding/base64"
)

type UserController struct {
     *revel.Controller
}

func (c UserController) Signup(user models.User) revel.Result {

    user.Password = b64.StdEncoding.EncodeToString([]byte(user.Password ))

    userJson, err := json.Marshal(user)
    if err != nil {
        fmt.Println("error:", err)
    }

    resultMap := make(map[string]interface{})
    resultMap = services.UserService{}.AddUser(userJson, c.Session)

    if(resultMap["result"]=="success"){
        c.Session["email"] = user.Email
        c.Session["password"] = user.Password
        resultMap1 := make(map[string]interface{})
        resultMap1 = services.UserService{}.Authenticate(userJson, c.Session)
        if(resultMap1["result"]=="success"){
            resultMap2 := make(map[string]interface{})
            resultMap2 = resultMap1["data"].(map[string]interface{})
            c.Session["user_id"] = strconv.Itoa(int(resultMap2["id"].(float64)))
            c.Session["user_name"] = resultMap2["first_name"].(string)+ " "+  resultMap2["last_name"].(string)
        }
    }
    return c.Redirect(routes.Application.Index())
}

func (c UserController) Signin(email string, password string ) revel.Result {

    password = b64.StdEncoding.EncodeToString([]byte(password )) 

    strJson := "{\"email\":\""+email+"\",\"password\":\""+password+"\"}"

    c.Session["email"] = email
    c.Session["password"] = password
      
    resultMap := make(map[string]interface{})
    resultMap = services.UserService{}.Authenticate([]byte(strJson), c.Session)

    response := "success"
    if(resultMap["result"]=="success"){
        resultMap1 := make(map[string]interface{})
        resultMap1 = resultMap["data"].(map[string]interface{})
        c.Session["user_id"] = strconv.Itoa(int(resultMap1["id"].(float64)))
        c.Session["user_name"] = resultMap1["first_name"].(string)+ " "+  resultMap1["last_name"].(string)
    }
    if(resultMap["result"]=="error"){
        for k := range c.Session {
            delete(c.Session, k)
        }
        return c.RenderJson(resultMap["data"])
    }
    return c.RenderJson(response)
}

func (c UserController) Signout() revel.Result {
    for k := range c.Session {
        delete(c.Session, k)
    }
    return c.Redirect(routes.Application.Index())
}
