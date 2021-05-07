package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type Meal struct{
	Name string
	Price int
}

type Menu struct{
	Meals []Meal
	Time string
	Chef string  
}

func main(){
	food1 := Meal{
		Name : "Chicken Kasha",
		Price: 234,
	} 
	food2 :=Meal{
		Name: "Chily Chicken",
		Price: 245,
	}
	food3 := Meal{
		Name: "Rooti",
		Price: 2,
	}
	food4 := Meal{
		Name : "Mutton Kasha",
		Price: 300,
	} 
	food5 :=Meal{
		Name: "Fried RIce",
		Price: 200,
	}
	food6 := Meal{
		Name: "Sweets",
		Price: 10,
	}
	food7 := Meal{
		Name : "Butter Chicken",
		Price: 220,
	} 
	food8 :=Meal{
		Name: "Paneer Varta",
		Price: 250,
	}
	food9 := Meal{
		Name: "Tarka",
		Price: 30,
	}

	item := []Menu{
		{
			Meals: []Meal{food1,food4,food7,food6},
			Time: "Dinner",
			Chef: "Sanjeev Kapoor",
		},
		{
			Meals: []Meal{food1,food2,food5,food9},
			Time: "Lunch",
			Chef: "Sanjeev Kapoor",
		},
		{
			Meals: []Meal{food3,food8,food7,food6},
			Time: "BreakFast",
			Chef: "Sanjeev Kapoor",
		},
	}

	err := tpl.Execute(os.Stdout, item)
	if err != nil{
		log.Fatal(err)
	}
}