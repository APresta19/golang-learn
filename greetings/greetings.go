package greetings

import "fmt"
import "strconv"

type Stack struct {
	Name string
	List []Player
}
type Player struct {
	Name string
	JerseyNumber int8
	Team string
}

func (s *Stack) AddPlayer(player Player) {
	//add new player to list
	s.List = append(s.List, player)
}

func (s *Stack) RemovePlayer(index int) {
	if len(s.List) == 0 {
		fmt.Println("Empty Player List")
		return
	}
	
	left := s.List[:index]
	right := s.List[index+1:]
	left = append(left, right...)
	s.List = left
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v.Welcome!", name)
	return message
}

func HelloInt(val int) string {
	var valStr string = strconv.Itoa(val); 
	message := fmt.Sprintf("Num is %v", valStr)
	return message
}

func AddNums(num1 int, num2 int) int {
	return num1 + num2
}

func ReverseString(str string) string {
	var strConcat string;
	for i := len(str) - 1; i >= 0; i-- {
		strConcat += string(str[i])
	}
	return strConcat
}

func Max(arr []int) int {
	var max int = arr[0]
	for _, val := range arr { // _ for no index in code
		if val > max {
			max = val
		}
	}
	return max
}