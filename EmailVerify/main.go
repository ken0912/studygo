package main

import (
	"log"

	trumail "github.com/sdwolfe32/trumail/verifier"
)

func main() {
	v := trumail.NewVerifier("salary.com", "ken.shi@salary.com")

	// Validate a single address
	log.Println(v.Verify("kensdfdsfsdfdsfds.shi@salary.com"))
}
