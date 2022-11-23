package eclosion

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Eclosion struct {
	AppName, Version, RootPath string
	Debug                      bool
	ErrorLog                   *log.Logger
	InfoLog                    *log.Logger
	Config                     internalConfig
	Routes                     *chi.Mux
	DB                         Database
}

type internalConfig struct {
	Port     string
	database databaseConfig
}

func (ecl *Eclosion) Hatch(rootPath string) error {

	pathConfig := appPaths{
		rootPath: rootPath,
		dirNames: []string{"handlers", "migrations", "views", "models", "public", "tmp", "logs", "middleware"},
	}

	err := ecl.Init(pathConfig)
	if err != nil {
		return err
	}

	err = ecl.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	//read .env file
	err = godotenv.Load(rootPath + "/.env")

	if err != nil {
		return err
	}

	//create loggers
	infoLog, errorLog := ecl.startLogging()

	//connect db
	if os.Getenv("DATABASE_TYPE") != "" {
		db, err := ecl.InitDB(os.Getenv("DATABASE_TYPE"), ecl.BuildDSN())

		if err != nil {
			errorLog.Println(err)
			os.Exit(1)
		}

		ecl.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

	ecl.InfoLog = infoLog
	ecl.ErrorLog = errorLog
	ecl.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	ecl.Version = version
	ecl.RootPath = rootPath

	ecl.Config = internalConfig{
		Port: os.Getenv("PORT"),
		database: databaseConfig{
			database: os.Getenv("DATABASE_TYPE"),
			dsn:      ecl.BuildDSN(),
		},
	}

	ecl.Routes = ecl.routes().(*chi.Mux)

	return nil
}

func (ecl *Eclosion) Init(p appPaths) error {

	root := p.rootPath

	for _, path := range p.dirNames {

		//create folder if doesn't exists
		err := ecl.CreateDirIfNotExists(root + "/" + path)

		if err != nil {
			return err
		}
	}
	return nil
}

func (ecl *Eclosion) checkDotEnv(path string) error {

	err := ecl.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))

	if err != nil {
		return err
	}
	return nil
}

func (ecl *Eclosion) startLogging() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger
	log.SetPrefix("Eclosion!")

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (ecl *Eclosion) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     ecl.ErrorLog,
		Handler:      ecl.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	defer ecl.DB.Pool.Close()

	ecl.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))

	err := srv.ListenAndServe()

	ecl.ErrorLog.Fatal(err)
}

func (ecl *Eclosion) BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_NAME"), os.Getenv("DATABASE_SSL_MODE"))

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}
	default:

	}
	return dsn
}
