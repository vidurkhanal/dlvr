package main

import (
	"auth-and-users/repo"
)

func main() {
	repo.Connect("postgres://lkbtxmfjlzhkaw:9f133ef2f7cfc968126e983141ce623e88f941ea68205eea2c6578374a1bfe44@ec2-54-87-179-4.compute-1.amazonaws.com:5432/da75tbdvv6c9p6")
	repo.Migrate()
}
