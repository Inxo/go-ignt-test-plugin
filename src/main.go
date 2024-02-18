package main

import (
	"bitbucket.org/nulljet/go_ignt_func_deps"
	"bitbucket.org/nulljet/go_ignt_func_deps/functions"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Handle(params go_ignt_func_deps.ParamsInterface, _ *http.Request) go_ignt_func_deps.ResponseInterface {
	hello, err := params.Get("hello")
	functions.CheckErr(err)
	second, err := params.Get("second")
	functions.CheckErr(err)
	log.Println("cll test_f")
	jsonData, err := json.Marshal("hello " + hello.Value + " second: " + second.Value + " when:" + time.Now().String())
	functions.CheckErr(err)
	return functions.JsonResponse{Data: string(jsonData)}
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
