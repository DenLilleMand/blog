package main

import (
	"database/sql"
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	_ "github.com/lib/pq"

	"github.com/alecthomas/kingpin"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

const (
	LOGGING_DATE_FORMAT = "Mon Jan 2 15:04:05 2006"
)

var validGetPath = regexp.MustCompile("^/(index|404|math)$")

type Controller struct {
	templates *template.Template
}

type PSQLLogHook struct {
	tableName string
	DB        *sql.DB
}

func getTemplates(path string) (*template.Template, error) {
	templates := template.New("templates")
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	templateFolder, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer templateFolder.Close()
	templatePathsRaw, err := templateFolder.Readdir(-1)
	if err != nil {
		return nil, err
	}
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, absPath+"/"+pathInfo.Name())
		}
	}
	templates.ParseFiles(*templatePaths...)
	return templates, nil
}

func NewPSQLLogHook(DB *sql.DB, tableName string) *PSQLLogHook {
	return &PSQLLogHook{DB: DB, tableName: tableName}
}

func (psqlLogHook *PSQLLogHook) Fire(entry *log.Entry) error {
	msg, ok := entry.Data["msg"]
	if !ok {
		return errors.New("Could not get entry.Data['msg'] in PSQLHook.Fire")
	}
	level, ok := entry.Data["level"]
	if !ok {
		return errors.New("Could not get entry.Data['level'] in PSQLHook.Fire")
	}
	var id string
	err := psqlLogHook.DB.QueryRow("INSERT INTO log (msg, level) VALUES ($1, $2) RETURNING id", msg, level).Scan(&id)
	if err != nil {
		return err
	}
	if id == "" {
		return errors.New("psqlLogHook didn't create a valid id")
	}
	return nil
}

func (psqlLogHook *PSQLLogHook) Levels() []log.Level {
	return log.AllLevels
}

func (controller *Controller) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templateName := "index.html"
	template := controller.templates.Lookup(templateName)
	if template == nil {
		errMsg := "template lookup on " + templateName + " returned nil"
		log.WithFields(log.Fields{
			"msg":   errMsg,
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(errMsg)
	}
	template.Execute(w, nil)
}

func (controller *Controller) Math(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templateName := "math.html"
	template := controller.templates.Lookup(templateName)
	if template == nil {
		errMsg := "template lookup on " + templateName + " returned nil"
		log.WithFields(log.Fields{
			"msg":   errMsg,
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(errMsg)
	}
	template.Execute(w, nil)
}

func (controller *Controller) NotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templateName := "404.html"
	template := controller.templates.Lookup(templateName)
	if template == nil {
		errMsg := "template lookup on " + templateName + " returned nil"
		log.WithFields(log.Fields{
			"msg":   errMsg,
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(errMsg)
	}
	template.Execute(w, nil)
}

func (controller *Controller) Static(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fileName := params.ByName("filename")
	if fileName == "" {
		errMsg := "Static request did not contain a filename"
		log.WithFields(log.Fields{
			"msg":   errMsg,
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(errMsg)
	}
	fileServer := http.FileServer(http.Dir("dist"))
	r.URL.Path = fileName
	fileServer.ServeHTTP(w, r)
}

func main() {
	log.SetOutput(os.Stdout)
	dbUser := kingpin.Flag("dbuser", "Postgresql username").String()
	dbName := kingpin.Flag("dbname", "Postgresql dbname").String()
	port := kingpin.Flag("port", "Webserver port").String()
	staticDir := kingpin.Flag("static", "Directory for static files").String()
	kingpin.Parse()

	connectionString := "user=" + *dbUser + " dbname=" + *dbName + " sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.WithFields(log.Fields{
			"msg":   err.Error(),
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(err.Error())
	}

	psqlLogHook := NewPSQLLogHook(db, "log")
	log.AddHook(psqlLogHook)

	templates, err := getTemplates("templates")
	if err != nil {
		log.WithFields(log.Fields{
			"msg":   err.Error(),
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(err.Error())
	}

	controller := &Controller{templates: templates}
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.GET("/math", controller.Math)
	router.GET("/404", controller.NotFound)
	router.GET("/static/:filename", controller.Static)

	infoMsg := "Started server at port:" + (*port)
	log.WithFields(log.Fields{
		"msg":   infoMsg,
		"level": log.InfoLevel.String(),
		"time":  time.Now().Format(LOGGING_DATE_FORMAT),
	}).Infoln(infoMsg)
	http.ListenAndServe(":"+(*port), router)
}
