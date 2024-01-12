package main

import "fmt"

type Apartment struct {
	Rooms  []Room
	SqArea float64
	Height float64
}

type Room struct {
	Name      string
	SqArea    float64
	Height    float64
	Furniture []string
}

func CreateHome() []Room {
	LivingRoom := Room{
		Name:      "Гостиная",
		SqArea:    20.0,
		Height:    6.0,
		Furniture: []string{"телевизор", "диван", "пианино"},
	}

	BedRoom1 := Room{
		Name:      "Спальня1",
		SqArea:    18.0,
		Height:    3.0,
		Furniture: []string{"кровать", "стол", "стул", "шкаф-купе"},
	}

	BedRoom2 := Room{
		Name:      "Спальня2",
		SqArea:    18.0,
		Height:    3.0,
		Furniture: []string{"кровать", "шкаф-купе", "тумбочка"},
	}

	BedRoom3 := Room{
		Name:      "Спальня3",
		SqArea:    11.0,
		Height:    3.0,
		Furniture: []string{"кровать", "стол", "стул", "шкаф-купе", "тумбочка"},
	}

	Kitchen := Room{
		Name:      "Кухня",
		SqArea:    16.0,
		Height:    6.0,
		Furniture: []string{"стол", "стулья", "шкаф", "духовка", "плита", "холодильник"},
	}

	Apartment := Apartment{
		SqArea: 130.0,
		Height: 3.0,
		Rooms:  []Room{LivingRoom, BedRoom1, BedRoom2, BedRoom3, Kitchen},
	}

	DescribeApartment(Apartment)
	return []Room{LivingRoom, BedRoom1, BedRoom2, BedRoom3, Kitchen}
}

func DescribeApartment(apartment Apartment) {
	fmt.Println("Описание квартиры:")
	fmt.Printf(" Площадь комнат: %.1f квадратных метров\n", apartment.SqArea)
	fmt.Printf(" Высота потолков: %.1f метров\n", apartment.Height)
	fmt.Printf(" Количество комнат: %d\n", len(apartment.Rooms))

	for _, room := range apartment.Rooms {
		fmt.Println("\nОписание комнаты:", room.Name)
		fmt.Printf("Площадь комнаты: %.1f квадратных метров", room.SqArea)
		fmt.Printf("\nВысота потолков: %.1f метров", room.Height)
		fmt.Printf("\nМебель в комнате:\n")

		for _, furniture := range room.Furniture {
			fmt.Printf(" - %s\n", furniture)
		}
	}
}
