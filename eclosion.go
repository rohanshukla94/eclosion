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
	Config                     InternalConfig
	Routes                     *chi.Mux
}

type InternalConfig struct {
	Port             string
	TemplateRenderer string
}

func (ecl *Eclosion) Hatch(rootPath string) error {

	pathConfig := appPaths{
		rootPath: rootPath,
		dirNames: []string{"handlers", "migrations", "views", "models", "controllers", "public", "tmp", "logs", "middleware"},
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

	ecl.InfoLog = infoLog
	ecl.ErrorLog = errorLog
	ecl.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	ecl.Version = version
	ecl.RootPath = rootPath
	ecl.Config = InternalConfig{
		Port:             os.Getenv("PORT"),
		TemplateRenderer: os.Getenv("RENDERER"),
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
	log.SetPrefix("Eclosion: ")

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (ecl *Eclosion) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     ecl.ErrorLog,
		Handler:      ecl.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	ecl.InfoLog.Println("Listening on port", os.Getenv("PORT"))

	err := srv.ListenAndServe()

	ecl.ErrorLog.Fatal(err)
}