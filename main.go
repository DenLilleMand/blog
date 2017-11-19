package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/alecthomas/kingpin"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

const (
	LOGGING_DATE_FORMAT = "Mon Jan 2 15:04:05 2006"
)

var validGetPath = regexp.MustCompile("^/(index|404)$")

func getTemplates(path string) (*template.Template, error) {
	fmt.Print("")
	templates := template.New("templates")
	absPath, _ := filepath.Abs(path)
	templateFolder, _ := os.Open(absPath)
	defer templateFolder.Close()
	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	templates.ParseFiles(*templatePaths...)
	return templates
}

type PSQLHook struct {
	tableName string
	DB        *sql.DB
}

func NewPSQLHook(DB *sql.DB, tableName string) *PSQLHook {
	return &PSQLHook{}
}

func (psqlHook *PSQLHook) Fire(entry *log.Entry) error {
	data := entry.Data
	DB := psqlHook.DB
	DB.Query("INSERT INTO public.log ")
}

func (psqlHook *PSQLHook) Levels(entry *log.Entry) []log.Level {
	return log.AllLevels()
}

func main() {
	dbUser := kingpin.Flag("dbuser", "Postgresql username").Short("dbu").String()
	dbName := kingpin.Flag("dbname", "Postgresql dbname").Short("dbn").String()
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

	psqlHook := NewPSQLHook(db, "blog")
	log.AddHook(psqlHook)

	templates, err := getTemplates("templates")
	if err != nil {
		log.WithFields(log.Fields{
			"msg":   err.Error(),
			"time":  time.Now().Format(LOGGING_DATE_FORMAT),
			"level": log.ErrorLevel.String(),
		}).Errorln(err.Error())
	}

	router := httprouter.New()
	httprouter.Handle()
}
