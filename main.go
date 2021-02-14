package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var courses []Course

func main() {
	router := mux.NewRouter()

	courses = append(courses, Course{ID: "1", Title: "My first course", Body: "This is the content of my first course"})

	router.HandleFunc("/courses", getCourses).Methods("GET")
	router.HandleFunc("/courses", createCourse).Methods("POST")
	router.HandleFunc("/course/{id}", getCourse).Methods("GET")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	found := false

	for _, item := range courses {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			found = true
			break
		}
	}

	if !found {
		json.NewEncoder(w).Encode("Course is not found")
	}
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	course.ID = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(&course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.ID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(&course)

			return
		}
	}

	json.NewEncoder(w).Encode(courses)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range courses {
		if item.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(courses)
}
