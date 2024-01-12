package main

import "fmt"

type Family struct {
	Members []FamMember
	Pets    []string
}

type FamMember struct {
	Name string
}

func CreateFamMembers() {
	FamilyMembers := []FamMember{
		{Name: "Я"},
		{Name: "Мама"},
		{Name: "Папа"},
		{Name: "Брат"},
		{Name: "Сестра"},
	}

	pets := []string{"Кошка"}
	family := Family{
		Members: FamilyMembers,
		Pets:    pets,
	}
	DescribeFamMembers(family)
}

func DescribeFamMembers(family Family) {
	fmt.Println("\nСемья:")
	for _, member := range family.Members {
		fmt.Println("-", member.Name)
	}

	fmt.Println("\nДомашнее животное:")
	for _, pet := range family.Pets {
		fmt.Println("-", pet)
	}
}
