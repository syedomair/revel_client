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
    return c.Render(user_id, data)
}

func (c Application) MyBooks() revel.Result {

    resultMap := make(map[string]interface{})
    resultMap = services.BookService{}.GetMyBooks(c.Session)
    resultMap1 := make(map[string]interface{})
    resultMap1 = resultMap["data"].(map[string]interface{})
    data := resultMap1["list"]

    user_id := c.Session["user_id"]
    return c.Render(user_id, data)
}

func (c Application) AddBookForm() revel.Result {
    user_id := c.Session["user_id"]
    return c.Render(user_id)
}

func (c Application) EditBookForm(book_id string) revel.Result {
    resultMap := make(map[string]interface{})
    resultMap = services.BookService{}.GetABook(book_id, c.Session)
    data := make(map[string]interface{})
    data = resultMap["data"].(map[string]interface{})

    user_id := c.Session["user_id"]
    return c.Render(user_id, data)
}

func (c Application) EditBook(book models.Book) revel.Result {

    bookJson, err := json.Marshal(book)
    if err != nil {
        fmt.Println("error:", err)
    }
    bookId := strconv.Itoa(book.Id)
    _ = services.BookService{}.EditBook(bookId, bookJson, c.Session)

    return c.Redirect(routes.Application.MyBooks())
}

func (c Application) AddBook(book models.Book) revel.Result {

    user_id := c.Session["user_id"]
    book.UserId, _ = strconv.ParseInt(user_id, 10, 64)

    bookJson, err := json.Marshal(book)
    if err != nil {
        fmt.Println("error:", err)
    }
    _ = services.BookService{}.AddBook(bookJson, c.Session)

    return c.Redirect(routes.Application.MyBooks())
}


func (c Application) Signup(user models.User) revel.Result {

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
        }
    }
    return c.Redirect(routes.Application.Index())

}

func (c Application) Signin(email string, password string ) revel.Result {

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
    }
    if(resultMap["result"]=="error"){
        for k := range c.Session {
            delete(c.Session, k)
        }
        return c.RenderJson(resultMap["data"])
    }
    return c.RenderJson(response)
}

func (c Application) GetJWTToken(email string, password string ) revel.Result {

    password = b64.StdEncoding.EncodeToString([]byte(password )) 

    jwtStr := services.CommonService{}.CreateJWTToken(email, password)

    return c.RenderJson(jwtStr)
}

func (c Application) Signout() revel.Result {
    for k := range c.Session {
        delete(c.Session, k)
    }
    return c.Redirect(routes.Application.Index())
}
