package main

import (
	"fmt"
	"os"
	"flag"
	"strings"
	"log"
)

func extentions(new, old string) (string, string) {
	newEx := strings.Split(new, ".")
	oldEx := strings.Split(old, ".")
	if len(newEx) != 2 || len(oldEx) != 2 || (newEx[1] != "xml" && newEx[1] != "json") || (oldEx[1] != "xml" && oldEx[1] != "json") {
		log.Fatal("wrong extention")
	}
	return newEx[1], oldEx[1]
}

func compareCakesNames(newCakes, oldCakes []Cake) {
	newNames := make(map[string]int)
	oldNames := make(map[string]int)
	for _, cake := range newCakes {
		newNames[cake.Name]++
	}
	for _, cake := range oldCakes {
		oldNames[cake.Name]++
		if newNames[cake.Name] == 0 {
			fmt.Println("REMOVED cake", "\"" + cake.Name + "\"")
		}
	}
	for name, _ := range newNames {
		if oldNames[name] == 0 {
			fmt.Println("ADDED cake", "\"" + name + "\"")
		}
	}
}

func compareTwoDataBases(newCakes, oldCakes []Cake) {
	compareCakesNames(newCakes, oldCakes)
}

func main() {
	newFlag := flag.String("new", "", "Filename with new database")
	oldFlag := flag.String("old", "", "Filename with original database")
	flag.Parse()
	if *newFlag == "" || *oldFlag == "" {
		os.Exit(1)
	}
	newFilename := *newFlag
	oldFilename := *oldFlag
	newEx, oldEx := extentions(newFilename, oldFilename)
	var new, old DataBase
	if newEx == "xml" {
		new = &XML{}
	} else {
		new = &JSON{}
	}
	if oldEx == "xml" {
		old = &XML{}
	} else {
		old = &JSON{}
	}
	newCakes := new.read(newFilename)
	oldCakes := old.read(oldFilename)
	compareTwoDataBases(newCakes, oldCakes)
}