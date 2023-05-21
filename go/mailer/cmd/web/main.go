package main

import (
	"database/sql"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/wooknight/going_in_circles/go/mailer/data"
)

const webPort = "8001"

func main() {
	//connect to a database to store data
	db := initDB()
	db.Ping()
	//connect to redis to cache info and sessions
	session := initSession()

	//create logger
	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)
	errLog := log.New(os.Stdout, "INFO\t", log.LstdFlags|log.Lshortfile|log.Lmsgprefix)

	//create channels so that we can
	//waitgroup so that we can wait on the app that sends emails
	wg := sync.WaitGroup{}
	app := Config{
		Session:  session,
		DB:       db,
		Wait:     &wg,
		InfoLog:  infoLog,
		ErrorLog: errLog,
		Models:   data.New(db),
	}
	//set up mai l
	app.Mailer = app.createMail()
	go app.listenForMail()

	//listen for termination signals
	go app.listenForShutdown()

	//listen for web conn
	app.serve()
}

func initDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	for {
		conn, err := sql.Open("pgx", dsn)
		if err == nil {
			err = conn.Ping()
			if err == nil {
				return conn
			}
		}
		counts++
		if counts >= 10 {
			log.Println("unable to connect to the database")
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func initSession() *scs.SessionManager {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	gob.Register(data.User{})
	session := scs.New()
	session.Store = redisstore.New(redisPool)
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	return session
}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT|syscall.SIGTERM)
	<-quit
	app.InfoLog.Println("waiting on waitgroup")
	app.Wait.Wait()
	app.Mailer.DoneChan <- true

	app.InfoLog.Println("running cleanup tasks")
	app.InfoLog.Println("closing channels and shutting down app")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	os.Exit(1)
}

func (app *Config) createMail() Mail {
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "info",
		FromAddress: "info@mycompany.com",
		Wait:        app.Wait,
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    mailerDoneChan,
	}
	return m
}
