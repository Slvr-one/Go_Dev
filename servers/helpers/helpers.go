package helpers

// https://www.blog.duomly.com/golang-course-with-building-a-fintech-banking-app-lesson-1-start-the-project/

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HandleErr(err error, msg string) {
	if err != nil {
		log.Fatalf("Error: %s", msg)
		panic(err.Error())

	}
}

// HashAndSalt take a pass as bytes, hashes & saltes it.
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err, "err on hash & salting pass")

	return string(hashed)
}
