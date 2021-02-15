package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func pickRandomFood(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// need to scan the file, not just read it
	scanner := bufio.NewScanner(file)
	// need some RNG
	randsource := rand.NewSource(time.Now().UnixNano())
	randgenerator := rand.New(randsource)

	lineNum := 1
	var pick string
	for scanner.Scan() {
		line := scanner.Text()
		// Instead of 1 to N it's 0 to N-1
		roll := randgenerator.Intn(lineNum)
		if roll == 0 {
			pick = line
		}
		lineNum += 1
	}

	return pick
}

func makeMeal() (meal []string, err error) {

	meal = make([]string, 0)

	// get the paths of the config files
	m, _ := filepath.Abs("config/meats.txt")
	v, _ := filepath.Abs("config/veggies.txt")
	o, _ := filepath.Abs("config/oils.txt")
	s, _ := filepath.Abs("config/spices.txt")
	if err != nil {
		return
	}
	// make the meal
	meat := pickRandomFood(m)
	veggie := pickRandomFood(v)
	oil := pickRandomFood(o)
	spice := pickRandomFood(s)
	// append the choices to the meal
	meal = append(meal, meat, veggie, oil, spice)
	// return the meal
	return meal, nil

}

func main() {

	fmt.Println("\n-----------------------------------------------------")
	fmt.Println("\n-------------------Paleo Diet Matrix-----------------")
	// make the meal
	result, _ := makeMeal()
	fmt.Println("\nHere's your meal! Get a frying pan. Add the following:")
	// loop through the items in the meal
	for _, ingredient := range result {
		// print them out. this is what you will eat!
		fmt.Println(ingredient)
	}
	fmt.Println("\nFry until your protein is cooked. Eat. Enjoy!")
	fmt.Println("\n----------------------------------------------------")
}
