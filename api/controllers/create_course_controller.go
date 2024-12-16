package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	o := orm.NewOrm()
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	_, err = o.Insert(&course)
	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]interface{}{"message": "Course created sucessfully",
		"id": course.Id}, "", " ")
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
