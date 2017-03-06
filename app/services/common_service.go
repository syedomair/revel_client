package services

import (
	"github.com/revel/revel"
        "net/http"
        "bytes"
        "io/ioutil"
        "github.com/dgrijalva/jwt-go"
        "encoding/json"
)

type CommonService struct {
}

func (c CommonService) CallGetBackend(methodType string, url string, session revel.Session) map[string]interface{} {
    return c.CallBackend(methodType, url, []byte(""), false, session)
}

func (c CommonService) CallMircoBackend(methodType string, url string, jsonStr []byte) []map[string]interface{} {

    completeURL := revel.Config.StringDefault("microservice_server_url", "") +  url

    request, err := http.NewRequest(methodType, completeURL, bytes.NewBuffer(jsonStr))

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    var bodyInterface []map[string]interface{}
    json.Unmarshal(body, &bodyInterface)
    
    return bodyInterface
}

func (c CommonService) CallBackend(methodType string, url string, jsonStr []byte, newuser bool, session revel.Session) map[string]interface{} {

    apiKey := revel.Config.StringDefault("api_key", "")

    var userName string
    var password string

    if(newuser){
        userName = "new_registration"
        password = "new_registration"
    }else{
        userName = session["email"]
        password = session["password"]
    }

    signedJwtToken := c.CreateJWTToken(userName, password)

    completeURL := revel.Config.StringDefault("backend_server_url", "") +  url

    request, err := http.NewRequest(methodType, completeURL, bytes.NewBuffer(jsonStr))
    request.Header.Set("x-key", apiKey)
    request.Header.Set("x-jwt", signedJwtToken)
    request.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

    body, _ := ioutil.ReadAll(response.Body)

    var bodyInterface map[string]interface{}
    json.Unmarshal(body, &bodyInterface)
    
    return bodyInterface
}

func (c CommonService) CreateJWTToken(userName string, password string) string {
    apiSecret := revel.Config.StringDefault("api_secret", "")
    signingKey := []byte(apiSecret)

    type Claims struct {
        Username string `json:"username"`
        Password string `json:"password"`
        jwt.StandardClaims
    }

    claims := Claims{
        userName,
        password,
        jwt.StandardClaims{
            Issuer: "test",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedJwtToken,_ := token.SignedString(signingKey)

    return signedJwtToken
}
