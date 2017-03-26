// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tUserController struct {}
var UserController tUserController


func (_ tUserController) Signup(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("UserController.Signup", args).URL
}

func (_ tUserController) Signin(
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("UserController.Signin", args).URL
}

func (_ tUserController) Signout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("UserController.Signout", args).URL
}


type tBookController struct {}
var BookController tBookController


func (_ tBookController) PublishedBooks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.PublishedBooks", args).URL
}

func (_ tBookController) MyBooks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.MyBooks", args).URL
}

func (_ tBookController) AddBookForm(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BookController.AddBookForm", args).URL
}

func (_ tBookController) EditBookForm(
		book_id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book_id", book_id)
	return revel.MainRouter.Reverse("BookController.EditBookForm", args).URL
}

func (_ tBookController) EditBook(
		book interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book", book)
	return revel.MainRouter.Reverse("BookController.EditBook", args).URL
}

func (_ tBookController) AddBook(
		book interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "book", book)
	return revel.MainRouter.Reverse("BookController.AddBook", args).URL
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).URL
}

func (_ tApplication) About(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.About", args).URL
}

func (_ tApplication) GetJWTToken(
		email string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Application.GetJWTToken", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


