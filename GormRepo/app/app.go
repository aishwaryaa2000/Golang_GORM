package app

import (
	"context"
	"fmt"
	"gorm/model"
	"gorm/router"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm/config"
)

type App struct {
	Config *config.Config
	DB     *gorm.DB
	Router *mux.Router
	server *http.Server
}
var db *gorm.DB

func New(router *mux.Router) *App{
	var defaults = make(map[string]interface{})
	appConfig := config.NewConfig(defaults)
	app := App{Config: appConfig} 
	app.Router = router
	app.Initialize()
	app.MigrateDB()
	db = app.DB
	return &app
}

func (app *App) GetConnectionString() string {
	dbHost := app.Config.GetString("DB_HOST")
	dbName := app.Config.GetString("DB_NAME")
	dbPort := app.Config.GetString("DB_PORT")
	dbUser := app.Config.GetString("DB_USER")
	dbPassword := app.Config.GetString("DB_PWD")

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true&charset=utf8&parseTime=True&loc=Local&tls=preferred", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func (app *App) Initialize() {

	app.Router = router.CreateRoute()

	apiPort := "80"

	app.server = &http.Server{
		Addr:         "0.0.0.0:" + apiPort,
		WriteTimeout: time.Second * time.Duration(app.Config.GetInt("HTTP_WRITE_TIMEOUT")),
		ReadTimeout:  time.Second * time.Duration(app.Config.GetInt("HTTP_READ_TIMEOUT")),
		IdleTimeout:  time.Second * time.Duration(app.Config.GetInt("HTTP_IDLE_TIMEOUT")),
		Handler:      app.Router,
	}
}

//Start http server and start listening to the requests
func (app *App) Start() {

	if err := app.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal("Unable to start server, exiting the application!")
		}
	}

}


func (app *App) Stop() {
	wait, _ := time.ParseDuration("2m")
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	app.server.Shutdown(ctx)

	sqlDB, err := app.DB.DB()
	if err != nil {
		sqlDB.Close()
	}
	
}

func (app *App) MigrateDB() {

	migrateDB, err := gorm.Open(mysql.Open(app.GetConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to open DB connection for migration, exiting the application!")
	}
	createTables(migrateDB)
}

func createTables(db *gorm.DB) {
	//db := connect.SetupDB().Debug()

	if !db.Migrator().HasTable(&model.User{}){
		db.AutoMigrate(&model.User{})
		fmt.Println("User table created")
	}
	if !db.Migrator().HasTable(&model.Course{}){
		db.AutoMigrate(&model.Course{})
		fmt.Println("Course table created")

	}
	if !db.Migrator().HasTable(&model.Hobby{}){
		db.AutoMigrate(&model.Hobby{})
		fmt.Println("Hobby table created")

	}
}

// NewUnitOfWork creates new UnitOfWork
// func (app *App) NewUnitOfWork(readOnly bool, logger zerolog.Logger) *repository.UnitOfWork {
// 	return repository.NewUnitOfWork(app.DB, readOnly, logger, log.Config{SlowThreshold: time.Duration(app.Config.GetInt(config.EvSuffixForGormSlowThreshold)) * time.Millisecond})
// }
