package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId     string  `json:"courseid"`
	CourseName   string  `json:"coursename"`
	CoursePrice  int     `json:"courceprice"`
	CourseAuthor *Author `json:"courseauthor"`
}

type Author struct {
	AuthorName    string `json:"authorname"`
	AuthorWebsite string `json:"authorwebsite"`
}

// fake DB

var courses []Course

//middleware - file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Golang API is running on port: 4000")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "JavaScript",
		CoursePrice: 10000,
		CourseAuthor: &Author{
			AuthorName:    "Amit V",
			AuthorWebsite: "www.tayloramitverma.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React",
		CoursePrice: 15000,
		CourseAuthor: &Author{
			AuthorName:    "Amit V",
			AuthorWebsite: "www.tayloramitverma.com",
		},
	})

	// routing
	r.HandleFunc("/", serveName).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", creareOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen server
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

func serveName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Building API</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func creareOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what I do to check data
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send data in JSON")
		return
	}

	// generate unique id and string conversion for course id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove and add new with id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course it not found!")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	// loop, id, remove and add new with id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			json.NewEncoder(w).Encode("This corse has been deleted.")
			return
		}
	}

	json.NewEncoder(w).Encode("Course it not found!")
	return
}
