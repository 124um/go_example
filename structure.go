package main

import (
	"fmt"
)

type User struct {
	Name  string
	Age   Age
	Weith int
	Sex   string
}

type Age int

func (a Age) isAdult() bool {
	return a >= 18 
}

type DumpDatabase struct{
	m map[string]string
}

//constructor initial db, так как  в пустую мапу или слайс нельзя добавлять ничего нового, 
//потому что по сути его размер равен 0 - после инициализации рабочая
func NewDatabase() *DumpDatabase{
	return &DumpDatabase{
		m: make(map[string]string),
	}
}

// функция структури
func (u User) printUserInfo(){
	fmt.Println(u.Age, u.Name, u.Sex, u.Weith)
}

// функция структури - работаем не с копией, а с областью памяти
func (u *User) printUserInfoR(name string){
	u.Name = name
	fmt.Println(u.Age, u.Name, u.Sex, u.Weith)
}

// функция структури
func (u User) getName()string{
	return u.Name
}

// функция структури - работаем не с копией, а с областью памяти
func (u *User) setName(name string){
	u.Name = name
}

// constructor
func NewUser(name string, age int, weith int, sex string) User {
	return User{
		Name: name,
		Age: Age(age),
		Weith: weith,
		Sex: sex,
	}
}

func mainStr() {
	user := NewUser("Bufe", 5, 54, "female")
	user2 := NewUser("Rommel", 65, 77, "male")
	user3 := NewUser("Huemel", 44, 58, "male")
	user4 := NewUser("Bukake", 22, 99, "female")

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user3)
	fmt.Printf("%+v\n", user4)

	fmt.Printf("%+v\n", user.Age)
	fmt.Printf("%+v\n", user2.Name)
	fmt.Printf("%+v\n", user3.Sex)
	fmt.Printf("%+v\n", user4.Weith)

	user.printUserInfo()
	user2.printUserInfoR("Bibs")
	user3.printUserInfoR("Buubs")

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", user2)
	fmt.Printf("%+v\n", user3)

	
	user.setName("NewName")
	fmt.Printf("%+v\n", user.getName())

	fmt.Printf("%+v\n", user.Age.isAdult())
	fmt.Printf("%+v\n", user2.Age.isAdult())

}
