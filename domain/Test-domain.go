package domain

import (
	proto "Test-output/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	mapstructure "github.com/mitchellh/mapstructure"
	"log"
)

var engine *xorm.Engine
var conn = "root:@tcp(localhost:3306)/GOGENETIC_SCHEMA?charset=utf8&parseTime=True&loc=Local"

type ToDoDomains struct {
	Sno         int    `mapstructure:"Id"`
	Description string `mapstructure:"Title"`
}

func ReadAllToDoDomains() []ToDoDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", conn)

	var toDoDomains []ToDoDomains
	engine.Find(&toDoDomains)
	log.Println("{}", toDoDomains)

	if err != nil {
		log.Println(err)
	}

	return toDoDomains

}

func ReadToDoDomains(Sno int) ToDoDomains {
	var err error
	engine, err = xorm.NewEngine("mysql", conn)

	var toDoDomains = ToDoDomains{Sno: Sno}
	has, err := engine.Get(&toDoDomains)
	log.Println("{}", toDoDomains)
	if err != nil {
		log.Println(err)
	}
	log.Println(has)

	return toDoDomains

}

func CreateToDoDomains(toDoDomains ToDoDomains) {

	var err error
	engine, err = xorm.NewEngine("mysql", conn)
	engine.Insert(&toDoDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Created {}", &toDoDomains)
	}

}

func DeleteToDoDomains(Sno int) {
	var err error
	engine, err = xorm.NewEngine("mysql", conn)

	var toDoDomains = ToDoDomains{Sno: Sno}
	engine.Delete(&toDoDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Deleted {}", &toDoDomains)
	}

}

func UpdateToDoDomains(Sno int, toDoDomains ToDoDomains) {
	var err error
	engine, err = xorm.NewEngine("mysql", conn)

	engine.Update(&toDoDomains)

	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully Updated {}", &toDoDomains)
	}
}

func ConvertToDoDomains2ToDo(tododomains interface{}) (todo *proto.ToDo) {
	err0 := mapstructure.Decode(tododomains, &todo)
	if err0 != nil {
		panic(err0)
	}
	return todo
}

func ConvertToDo2ToDoDomains(todo *proto.ToDo) (tododomains ToDoDomains) {
	err1 := mapstructure.Decode(todo, &tododomains)
	if err1 != nil {
		panic(err1)
	}
	return tododomains
}
