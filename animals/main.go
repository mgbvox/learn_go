package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	food, noise, locomotion string
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func prompt(msg string) string {
	fmt.Println(msg)
	fmt.Print("> ")
	scn := bufio.NewScanner(os.Stdin)
	if scn.Scan() {
		inp := scn.Text()
		return strings.Trim(strings.ToLower(inp), " ")
	} else {
		return ""
	}
}

func publicMethods(i interface{}) []string {
	var names []string
	typ := reflect.TypeOf(i)
	for j := 0; j < typ.NumMethod(); j++ {
		method := typ.Method(j)
		names = append(names, method.Name)
	}
	return names
}

func getAnimal() Animal {
	inp := prompt("Animal? (q to quit)")
	switch inp {
	case "cow":
		return Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		return Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		return Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	case "q":
		os.Exit(0)
	default:
		fmt.Printf("Invalid input, please try again, got: %s\n", inp)
		return getAnimal()
	}
	panic("Unreachable, added to satisfy the compiler :P")
}

func (a Animal) doMethod() {
	options := strings.Join(publicMethods(a), ", ")
	inp := prompt(fmt.Sprintf("Method? (%s; q to quit)", options))
	if inp == "q" {
		os.Exit(0)
	}
	caser := cases.Title(language.AmericanEnglish)
	methodName := caser.String(inp)
	method := reflect.ValueOf(a).MethodByName(methodName)
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Printf("Invalid input, please try again, got: %s\n", methodName)
		a.doMethod()
	}
}

func main() {
	for {
		getAnimal().doMethod()
	}
}
