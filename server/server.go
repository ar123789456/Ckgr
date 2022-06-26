package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cgr/tool/logger"

	//link
	handlerLink "cgr/link/delivery/http"
	repoLink "cgr/link/repository/mongo"
	usecaseLink "cgr/link/usecase"

	//news
	handlerNews "cgr/news/delivery/http"
	repoNews "cgr/news/repository/sqlite"
	usecaseNews "cgr/news/usecase"

	//project
	handlerProject "cgr/project/delivery/http"
	repoproject "cgr/project/repository"
	usecaseproject "cgr/project/usecase"

	//user
	handleruser "cgr/user/delivery/http"
	repouser "cgr/user/repository/mongo"
	usecaseuser "cgr/user/usecase"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server
	router     *gin.Engine
}

func NewApp() *App {
	db := initDB()
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	//file serve
	router.StaticFS("/more_static", http.Dir("./"))

	logg := logger.NewLogger()

	//User
	userRepo := repouser.NewRepository(db)
	usecaseUser := usecaseuser.NewUsecase(userRepo)
	handlerUser := handleruser.NewHandler(usecaseUser)

	userEndlints := router.Group("/users")
	{
		userEndlints.POST("", handlerUser.AuthMiddleWare(), handlerUser.Create)
		userEndlints.PUT("/:id", handlerUser.AuthMiddleWare(), handlerUser.Update)
		userEndlints.DELETE("/:id", handlerUser.AuthMiddleWare(), handlerUser.Delete)
		userEndlints.POST("/:id", handlerUser.LogIn)
	}

	//link
	LinkRepo := repoLink.NewRepository(db)
	LinkUsecase := usecaseLink.NewUsecase(LinkRepo, logg)
	LinkHandler := handlerLink.NewHandler(LinkUsecase, logg)

	linkEndpoints := router.Group("/link")
	{
		linkEndpoints.GET("", LinkHandler.GetAll)
		linkEndpoints.POST("/new", handlerUser.AuthMiddleWare(), LinkHandler.Create)
		linkEndpoints.DELETE("/delete", handlerUser.AuthMiddleWare(), LinkHandler.Delete)
	}

	//News
	NewsRepo := repoNews.NewRepository(db)
	Usecase := usecaseNews.NewUsecase(NewsRepo, logg)
	newsHandler := handlerNews.NewHandler(Usecase, logg)

	newsEndpoints := router.Group("/news")
	{
		newsEndpoints.GET("", newsHandler.GetAllForClient)
		newsEndpoints.GET("/:id", newsHandler.Get)
		newsEndpoints.GET("/admin", handlerUser.AuthMiddleWare(), newsHandler.GetAllForAdmin)
		newsEndpoints.PUT("/:id/put", handlerUser.AuthMiddleWare(), newsHandler.Update)
		newsEndpoints.DELETE("/:id/delete", handlerUser.AuthMiddleWare(), newsHandler.Delete)
		newsEndpoints.POST("/new", handlerUser.AuthMiddleWare(), newsHandler.Create)
	}

	//Project
	ProjectRepo := repoproject.NewRepository(db)
	ProjectUsecase := usecaseproject.NewUsecase(ProjectRepo, logg)
	ProjectHandler := handlerProject.NewHandler(ProjectUsecase)

	projectEndpoints := router.Group("/projects")
	{
		projectEndpoints.GET("", ProjectHandler.GetAllForClient)
		projectEndpoints.GET("/:id", ProjectHandler.Get)
		projectEndpoints.GET("/admin", handlerUser.AuthMiddleWare(), ProjectHandler.GetAllForAdmin)
		projectEndpoints.PUT("/:id", handlerUser.AuthMiddleWare(), ProjectHandler.Update)
		projectEndpoints.DELETE("/:id", handlerUser.AuthMiddleWare(), ProjectHandler.Delete)
		projectEndpoints.POST("/new", handlerUser.AuthMiddleWare(), ProjectHandler.Create)
	}

	return &App{
		router: router,
	}
}

func (a *App) Run(port string) error {

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        a.router,
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

func initDB() *bolt.DB {
	db, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte("link"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("news"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("project"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("token"))
		if err != nil {
			return err
		}
		_, err = t.CreateBucketIfNotExists([]byte("user_name"))
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		log.Panicln(err)
	}
	return db
}
