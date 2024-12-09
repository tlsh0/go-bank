package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Askhat", "Tolesh", "12345678")
	seedAccount(s, "John", "Doe", "P@ssw0rd123")
	seedAccount(s, "Emma", "Smith", "Qwerty!987")
	seedAccount(s, "Liam", "Brown", "Hunter2@56")
	seedAccount(s, "Olivia", "Johnson", "Passw0rd!89")
	seedAccount(s, "Noah", "Davis", "Secure@45")
	seedAccount(s, "Sophia", "Miller", "Admin!234")
	seedAccount(s, "James", "Wilson", "1234Abc!")
	seedAccount(s, "Isabella", "Moore", "Zxcvbnm!09")
	seedAccount(s, "Benjamin", "Taylor", "Sunshine#67")
	seedAccount(s, "Mia", "Anderson", "Welcome!123")
	seedAccount(s, "Lucas", "Thomas", "Rocket!42")
	seedAccount(s, "Charlotte", "Jackson", "Flower#99")
	seedAccount(s, "Elijah", "White", "Ocean@21")
	seedAccount(s, "Ava", "Harris", "St@rt123")
	seedAccount(s, "Ethan", "Martin", "Storm!56")
	seedAccount(s, "Amelia", "Thompson", "P@rad1se!")
	seedAccount(s, "Mason", "Garcia", "CoolCat#78")
	seedAccount(s, "Harper", "Martinez", "Ninja!90")
	seedAccount(s, "Logan", "Robinson", "Galaxy@88")
	seedAccount(s, "Evelyn", "Clark", "River#12")

}

func main() {
	seed := flag.Bool("seed", false, "seed the db")

	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the databse")
		//seed stuff
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
