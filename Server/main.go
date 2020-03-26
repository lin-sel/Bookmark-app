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
)

func main() {
	con := conn()
	muxs := mux.NewRouter()
	route := muxs.PathPrefix("/api/v1/user").Subrouter()
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
	authservice := services.NewUserService(con, repo)
	bookmarkcategoryservice := services.NewBookmarkCategoryService(repo, con)
	bookmarkservice := services.NewBookmarkService(repo, con)
	authcontroller := controllers.NewUserController(authservice)
	bookmarkcontroller := controllers.NewController(bookmarkservice, bookmarkcategoryservice)
	authcontroller.RouterRgstr(route)
	bookmarkcontroller.RouterRegstr(route)
}
