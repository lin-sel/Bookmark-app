package main

import (
	"context"
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
	"github.com/lin-sel/bookmark-app/auth"
	"github.com/lin-sel/bookmark-app/bookmark"
	"github.com/lin-sel/bookmark-app/repository"
)

func main() {
	con := conn()
	route := mux.NewRouter()
	repo := repository.NewRepository()
	srvs := auth.NewAuthsrv(con, repo)
	control := auth.NewAuthController(srvs)
	control.RouterRgstr(route)
	srvs1 := bookmark.NewBMCService(repo, con)
	// control1 := bookmark.NewBMCController(srvs1)
	// control1.RouterRgstr(route)
	srvs2 := bookmark.NewBMService(repo, con)
	control2 := bookmark.NewController(srvs2, srvs1)
	control2.RouterRegstr(route)
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})
	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(route),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}
	var wait time.Duration

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	defer func() {
		con.Close()
	}()
	srv.Shutdown(ctx)
	fmt.Println("Server ShutDown....")
}

func conn() *gorm.DB {
	con, err := gorm.Open("mysql", "swabhav:swabhav@tcp(127.0.0.1)/Swabhav?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	return con
}
