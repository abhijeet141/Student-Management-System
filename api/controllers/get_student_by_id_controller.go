package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gorilla/mux"
)

func GetStudentsByID(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	o := orm.NewOrm()
	vars := mux.Vars(r)
	id := vars["id"]
	err := o.QueryTable("student").Filter("id", id).RelatedSel("course").One(&student)
	if err == orm.ErrNoRows {
		http.Error(w, "No student found with given ID", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error fetching student by ID", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]interface{}{
		"message": "Student Fetched Successfully",
		"data":    student}, "", " ")
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
