package main

import (
	"auth-and-users/repo"
)

func main() {
	repo.Connect("")
	repo.Migrate()
}
