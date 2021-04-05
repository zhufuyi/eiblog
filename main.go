package main

import (
	"eiblog/setting"
	"eiblog/utils/mgo"

	"fmt"
)

func main() {
	initServer()
	Run()
}

func initServer() {
	mongodbAddr := setting.Conf.DbServer.MongodbAddr
	elasticsearchAddr := setting.Conf.DbServer.ElasticsearchAddr
	// 初始化mongodb
	fmt.Printf("mongodbAddr = %s\nelasticsearchAddr = %s\n", mongodbAddr, elasticsearchAddr)
	err := mgo.Init(mongodbAddr)
	if err != nil {
		panic("connetc mongodb failed, " + err.Error())
	}
	fmt.Println("connetc mongodb success")

	// 初始化es(非必需)
	err = InitElasticsearch(elasticsearchAddr)
	if err != nil {
		fmt.Printf("connect elasticsearch failed, err = %s\n", err.Error())
	} else {
		fmt.Println("connect elasticsearch success")
	}

	InitDB()
	PingDomain()
	initXML()
	initWebServer()
}
