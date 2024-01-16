package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//type Animal interface {
//	Eat()
//	Move()
//	Speak()
//}
//
//type Cow struct {
//	name string
//}
//
//func (c Cow) Eat() {
//	fmt.Println("grass")
//}
//func (c Cow) Move() {
//	fmt.Println("walk")
//}
//func (c Cow) Speak() {
//	fmt.Println("moo")
//}
//
//type Bird struct {
//	name string
//}
//
//func (b Bird) Eat() {
//	fmt.Println("worm")
//}
//func (b Bird) Move() {
//	fmt.Println("fly")
//}
//func (b Bird) Speak() {
//	fmt.Println("peep")
//}
//
//type Snake struct {
//	name string
//}
//
//func (s Snake) Eat() {
//	fmt.Println("mice")
//}
//func (s Snake) Move() {
//	fmt.Println("slither")
//}
//func (s Snake) Speak() {
//	fmt.Println("hsss")
//}
//
//func main() {
//	animals := make(map[string]Animal)
//	for {
//		fmt.Print("> ")
//		scanner := bufio.NewReader(os.Stdin)
//		input, _ := scanner.ReadString('\n')
//		input = input[:len(input)-1]
//		request := strings.Split(input, " ")
//		if len(request) != 3 {
//			fmt.Println("Wrong syntax")
//		} else if request[0] == "newanimal" {
//			switch animalType := request[2]; animalType {
//			case "cow":
//				animals[request[1]] = Cow{request[1]}
//			case "bird":
//				animals[request[1]] = Bird{request[1]}
//			case "snake":
//				animals[request[1]] = Snake{request[1]}
//			}
//			fmt.Println("New", request[2], "created successfully")
//		} else if request[0] == "query" {
//			if animal, exist := animals[request[1]]; exist {
//				switch quest := request[2]; quest {
//				case "eat":
//					animal.Eat()
//				case "move":
//					animal.Move()
//				case "speak":
//					animal.Speak()
//				}
//			} else {
//				fmt.Println("No such animal named", request[1], "exists")
//			}
//		} else {
//			fmt.Println("Wrong syntax, Please enter again")
//		}
//	}
//}
