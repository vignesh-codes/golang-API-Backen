package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	config "todo/v1/configs"
	"todo/v1/models"
	model "todo/v1/models"

	"github.com/gorilla/mux"
)

var (
	id         int
	task       string
	desc       string
	status     string
	created_at string
	updated_at string
	database   = config.Database()
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// sortTime := r.URL.Query().Get("valB")

	db_exec, err := database.Query(`SELECT * FROM todo_table`)

	// db_exec, err := database.Query(`SELECT * FROM todo_table WHERE status = ? ORDER BY created_at ASC`, filterVal)
	// fmt.Printf("DB TYPE: %T", db_exec)

	if err != nil {
		fmt.Println(err)
		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		return
	}

	var todos []model.TodoModel

	for db_exec.Next() {
		err := db_exec.Scan(&id, &task, &desc, &status, &created_at, &updated_at)

		if err != nil {
			es := &models.ErrorHandler{"Bad value", "400"}
			jsonResponse(w, es)
			return
		}

		todo := model.TodoModel{
			Id:        id,
			Task:      task,
			Desc:      desc,
			Status:    status,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
		}

		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
	handler := &models.ResponseHandler{"Success", 201}
	jsonResponse(w, handler)

}

func CreateOne(w http.ResponseWriter, r *http.Request) {

	task := r.FormValue("task")
	desc := r.FormValue("desc")
	status := "not-started"

	taskV, descV := Validate(task, desc)

	if !taskV {
		fmt.Println("error")

		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		return
	}
	if !descV {
		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		return
	}

	_, err := database.Exec(`INSERT INTO todo_table (task, description, status) VALUE(?,?,?)`, task, desc, status)

	if err != nil {
		fmt.Println(err)
	}
	handler := &models.ResponseHandler{"Success", 200}
	jsonResponse(w, handler)

}

func jsonResponse(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if error_check(res, err) {
		return
	}

	fmt.Fprintf(res, string(payload))
}

func error_check(res http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
func Validate(task string, desc string) (bool, bool) {
	a := true
	b := true
	if len(task) > 40 {
		fmt.Println("Invalid text", task, len(task))
		a = false

	}
	if len(desc) > 256 {
		fmt.Println("Invalid text")
		b = false
	}
	return a, b
}
func DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM todo_table WHERE id = ?`, id)

	if err != nil {
		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		fmt.Println(err)
		return
	}

	handler := &models.ResponseHandler{"Success", 201}
	jsonResponse(w, handler)
}

func CompleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// newStatus := "completeOned"
	_, err := database.Exec(`UPDATE todo_table SET status = "completed" WHERE id = ?`, id)

	if err != nil {

		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		fmt.Println(err)
		return
	}

	handler := &models.ResponseHandler{"Success", 201}
	jsonResponse(w, handler)
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// newStatus := "completeOned"
	task := r.FormValue("task")
	desc := r.FormValue("desc")

	taskV, descV := Validate(task, desc)

	if !taskV {
		fmt.Println("error")

		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)

		return
	}
	if !descV {
		fmt.Println("false")
		es := &models.ErrorHandler{"Bad value", "400"}
		jsonResponse(w, es)
		return
	}

	_, err := database.Exec(`UPDATE todo_table SET task = ?, description = ? WHERE id = ?;`, task, desc, id)

	if err != nil {
		fmt.Println(err)
	}

	jsonResponse(w, "Success")
	handler := &models.ResponseHandler{"Success", 201}
	jsonResponse(w, handler)
}
