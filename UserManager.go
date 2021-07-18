package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"net/http"
	"sync"
)

type UserManager struct {
	db *sql.DB
	m  sync.Mutex
}

func NewManager() (*UserManager, error) {
	err := OpenCfg()
	if err != nil {
		return nil, err
	}

	usr := UserManager{m: sync.Mutex{}}

	usr.db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDB))
	if err != nil {
		return nil, err
	}

	return &usr, err
}

func StartUserManager() {
	srv, err := NewManager()
	if err != nil {
		panic(err)
	}
	defer srv.db.Close()

	m := chi.NewMux()
	m.Get("/", srv.HandleGet)
	m.Post("/", srv.HandlePost)

	addr := cfg.ServerHost + ":" + cfg.ServerPort
	fmt.Print("Starting server at ", addr)
	http.ListenAndServe(addr, m)
}

func (u *UserManager) HandleGet(r http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("u")
	if name == "" {
		r.Write([]byte("Empty request"))
		return
	}

	q, err := u.db.Query(`SELECT full_name FROM "user" WHERE full_name=$1`, name)
	if err != nil {
		r.Write([]byte("Query error"))
		return
	}
	defer q.Close()

	if !q.Next() {
		r.Write([]byte("No user found"))
		return
	}

	r.Write([]byte("Successful GET request: " + name))
}

func (u *UserManager) HandlePost(r http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("p")
	if name == "" {
		r.Write([]byte("Empty request"))
		return
	}

	u.m.Lock()
	q, err := u.db.Query(`SELECT full_name FROM "user" WHERE full_name=$1`, name)
	if err != nil {
		r.Write([]byte("Query error"))
		return
	}

	if q.Next() {
		r.Write([]byte("User " + name + " already exists"))
		return
	}

	q, err = u.db.Query(`INSERT INTO "user" VALUES ($1)`, name)
	if err != nil {
		r.Write([]byte("Query error"))
		return
	}
	u.m.Unlock()

	r.Write([]byte("Successful POST request: " + name))
}
