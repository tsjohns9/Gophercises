package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Phone is a phone
type Phone struct {
	gorm.Model
	Number uint64
	Name   string
}

// JSONPhone is json
type JSONPhone struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

var db *gorm.DB

func main() {
	file := flag.String("file", "phones.json", "JSON file of phone numbers to normalize")
	flag.Parse()

	defer db.Close()
	db.AutoMigrate(&Phone{})

	phones := openFile(*file)
	loopPhones(phones)
}

func loopPhones(phones []JSONPhone) {
	for _, p := range phones {
		uintNumber, isOk := normalize(p)
		savedPhone := find(uintNumber)
		fmt.Println(savedPhone)
		if savedPhone.Number == 0 && isOk {
			create(uintNumber, p.Name)
			return
		}

		if savedPhone.Number != 0 && isOk {
			update(savedPhone.ID, uintNumber, p.Name)
		}
	}
}

func create(num uint64, name string) {
	db.Create(&Phone{Number: num, Name: name})
}

func find(num uint64) Phone {
	var phone Phone
	db.Where(&Phone{Name: "", Number: num}).First(&phone)
	return phone
}

func update(id uint, num uint64, name string) Phone {
	phone := Phone{Name: name, Number: num}
	var p Phone
	db.Model(p).Where("id = ?", id).Update(&phone)
	return phone
}

func openFile(fileName string) []JSONPhone {
	var data []JSONPhone

	bts, err := ioutil.ReadFile(fileName)
	checkErr(err)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(bts, &data)
	return data
}

func normalize(p JSONPhone) (uint64, bool) {
	var runes []rune
	for _, ch := range p.Number {
		if ch >= '0' && ch <= '9' {
			runes = append(runes, ch)
		}
	}
	strNumber := string(runes)
	isOk := len(strNumber) == 10
	x, _ := strconv.Atoi(strNumber)
	return uint64(x), isOk
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:password1@/golang_phones?charset=utf8&parseTime=True&loc=Local")
	checkErr(err)
}
