package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/darren-rose/golang-xml/model"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var s model.Specification
	err := envconfig.Process("golang_xml", &s)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if root, err := getXML(fmt.Sprintf("%s?siteCode=%s&operatorCode=%s&country=%s&language=%s", s.Url, s.SiteCode, s.OperatorCode, s.Country, s.Language)); err != nil {
		log.Fatalf("Failed to get XML: %v", err)
		return
	} else {
		filterAfterTime, err := time.Parse("2006-01-02", "2020-02-13")
		if err != nil {
			log.Fatalf("parse time error: %v\n", err)
		} else {
			reduced := root.Medias[:0]
			for _, media := range root.Medias {
				if media.Modified.After(filterAfterTime) {
					reduced = append(reduced, media)
				}
			}
			for _, media := range reduced {
				log.Printf("%v\n", media)
			}
			log.Println("finished")
		}
	}
}

func getXML(url string) (model.Root, error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.Root{}, fmt.Errorf("GET error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Root{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/xml; charset=utf-8" {
		return model.Root{}, fmt.Errorf("Content-Type error: %v", contentType)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Root{}, fmt.Errorf("Read body: %v", err)
	}

	var root model.Root
	xml.Unmarshal(data, &root)

	return root, nil
}
