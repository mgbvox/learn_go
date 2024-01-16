package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Cow) Move() {
	fmt.Println("walk")
}

type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}

func (c Bird) Speak() {
	fmt.Println("peep")
}
func (c Bird) Move() {
	fmt.Println("fly")
}

type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Snake) Speak() {
	fmt.Println("hsss")
}
func (c Snake) Move() {
	fmt.Println("slither")
}

func prompt(args ...string) string {
	var msg string
	quitseq := "ctrl+c"
	helpseq := "h"
	helptxt := ""
	terminus := ">"

	for idx, arg := range args {
		if arg != "" {
			switch idx {
			case 0:
				msg = arg
			case 1:
				quitseq = arg
			case 2:
				helpseq = arg
			case 3:
				helptxt = arg
			case 4:
				terminus = arg

			}
		}

	}

	formatted := strings.TrimLeft(fmt.Sprintf("%s (%s to quit, %s for help) %s ", msg, quitseq, helpseq, terminus), " ")
	fmt.Print(formatted)
	scn := bufio.NewScanner(os.Stdin)
	result := ""
	if scn.Scan() {
		inp := scn.Text()
		clean := strings.Trim(strings.ToLower(inp), " ")
		switch clean {
		case quitseq:
			if quitseq != "" {
				fmt.Println("Quitting")
				os.Exit(0)
			}
		case helpseq:
			fmt.Println(helptxt)
			return prompt(args...)
		default:
			result = clean
		}

	}
	return result
}

func makeAnimal(name, animalType string, animap *map[string]Animal) {
	var animal Animal
	switch animalType {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}

	default:
		fmt.Printf("Invalid animal type '%s', please try again.\n", animalType)
	}

	(*animap)[name] = animal

}

func Map[T any](vs []T, f func(T) T) []T {
	vsm := make([]T, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Title(s string) string {
	runes := []rune(s)
	firstUpper := []rune{unicode.ToUpper(runes[0])}
	rest := Map(runes[1:], unicode.ToLower)
	all := append(firstUpper, rest...)
	return string(all)

}

func doMethod(a Animal, command string) {
	methodName := Title(command)
	method := reflect.ValueOf(a).MethodByName(methodName)
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Printf("Invalid method, please try again, got: %s\n", methodName)
	}
}

func invokeAnimal(name, command string, animap *map[string]Animal) {
	animal := (*animap)[name]
	doMethod(animal, command)
}

func printAnimals(animals *map[string]Animal) {
	if len(*animals) > 0 {
		fmt.Println("Animals: ")
	}
	for k, v := range *animals {
		fmt.Println(k, "->", reflect.TypeOf(v))
	}
}

func main() {

	animals := make(map[string]Animal)

	for {

		inp := strings.Split(prompt("command", "q", "h", "Sample Commands: newanimal <name> <cow|bird|snake> | query <name> <eat|move|speak>"), " ")
		if len(inp) != 3 {
			fmt.Println("Must provide three whitespace-separated commands, e.g. 'newanimal bob cow'")
		} else {

			mode := inp[0]
			name := inp[1]
			command := inp[2]

			switch mode {
			case "newanimal":
				makeAnimal(name, command, &animals)
				printAnimals(&animals)
			case "query":
				invokeAnimal(name, command, &animals)
			default:
				fmt.Printf("Invalid option '%s'\n", mode)
			}
		}
	}
}
