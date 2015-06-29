package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Twitter struct {
	Text string `"bson: "text"`
}

func getTwitters() (*[]Twitter, error) {
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
	getTwitters()
}
