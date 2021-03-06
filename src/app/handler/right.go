package handler

import (
	"goproject/src/app/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllRights(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rights := []model.Right{}
	db.Find(&rights)
	respondJSON(w, http.StatusOK, rights)
}

func CreateRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	right := model.Right{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, right)
}

func GetRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}
	respondJSON(w, http.StatusOK, right)
}

func UpdateRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&right); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, right)
}

func DeleteRight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["RightID"])
	if err != nil {
		return
	}
	right := getRightOr404(db, id, w, r)
	if right == nil {
		return
	}
	if err := db.Delete(&right).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func getRightOr404(db *gorm.DB, rightID int, w http.ResponseWriter, r *http.Request) *model.Right {
	right := model.Right{}
	if err := db.First(&right, model.Right{Model: gorm.Model{ID: uint(rightID)}}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &right
}
