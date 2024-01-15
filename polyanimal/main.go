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
	terminus := ">"
	switch len(args) {
	case 1:
		msg = args[0]
	case 2:
		msg = args[0]
		quitseq = args[1]
	}

	formatted := strings.TrimLeft(fmt.Sprintf("%s (%s to quit) %s ", msg, quitseq, terminus), " ")
	fmt.Print(formatted)
	scn := bufio.NewScanner(os.Stdin)
	result := ""
	if scn.Scan() {
		inp := scn.Text()
		clean := strings.Trim(strings.ToLower(inp), " ")
		if clean == quitseq && quitseq != "" {
			fmt.Println("Quitting")
			os.Exit(0)
		} else {
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

func doMethod(a Animal, command string) {
	caser := cases.Title(language.AmericanEnglish)
	methodName := caser.String(command)
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

func main() {

	animals := make(map[string]Animal)

	for {
		inp := strings.Split(prompt("command", "q"), " ")
		if len(inp) != 3 {
			fmt.Println("Must provide three whitespace-separated commands, e.g. 'newanimal bob cow'")
		} else {

			mode := inp[0]
			name := inp[1]
			command := inp[2]

			switch mode {
			case "newanimal":
				makeAnimal(name, command, &animals)
			case "query":
				invokeAnimal(name, command, &animals)
			default:
				fmt.Printf("Invalid option '%s'\n", mode)
			}
		}
	}
}
