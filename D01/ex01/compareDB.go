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

func compareCakesNames(newCakes, oldCakes []Cake) []string{
	commonNames := make([]string, 0)
	newNames := make(map[string]int)
	oldNames := make(map[string]int)
	for _, cake := range newCakes {
		newNames[cake.Name]++
	}
	for _, cake := range oldCakes {
		oldNames[cake.Name]++
		if newNames[cake.Name] == 0 {
			fmt.Printf("REMOVED cake %q\n", cake.Name)
		} else {
			commonNames = append(commonNames, cake.Name)
		}
	}
	for name, _ := range newNames {
		if oldNames[name] == 0 {
			fmt.Printf("ADDED cake %q\n", name)
		}
	}
	return commonNames
}

func findCake(cakes []Cake, name string) Cake {
	for _, cake := range cakes {
		if cake.Name == name {
			return cake
		}
	}
	return cakes[0]
}

func compareTwoIngredients(cakeName string, oldItem, newItem Ingredient) {
	name := oldItem.Name
	if oldItem.Count != newItem.Count {
		fmt.Printf("CHANGED unit count for ingredient %q for cake  %q - %q instead of %q\n",
					name, cakeName, newItem.Count, oldItem.Count)
	}
	if oldItem.Unit != newItem.Unit {
		if len(oldItem.Unit) == 0 {
			fmt.Printf("ADDED unit %q for ingredient %q for cake  %q\n", newItem.Unit, name, cakeName)
		} else if len(newItem.Unit) == 0 {
			fmt.Printf("REMOVED unit %q for ingredient %q for cake  %q\n", oldItem.Unit, name, cakeName)
		} else {
			fmt.Printf("CHANGED unit for ingredient %q for cake  %q - %q instead of %q\n",
						name, cakeName, newItem.Unit, oldItem.Unit)
		}
	}
}

func compareIngredientsNames(newCake, oldCake Cake) []string {
	cakeName := newCake.Name
	commonIngredients := make([]string, 0)
	newIngredients := make(map[string]int)
	oldIngredients := make(map[string]int)
	for _, item := range newCake.Ingredients {
		newIngredients[item.Name]++
	}
	for _, item := range oldCake.Ingredients {
		oldIngredients[item.Name]++
		if newIngredients[item.Name] == 0 {
			fmt.Printf("REMOVED ingredient %q for cake  %q\n", item.Name, cakeName)
		} else {
			commonIngredients = append(commonIngredients, item.Name)
		}
	}
	for name, _ := range newIngredients {
		if oldIngredients[name] == 0 {
			fmt.Printf("ADDED ingredient %q for cake  %q\n", name, cakeName)
		}
	}
	return commonIngredients
}

func findIngredient(items []Ingredient, name string) Ingredient {
	for _, item := range items {
		if item.Name == name {
			return item
		}
	}
	return items[0]
}

func compareTwoCakes(newCake, oldCake Cake) {
	cakeName := newCake.Name
	if newCake.Time != oldCake.Time {
		fmt.Printf("CHANGED cooking time for cake %q - %q instead of %q\n", cakeName, newCake.Time, oldCake.Time)
	}
	commonIngredients := compareIngredientsNames(newCake, oldCake)
	for _, item := range commonIngredients {
		oldItem := findIngredient(oldCake.Ingredients, item)
		newItem := findIngredient(newCake.Ingredients, item)
		compareTwoIngredients(cakeName, oldItem, newItem)
	}
}

func compareTwoDataBases(newCakes, oldCakes []Cake) {
	commonNames := compareCakesNames(newCakes, oldCakes)
	for _, name := range commonNames {
		oldCake := findCake(oldCakes, name)
		newCake := findCake(newCakes, name)
		compareTwoCakes(newCake, oldCake)
	}
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
