package eclosion

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rohanshukla94/eclosion/render"
	"github.com/rohanshukla94/eclosion/session"
)

const version = "1.0.0"

type Eclosion struct {
	AppName, Version, RootPath string
	Debug                      bool
	ErrorLog                   *log.Logger
	InfoLog                    *log.Logger
	Config                     internalConfig
	Routes                     *chi.Mux
	Render                     *render.Render
	JetViews                   *jet.Set
	Session                    *scs.SessionManager
}

type internalConfig struct {
	Port             string
	TemplateRenderer string
	Cookie           CookieConfig
	SessionType      string
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

	ecl.Config = internalConfig{
		Port:             coalesce(os.Getenv("PORT"), "8080"),
		TemplateRenderer: coalesce(os.Getenv("RENDERER"), "jet"),
		Cookie: CookieConfig{
			Name:     coalesce(os.Getenv("COOKIE_NAME"), "eclosion_session"),
			Lifetime: coalesce(os.Getenv("COOKIE_LIFETIME"), "60"),
			Persist:  coalesce(os.Getenv("COOKIE_PERSIST"), "false"),
			Secure:   coalesce(os.Getenv("COOKIE_SECURE"), ternary(!ecl.Debug, "true", "false"))
			Domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		SessionType: coalesce(os.Getenv("SESSION_TYPE"), "cookie"),
	}

	//create session

	sess := session.Session{
		CookieLifetime: ecl.Config.Cookie.Lifetime,
		CookiePersist:  ecl.Config.Cookie.Persist,
		CookieName:     ecl.Config.Cookie.Name,
		CookieDomain:   ecl.Config.Cookie.Domain,
		SessionType:    ecl.Config.SessionType,
		CookieSecure:   ecl.Config.Cookie.Secure
	}

	ecl.Session = sess.InitSession()

	ecl.Routes = ecl.routes().(*chi.Mux)

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	ecl.JetViews = views

	ecl.createRenderer()
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
		Addr:         fmt.Sprintf(":%s", ecl.Config.Port),
		ErrorLog:     ecl.ErrorLog,
		Handler:      ecl.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	ecl.InfoLog.Printf("Listening on port %s", ecl.Config.Port)

	err := srv.ListenAndServe()

	ecl.ErrorLog.Fatal(err)
}

func (ecl *Eclosion) createRenderer() {
	myRender := render.Render{
		Renderer: ecl.Config.TemplateRenderer,
		RootPath: ecl.RootPath,
		Port:     ecl.Config.Port,
		JetViews: ecl.JetViews,
	}

	ecl.Render = &myRender
}
