package dao

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"regexp"
)

type Twitter struct {
	Text string `bson: "text"`
}

type NewsData struct {
	Title       string `bson:"title" json:"title"`
	TimeStamp   string `bson:"timeStamp" json:"timeStamp"`
	Description string `bson:"description" json:"description"`
	MainStory   string `bson:"mainStory" json:"mainStory"`
}

func GetNewsDataOnID(sid string) (*NewsData, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Connect MongoDB failed")
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	conn := session.DB("gms").C("newsdata")
	// Find
	var result NewsData
	err = conn.FindId(bson.ObjectIdHex(sid)).Select(bson.M{"title": 1, "timeStamp": 1, "description": 1, "mainStory": 1}).One(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	re, _ := regexp.Compile("</?\\w+[^>]*>")
	result.MainStory = re.ReplaceAllString(result.MainStory, "")
	// fmt.Println(result.Title)
	// fmt.Println(result.TimeStamp)
	return &result, nil
}

func GetNewsData() (*[]NewsData, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Connect MongoDB failed")
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	conn := session.DB("gms").C("newsdata")
	// Find
	var result []NewsData
	err = conn.Find(nil).Select(bson.M{"title": 1, "timeStamp": 1, "description": 1, "mainStory": 1}).All(&result)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, v := range result {
		fmt.Println(v.TimeStamp)
	}
	return &result, nil
}

func GetTwitters() (*[]Twitter, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Connect MongoDB failed")
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	conn := session.DB("gms").C("twitter-3")
	// Find
	var result []Twitter
	err = conn.Find(nil).Select(bson.M{"text": 1}).All(&result)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, v := range result {
		fmt.Println(v.Text)
	}

	return &result, nil
}

func main() {
	fmt.Println("...")
	// GetTwitters()
	GetNewsData()
}
