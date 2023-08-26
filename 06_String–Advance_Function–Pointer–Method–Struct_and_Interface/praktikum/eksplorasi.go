package main

import "fmt"

type student struct {
	name string
	nameEncode string
	score int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	key := 10
	for _, value := range s.name{
		ascii := int(value)
		ascii = ascii + key
		if value >= 'a' && value <= 'z' {
			ascii = ((ascii - int('a')) % 26) + int('a')
		} 
		nameEncode += string(rune(ascii))
	}
	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	key := 10
	for _, value := range s.nameEncode {
		ascii := int(value)
		ascii = ascii - key
		if value >= 'a' && value <= 'z' {
			ascii = ((ascii - int('a')) % 26) + int('a')
		}
		nameDecode += string(rune(ascii))
	}
  return nameDecode
}

func main() {
	for {
	var menu int
  	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)
	
	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + " is : " + c.Decode())
	} else if menu == 3 {
		break 
	}
		fmt.Println("\n\n")
	}
}