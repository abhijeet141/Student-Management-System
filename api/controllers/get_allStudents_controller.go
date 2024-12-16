package controllers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	o := orm.NewOrm()
	num, err := o.QueryTable("student").RelatedSel("course").OrderBy("id").All(&students)

	if err != nil {
		http.Error(w, "Error fetching students", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(map[string]interface{}{
		"message": "Students Fetched Successfully",
		"count":   num,
		"data":    students}, "", " ")
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(students)
	w.Write(response)
}
