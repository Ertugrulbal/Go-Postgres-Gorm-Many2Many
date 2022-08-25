package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ertugrulbal/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAllRoles(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	roles := []model.Roles{}
	db.Preload("Processes").Find(&roles)
	respondJSON(w, http.StatusOK, roles)
}

func CreateRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	role := []model.Roles{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&role).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, role)
}

func GetRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}

	role := getRoleOr404(db, id, w, r)
	if role == nil {
		return
	}
	respondJSON(w, http.StatusOK, role)
}
func UpdateRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}
	role := getRoleOr404(db, id, w, r)
	if role == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&role).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, role)
}

func GetAllProcesses(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	processes := []model.Process{}
	db.Preload("Roles").Find(&processes)
	respondJSON(w, http.StatusOK, processes)
}
func CreateProcess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	process := []model.Process{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&process); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&process).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, process)
}

func GetProcess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}
	process := getProcessOr404(db, id, w, r)
	if process == nil {
		return
	}
	respondJSON(w, http.StatusOK, process)
}

func UpdateProcess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}
	process := getProcessOr404(db, id, w, r)
	if process == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&process); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&process).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, process)
}

func DeleteRole(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}
	role := getRoleOr404(db, id, w, r)
	if role == nil {
		return
	}
	if err := db.Delete(&role).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
func DeleteProcess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Incorrect ID value.")
	}
	process := getProcessOr404(db, id, w, r)
	if process == nil {
		return
	}
	if err := db.Delete(&process).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getRoleOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Roles {
	role := model.Roles{}
	if err := db.Preload("Processes").First(&role, model.Roles{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &role
}
func getProcessOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Process {
	process := model.Process{}
	if err := db.First(&process, model.Process{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &process
}
