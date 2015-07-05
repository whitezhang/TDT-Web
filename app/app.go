package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const indices2id_file = "../plsa/file-path.txt"
const entity_path = "/Users/wyatt/Documents/Code/Gla/Final/Sources/web/db/gms/r_month-4/"

type EntityNode struct {
	ExpEntity []string
}

type EntitySet struct {
	ExpEntityNode map[string]EntityNode
}

var ExpEntitySet EntitySet

/*
 * Usage: find the id according to the index
 * e.g.
 *		0 -> 5536def4e4b0644323e219a8
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

func SplitDate(date string) (int, int, int) {
	info := strings.Split(date, "/")
	day, _ := strconv.Atoi(info[0])
	month, _ := strconv.Atoi(info[1])
	year, _ := strconv.Atoi(info[2])
	return day, month, year
}

func GetIdsFromEntity(targetEntity string) []string {
	var ids []string
	for sid, eNode := range ExpEntitySet.ExpEntityNode {
		for _, entityName := range eNode.ExpEntity {
			if entityName == targetEntity {
				ids = append(ids, sid)
				break
			}
		}
	}
	return ids
}

func GenerateEntitySet() error {
	ExpEntitySet.ExpEntityNode = make(map[string]EntityNode)
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
			ExpEntitySet.ExpEntityNode[sid] = EntityNode{ExpEntity: entityNames}
			// fmt.Printf("%s Entity Set: %s\n", sid, ExpEntitySet.ExpEntityNode)
		}
		return nil
	})
	fmt.Println("Have generated the entity set")
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
