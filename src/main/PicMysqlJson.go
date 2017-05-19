package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/emirpasic/gods/sets/hashset"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// CREATE TABLE `pic` (
//   `id` int(11) NOT NULL AUTO_INCREMENT,
//   `title` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
//   `type` varchar(50) COLLATE utf8_unicode_ci NOT NULL,
//   `imgUrl` varchar(1024) COLLATE utf8_unicode_ci NOT NULL,
//   PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
type Pic struct {
	ID     int    `json:"id"`
	Title  string `json:"title" form:"title"`
	Type   string `json:"type" form:"type"`
	ImgURL string `json:"imgUrl" form:"imgUrl"`
}

func (p Pic) add() (ID int, err error) {
	stmt, err := db.Prepare("INSERT INTO pic(title, type, imgUrl) VALUES (?, ?, ?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.Title, p.Type, p.ImgURL)
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	ID = int(id)
	defer stmt.Close()
	return
}

// https://github.com/emirpasic/gods
func main() {
	var err error
	db, err = sql.Open("mysql", "root:root123xw@tcp(127.0.0.1:3306)/goginapi?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	///Users/xiwi/Documents/design/pic/emotions/json.json
	// file, e := ioutil.ReadFile("/Users/xiwi/Documents/design/pic/new/json.json")
	// file, e := ioutil.ReadFile("/Users/xiwi/Documents/design/pic/emotions/json.json")
	file, e := ioutil.ReadFile("/Users/xiwi/Documents/design/pic/best/json.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	// fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsonPics []Pic
	json.Unmarshal(file, &jsonPics)
	logger := log.New(os.Stdout, "log id=", log.Lshortfile)
	for _, pic := range jsonPics {
		ID, _ := pic.add()
		fmt.Println(ID)
		// logger.Println(pic)
	}

	set := hashset.New()   // empty
	set.Add(1)             // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	logger.Println(set)
	set.Remove(4)      // 5, 3, 2, 1 (random order)
	set.Remove(2, 3)   // 1, 5 (random order)
	set.Contains(1)    // true
	set.Contains(1, 5) // true
	set.Contains(1, 6) // false
	_ = set.Values()   // []int{5,1} (random order)
	set.Clear()        // empty
	set.Empty()        // true
	set.Size()         // 0

}
