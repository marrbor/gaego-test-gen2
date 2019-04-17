/*
 * Copyright (c) 2017 Genetec corporation
 * -*- coding:utf-8 -*-
 *
 * ファイルの説明
 *
 */
package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
	util "github.com/marrbor/gaego-util-gen2"
)

func main() {
	projectID := util.GetProjectID()
	log.Printf("project id: %s\n", projectID)

	dc, err := datastore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dc.Close(); err != nil {
			log.Print(err)
		}
	}()

	http.HandleFunc("/", indexHandler)
	log.Fatal(util.StartServer(8080,nil))
}

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	util.TextResponse(w, "Hello, World!")
}
