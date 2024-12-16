package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	o := orm.NewOrm()
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	var user models.User
	err = o.QueryTable("user").Filter("id", student.User.Id).One(&user)
	if err != nil {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	var course models.Course
	err = o.QueryTable("course").Filter("id", student.Course.Id).One(&course)
	if err != nil {
		http.Error(w, "Course does not exist", http.StatusBadRequest)
		return
	}
	_, err = o.Insert(&student)
	if err != nil {
		http.Error(w, "Failed to create student", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]interface{}{"message": "Student created sucessfully",
		"id": student.Id}, "", " ")
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
