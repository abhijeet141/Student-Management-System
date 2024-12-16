package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gorilla/mux"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := o.QueryTable("student").Filter("id", id).Delete()
	if err != nil {
		http.Error(w, "Failed to delete the student", http.StatusInternalServerError)
		return
	}
	if num == 0 {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	response, err := json.MarshalIndent(map[string]interface{}{"message": "Student deleted sucessfully"}, "", " ")
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
