package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"studentPackage/internal/storage"
	typesFile "studentPackage/internal/type"
	"studentPackage/internal/utils/response"

	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a new student")
		var student typesFile.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}
		if err := validator.New().Struct(student); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors), http.StatusBadRequest))
			return
		}
		lastid, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrorResponse(err.Error(), http.StatusBadRequest))
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]any{"status": "Student Created", "statusCode": http.StatusCreated, "lastId": lastid})
	}
}

func GetStudent(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {

			return
		}

		stud, err := storage.GetStudent(idInt)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErrorResponse(err.Error(), http.StatusInternalServerError))
			return
		}

		response.WriteJson(w, http.StatusOK, stud)

	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stud, err := storage.GetStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusOK, stud)

	}
}
