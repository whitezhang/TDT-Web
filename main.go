package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var top_words_path = "../plsa/model/top_words"

// var templates = template.Must(template.ParseGlob("./templates/*"))

type TopicModels struct {
	Topics []string
}

func loadTopicModels(filepath string) (*TopicModels, error) {
	filename := filepath + ".txt"
	fmt.Println(filename)
	var topics [2]string
	topics[0] = "Sport Nice"
	topics[1] = "Nice Sport"
	return &TopicModels{Topics: topics[:]}, nil
	// return &TopicModels{Topics: "t111est"}, nil
}

/*
 * Usage: views handler
 */
func viewHandler(w http.ResponseWriter, r *http.Request) {
	p, _ := loadTopicModels(top_words_path)

	templates := template.Must(template.ParseGlob("./templates/*"))
	err := templates.ExecuteTemplate(w, "indexPage", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// page, _ := loadTopicModels("topics")
	// fmt.Println("%v\n", page.Topics)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/index", viewHandler)
	http.ListenAndServe(":8090", nil)
}
