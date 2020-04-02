package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lin-sel/bookmark-app/controllers"
	"github.com/lin-sel/bookmark-app/repository"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

func main() {
	con := conn()
	muxs := mux.NewRouter()
	route := muxs.PathPrefix("/api/v1").Subrouter()
	repo := repository.NewRepository()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTION"})
	origin := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(route),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":" + port("SERVERPORT"),
	}
	prepareController(con, route, repo)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	defer func() {
		con.Close()
	}()
	fmt.Println("Server ShutDown....")
}

func conn() *gorm.DB {
	fmt.Println(getConncetionString())
	con, err := gorm.Open("mysql", getConncetionString())
	if err != nil {
		fmt.Println(err)
	}
	return con
}

func getConncetionString() string {
	username := os.Getenv("UNAME")
	if username == "" {
		username = "swabhav"
	}
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "swabhav"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	database := os.Getenv("DATABASE")
	if database == "" {
		database = "Swabhav"
	}
	port := port("MYSQLPORT")
	return username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=true"
}

func port(key string) string {
	port := os.Getenv(key)
	if port == "" {
		if key == "SERVERPORT" {
			port = "8080"
		}
		if key == "MYSQLPORT" {
			port = "3306"
		}
	}
	return port
}

func prepareController(con *gorm.DB, route *mux.Router, repo *repository.Repositorysrv) {
	auth := controllers.NewAuthController([]byte(getSecurityKey()))
	userservice := services.NewUserService(con, repo)
	categoryservice := services.NewBookmarkCategoryService(repo, con)
	bookmarkservice := services.NewBookmarkService(repo, con)
	adminservice := services.NewAdminService(con, repo)
	usercontroller := controllers.NewUserController(userservice, auth)
	bookmarkcontroller := controllers.NewController(bookmarkservice, categoryservice, auth)
	admincontroller := controllers.NewAdminController(bookmarkservice, categoryservice, userservice, adminservice, auth)
	fmt.Println(web.GetUUID())
	usercontroller.UserRouteRegister(route)
	bookmarkcontroller.BookmarkRouteRegister(route)
	admincontroller.AdminRouteRegister(route)
}

func getSecurityKey() string {
	if key := os.Getenv("SECURITYKEY"); key != "" {
		return key
	}
	return "mykey"
}
