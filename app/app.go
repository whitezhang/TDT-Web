package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const indices2id_file = "../plsa/file-path.txt"
const entity_path = "/Users/wyatt/Documents/Code/Gla/Final/Sources/web/db/gms/r_month-4/"

type EntityNode struct {
	entity []string
}

type EntitySet struct {
	entityNode map[string]EntityNode
}

var entitySet EntitySet

/*
 * Usage: find the id according to the index
 */
func Index2Id(index int) (string, error) {
	fin, err := os.Open(indices2id_file)
	defer fin.Close()
	if err != nil {
		panic(err)
		return "", err
	}
	reader := bufio.NewReader(fin)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		if i == index {
			info := strings.Split(line, "/")
			return info[len(info)-1], nil
		}
		i++
	}
	return "", err
}

func GetIdsHasEntity(targetEntity string) []string {
	var ids []string
	for sid, eNode := range entitySet.entityNode {
		for _, entityName := range eNode.entity {
			if entityName == targetEntity {
				ids = append(ids, sid)
				break
			}
		}
	}
	return ids
}

func GenerateEntitySet() error {
	err := filepath.Walk(entity_path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			fmt.Println("Error:", err)
			return err
		}
		if info.IsDir() {
			// EntityConterOnEntity(path)
			// fmt.Println("dir")
		} else {
			// fmt.Println(path)
			fin, err := os.Open(path)
			defer fin.Close()
			if err != nil {
				panic(err)
				return err
			}
			_sid := strings.Split(path, "/")
			sid := _sid[len(_sid)-1]
			reader := bufio.NewReader(fin)
			// Read file(sid)
			var entityNames []string
			for {
				line, err := reader.ReadString('\n')
				if err != nil || io.EOF == err {
					break
				}
				line = strings.Replace(line, "\n", "", -1)
				// Find entities
				re := regexp.MustCompile(`\/resource\/(\w*)\" title`)
				entity := re.FindAllStringSubmatch(line, -1)

				for _, v := range entity {
					entityNames = append(entityNames, v[1])
				}
			}
			entitySet.entityNode = make(map[string]EntityNode)
			entitySet.entityNode[sid] = EntityNode{entity: entityNames}
			fmt.Printf("%s Entity Set: %s\n", sid, entitySet.entityNode[sid])
		}
		return nil
	})
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

// func EntityParserOnID(sid string) {
// 	entity_file := entity_path + sid
// 	fin, err := os.Open(entity_file)
// 	defer fin.Close()
// 	if err != nil {
// 		panic(err)
// 		return "", err
// 	}
// 	reader := bufio.NewReader(fin)

// }
