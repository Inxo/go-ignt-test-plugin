package main

import (
	"bitbucket.org/nulljet/go_ignt_func_deps"
	"encoding/json"
	"log"
	"net/http"
	"time"
	functions2 "viads-cron-plugin/src/functions"
)

func Handle(params go_ignt_func_deps.ParamsInterface, _ *http.Request) go_ignt_func_deps.ResponseInterface {
	hello, err := params.Get("hello")
	functions2.CheckErr(err)
	second, err := params.Get("second")
	functions2.CheckErr(err)
	log.Println("cll test_f")
	jsonData, err := json.Marshal("hello " + hello.Value + " second: " + second.Value + " when:" + time.Now().String())
	functions2.CheckErr(err)
	return functions2.JsonResponse{Data: string(jsonData)}
}

func GetParams() []string {
	list := []string{"hello", "second"}
	return list
}

func GetHelp() string {
	return "Test plugin function"
}

func getDescription() string {
	return "Test plugin function"
}

func GetName() string {
	return "test-plugin"
}
