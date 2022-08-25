package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ertugrulbal/handler"
	"github.com/ertugrulbal/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	var err error
	dsn := "host=18.185.93.196 user=postgres password=postgres dbname=testErtugrul port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Established")
	}
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/roles", a.GetAllRoles)
	a.Get("/roles/{id}", a.GetRole)
	a.Post("/roles", a.CreateRole)
	a.Put("/roles/{id}", a.UpdateRole)
	a.Delete("/roles/{id}", a.DeleteRole)

	a.Get("/process", a.GetAllProcesses)
	a.Get("/process/{id}", a.GetProcess)
	a.Post("/process", a.CreateProcess)
	a.Put("/process/{id}", a.UpdateProcess)
	a.Delete("/process/{id}", a.DeleteProcess)

}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a *App) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	handler.GetAllRoles(a.DB, w, r)
}

func (a *App) GetRole(w http.ResponseWriter, r *http.Request) {
	handler.GetRole(a.DB, w, r)
}
func (a *App) CreateRole(w http.ResponseWriter, r *http.Request) {
	handler.CreateRole(a.DB, w, r)
}
func (a *App) UpdateRole(w http.ResponseWriter, r *http.Request) {
	handler.UpdateRole(a.DB, w, r)
}

func (a *App) GetAllProcesses(w http.ResponseWriter, r *http.Request) {
	handler.GetAllProcesses(a.DB, w, r)
}
func (a *App) GetProcess(w http.ResponseWriter, r *http.Request) {
	handler.GetProcess(a.DB, w, r)
}
func (a *App) CreateProcess(w http.ResponseWriter, r *http.Request) {
	handler.CreateProcess(a.DB, w, r)
}

func (a *App) UpdateProcess(w http.ResponseWriter, r *http.Request) {
	handler.UpdateProcess(a.DB, w, r)
}

func (a *App) DeleteRole(w http.ResponseWriter, r *http.Request) {
	handler.DeleteRole(a.DB, w, r)
}
func (a *App) DeleteProcess(w http.ResponseWriter, r *http.Request) {
	handler.DeleteProcess(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
