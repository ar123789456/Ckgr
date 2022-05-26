package server

import (
	"cgr/link"
	"cgr/news"
	"cgr/tool/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	//repo
	mongoLink "cgr/link/repository/mongo"
	mongoNews "cgr/news/repository/sqlite"

	//usecase
	linkuc "cgr/link/usecase"
	newsuc "cgr/news/usecase"

	//handlers
	linkhttp "cgr/link/delivery/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	httpServer *http.Server

	LinkUC link.UseCase
	NewsUC news.UseCase
}

func NewApp() *App {
	db := initDB()

	linkRepo := mongoLink.NewRepository(db.Collection("link"))
	newsRepo := mongoNews.NewRepository(db.Collection("news"))
	logg := logger.NewLogger()
	return &App{
		LinkUC: linkuc.NewUsecase(linkRepo, logg),
		NewsUC: newsuc.NewUsecase(newsRepo, logg),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	linkhttp.RegisterLink(router, a.LinkUC, logger.NewLogger())

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database("dbCGA")
}
