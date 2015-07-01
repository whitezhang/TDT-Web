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

type SimpleNewsData struct {
	ID          string
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	TimeStamp   string `bson:"timeStamp" json:"timeStamp"`
}

type NewsData struct {
	ID          string
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	TimeStamp   string `bson:"timeStamp" json:"timeStamp"`
	Category    string `bson:"category" json:"category"`
	URL         string `bson:"url" json:"url"`
	Source      string `bson:"source" json:"source"`
	MainStory   string `bson:"mainStory" json:"mainStory"`
}

func GetSimpleNewsDataOnID(sid string) (*SimpleNewsData, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Connect MongoDB failed")
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	conn := session.DB("gms").C("newsdata")
	// Find
	var result SimpleNewsData
	err = conn.FindId(bson.ObjectIdHex(sid)).Select(bson.M{"title": 1, "description": 1, "timeStamp": 1}).One(&result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result.ID = sid
	// fmt.Println(result.Title)
	// fmt.Println(result.TimeStamp)
	return &result, nil
}

/*
 * Usage: get the whole newsdata
 */
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
	err = conn.FindId(bson.ObjectIdHex(sid)).One(&result)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	re, _ := regexp.Compile("</?\\w+[^>]*>")
	result.MainStory = re.ReplaceAllString(result.MainStory, "")
	result.ID = sid

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

// func main() {
// fmt.Println("...")
// GetTwitters()
// GetNewsData()
// }
