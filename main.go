package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

var top_words_path = "../plsa/model/top_words"

// var templates = template.Must(template.ParseGlob("./templates/*"))

type TopicModels struct {
	Topics []string
}

/*
 * Usage: load topic models from file
 * @param: filename string = top_words_path + "txt"
 */
func loadTopicModels(filepath string) (*TopicModels, error) {
	filename := filepath + ".txt"
	fmt.Println(filename)

	// Read files
	fin, err := os.Open(filename)
	defer fin.Close()
	if err != nil {
		panic(err)
		return nil, err
	}
	reader := bufio.NewReader(fin)
	var topics [20]string
	index := -1
	num_topics := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if strings.Contains(line, "----") {
			continue
		} else if strings.Contains(line, "Topic #") {
			index++
			num_topics = 0
		} else if num_topics > 5 {
			continue
		} else {
			info := strings.Split(line, " ")
			topics[index] += info[0] + " "
			num_topics++
		}
	}

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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/index", viewHandler)
	http.ListenAndServe(":8090", nil)
}
