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

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json: "website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - coursewebsite.in")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "Reactjs", CoursePrice: 123, Author: &Author{FullName: "Vikash Kumar", Website: "Vikashkumar.live"}})
	courses = append(courses, Course{CourseID: "1", CourseName: "K8s", CoursePrice: 213, Author: &Author{FullName: "Vikash Kumar", Website: "Vikashkumar.live"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":5000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id return the response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with Given ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One course")
	w.Header().Set("Content-Type", "application/json")
	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")

	}
	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("please send some data")
		return
	}
	// for index, course := range courses{
	// 	if course.CourseID == {

	// 	}
	// }

	// generate unique id, string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	json.NewEncoder(w).Encode(course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Course Added Succesfully")
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from response
	params := mux.Vars(r)
	// loop, id, remove, add with my id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			json.NewEncoder(w).Encode("Couse Update Succesfully")
			return
		}
	}
	// TODO; send a response when id is not found
}
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id , remove, (index, index+1)
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Course Not found")
}
