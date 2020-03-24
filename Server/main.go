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
	route := mux.NewRouter()
	repo := repository.NewRepository()
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(route),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
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
	con, err := gorm.Open("mysql", "swabhav:swabhav@tcp(127.0.0.1)/Swabhav?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	return con
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
