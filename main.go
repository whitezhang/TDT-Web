package main

import (
	"./app"
	"./dao"
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Const Variable
// File path
const top_words_file = "../plsa/model/top_words.txt"
const pzd_file = "../plsa/model/p_z_d.txt"

// number of topics that shown in the home page
const num_topics = 10

// number of keywords for each topic
const num_keywords = 5
const num_documents = 10

// var templates = template.Must(template.ParseGlob("./templates/*"))

/*
 * Strcut that in page returned
 */
type RetDocPage struct {
	EntitiesTrends EntitiesTrends
	NewsDoc        dao.NewsData
}

/*
 * Common struct
 */
type TopicModels struct {
	Topics []string
}

type TopicPostingList struct {
	DocumentsProb [num_topics][]float64
}

/*
 * Usage: Trends struct
 */

type EntityTrendOnTime struct {
	EntityCount []int
}

type EntitiesTrends struct {
	EntityTrendOnTime map[string]EntityTrendOnTime
}

// Usage: get the indices of the sorted slice
type Slice struct {
	sort.Float64Slice
	idx []int
}

func (s Slice) Swap(i, j int) {
	s.Float64Slice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func NewSlice(n []float64) *Slice {
	s := &Slice{Float64Slice: sort.Float64Slice(n), idx: make([]int, len(n))}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

/*
 * Usage: load topic models from file
 */
func loadTopicModels() (*TopicModels, error) {
	fmt.Println("Load topic models from ", top_words_file)

	// Read files
	fin, err := os.Open(top_words_file)
	defer fin.Close()
	if err != nil {
		panic(err)
		return nil, err
	}
	reader := bufio.NewReader(fin)
	var topics [num_topics]string
	index := -1
	count_topics := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)

		if strings.Contains(line, "----") {
			continue
		} else if strings.Contains(line, "Topic #") {
			index++
			count_topics = 0
		} else if count_topics >= num_keywords {
			continue
		} else {
			info := strings.Split(line, " ")
			topics[index] += info[0] + " "
			count_topics++
		}
	}

	return &TopicModels{Topics: topics[:]}, nil
}

// Usage: generate topic posting indices based on index
func generateTopicPostingList(index int) ([]int, error) {
	var topicPostingList TopicPostingList

	fin, err := os.Open(pzd_file)
	defer fin.Close()
	if err != nil {
		panic(err)
		return nil, err
	}
	reader := bufio.NewReader(fin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		info := strings.Split(line, ": ")
		probsString := info[1]
		for index, probString := range strings.Split(probsString, " ") {
			prob, _ := strconv.ParseFloat(probString, 64)
			topicPostingList.DocumentsProb[index] = append(topicPostingList.DocumentsProb[index], prob)
		}
	}

	probsList := NewSlice(topicPostingList.DocumentsProb[index])
	sort.Sort(probsList)
	// s.idx is the indices of the slice
	// fmt.Println(probsList.idx)
	return probsList.idx, nil
}

// Discarded
// func loadDocumentsOnTopics(index int) {
// 	fin, err := os.Open(pzd_file)
// 	defer fin.Close()
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	reader := bufio.NewReader(fin)
// 	// var docProbs []float64
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err != nil || io.EOF == err {
// 			break
// 		}
// 		line = strings.Replace(line, "\n", "", -1)
// 		info := strings.Split(line, ": ")
// 		probsString := info[1]
// 		// sortedProbIndex
// 		fmt.Println("...", probsString, "..")
// 		for _, probString := range strings.Split(probsString, " ") {
// 			prob, _ := strconv.ParseFloat(probString, 64)
// 			fmt.Println(prob)
// 		}
// 	}
// 	return
// }

/*
 * Usage: find the position of the topic
 */
func findTopicPosition(topics string) int {
	topicModels, _ := loadTopicModels()
	position := 0
	for _, b := range topicModels.Topics {
		if b == topics {
			return position
		}
		position++
	}
	return -1
}

/*
 * Usage: home page handler(/)
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	topicsModels, _ := loadTopicModels()

	// Views loading
	templates := template.Must(template.ParseGlob("./templates/*"))
	err := templates.ExecuteTemplate(w, "indexPage", topicsModels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
 * Usage: topic page handler(/topic?keyworkds=xxx)
 */
func topicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	keywords := r.Form["keywords"][0]
	position := findTopicPosition(keywords)

	documentsPostingIndices, err := generateTopicPostingList(position)
	if err != nil {
		panic(err)
		return
	}
	// Map indices to ids
	documetnsPostingIds := make([]string, num_documents)
	for index, value := range documentsPostingIndices[:num_documents] {
		documetnsPostingIds[index], err = app.Index2Id(value)
	}

	// News
	docsInPage := make([]dao.SimpleNewsData, num_documents)
	for index, value := range documetnsPostingIds {
		newsData, err := dao.GetSimpleNewsDataOnID(value)
		if err != nil {
			panic(err)
			return
		}
		docsInPage[index] = *newsData
		fmt.Println(docsInPage[index].Title)
	}

	// fmt.Println(docsInPage)
	// Passed parameter
	pageDict := make(map[string]string)
	pageDict["keywords"] = keywords

	// Views loading
	templates := template.Must(template.ParseGlob("./templates/*"))
	err = templates.ExecuteTemplate(w, "topicPage", docsInPage)
	// err = templates.ExecuteTemplate(w, "topicPage", documetnsPostingIds[:num_documents])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
 * Usage: document page handler(/document?id=xxx)
 */
func documentHandler(w http.ResponseWriter, r *http.Request) {
	var retDocPage RetDocPage

	r.ParseForm()
	sid := r.Form["id"][0]
	newsDoc, err := dao.GetNewsDataOnID(sid)
	retDocPage.NewsDoc = *newsDoc
	// Demo
	// Get IDs that has this entity
	// fmt.Println(sid, "----", app.ExpEntitySet.ExpEntityNode[sid])
	var entitiesTrends EntitiesTrends
	entitiesTrends.EntityTrendOnTime = make(map[string]EntityTrendOnTime)

	for _, name := range app.ExpEntitySet.ExpEntityNode[sid].ExpEntity {
		entitiesTrends.EntityTrendOnTime[name] = EntityTrendOnTime{EntityCount: make([]int, 31)}
		fmt.Println(name)
		idsHasEntity := app.GetIdsFromEntity(name)
		for _, id := range idsHasEntity {
			timeStamp, _ := dao.GetTimeStampOnID(id)
			fmt.Println(id, timeStamp.TimeStamp)
			day, _, _ := app.SplitDate(timeStamp.TimeStamp)
			entitiesTrends.EntityTrendOnTime[name].EntityCount[day]++
		}
	}
	fmt.Println(entitiesTrends)
	retDocPage.EntitiesTrends = entitiesTrends
	// Test
	// idsHasEntity := app.GetIdsFromEntity("Pupil")
	// var timeStampSet map[string]int
	// For each id, find the timestamp
	// for _, id := range idsHasEntity {
	// 	timeStamp, _ := dao.GetTimeStampOnID(id)
	// 	fmt.Println(idsHasEntity, timeStamp.TimeStamp)
	// }
	// Views loading
	templates := template.Must(template.ParseGlob("./templates/*"))
	err = templates.ExecuteTemplate(w, "documentPage", retDocPage)
	// err = templates.ExecuteTemplate(w, "documentPage", newsDoc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	app.GenerateEntitySet()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/topic", topicHandler)
	http.HandleFunc("/document", documentHandler)
	http.ListenAndServe(":8090", nil)

}
