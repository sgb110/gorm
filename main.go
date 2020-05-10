package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	gorm.Model
	Name   string
	Age    uint
	Phones []Phone `gorm:"forgin_key:PersonID"`
}
type Phone struct {
	gorm.Model
	PhoneNumber string
	PersonID    uint
}

func main() {

	db, err := gorm.Open("sqlite3", "mydb.db")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(400)

	}
	defer db.Close()

	db.AutoMigrate(&Person{}, &Phone{})
	var row Person
	//db.Create(&Person{Age: 51, Name: "Gabriel"})
	//db.Create(&Person{Age: 34, Name: "Reza"})
	//db.Create(&Phone{PersonID: 1, PhoneNumber: "09215620525"})

	x := db.Model(&Person{}).Find(&row, "id = 1").Related(&row.Phones)

	// x := db.Model(&Person{}).Where(&Person{Name: "Ali"})
	rows, _ := x.Rows()
	//fmt.Println(err2)

	for rows.Next() {
		db.ScanRows(rows, &row)
		//rows.Scan(%row)
		fmt.Println(row)
	}
	fmt.Println("1----------")
	// y := db.Model(&Person{}).Where("age > 50")
	// //x.First(&row, "Age > 50")
	// rows2, _ := y.Rows()
	// fmt.Println("----------")
	// for rows2.Next() {
	// 	db.ScanRows(rows2, &row)
	// 	//rows.Scan(%row)
	// 	fmt.Println(row)
	// }
	//fmt.Println(row)

	fmt.Println("2----------")

	_ = db.Model(&Person{}).First(&row)
	fmt.Println(row)

}
