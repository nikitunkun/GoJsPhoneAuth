package apiserver

import (
	"backend/internal/app/db"
	"backend/internal/app/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	db     db.Database
}

func newServer(database db.Database) *server {
	server := &server{
		router: mux.NewRouter(),
		db:     database,
	}

	server.configurateRouter()

	return server
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *server) configurateRouter() {
	server.router.HandleFunc("/user", server.handleUserCreate()).Methods("POST")
	server.router.HandleFunc("/user/{phone}", server.handleUserGet()).Methods("GET")
	server.router.HandleFunc("/user/{phone}", server.handleUserDelete()).Methods("DELETE")
}

func (server *server) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func (server *server) handleUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.enableCors(&w)
		phone := mux.Vars(r)["phone"]
		user, err := server.db.User().Get(phone)
		if err != nil {
			server.error(w, r, http.StatusBadRequest, err)
			return
		}
		server.respond(w, r, http.StatusOK, user)
	}
}

func (server *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		server.enableCors(&w)
		req := &request{}

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			server.error(w, r, http.StatusBadRequest, err)
			return
		}

		user := &model.User{
			Name:  req.Name,
			Phone: req.Phone,
		}
		if err := server.db.User().Create(user); err != nil {
			server.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		server.respond(w, r, http.StatusCreated, user)
	}
}

func (server *server) handleUserDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.enableCors(&w)
		phone := mux.Vars(r)["phone"]
		user, err := server.db.User().Delete(phone)
		if err != nil {
			server.error(w, r, http.StatusBadRequest, err)
			return
		}
		server.respond(w, r, http.StatusOK, user)
	}
}

func (server *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	server.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
