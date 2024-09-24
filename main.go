// package main seeds an itpg database with fake data.
package main

import (
	"context"
	"flag"
	"log"
	"math/rand"

	"github.com/jaswdr/faker"
	"github.com/vanillaiice/itpg/db"
	"github.com/vanillaiice/itpg/db/postgres"
	"github.com/vanillaiice/itpg/db/sqlite"
)

// maxRandScore is the maximum score that can be randomly generated.
const maxRandScore = 6

func main() {
	dbBackend := flag.String("db-backend", "sqlite", "database backend to use (sqlite, postgres)")
	dbUrl := flag.String("db-url", "itpg.db", "database url")
	codeLen := flag.Int("code-len", 5, "length of generated course codes")
	sampleSize := flag.Int("sample-size", 50, "number of sample data to generate")

	flag.Parse()

	var d db.DB
	var err error

	switch *dbBackend {
	case "sqlite":
		d, err = sqlite.New(*dbUrl, "", 0, context.Background())
	case "postgres":
		d, err = postgres.New(*dbUrl, "", 0, context.Background())
	default:
		log.Fatalf("backend %s not supported", *dbBackend)
	}

	if err != nil {
		log.Fatal(err)
	}

	faker := faker.New()

	var courses []*db.Course
	for i := 0; i < *sampleSize; i++ {
		c := &db.Course{Name: faker.Company().CatchPhrase(), Code: faker.RandomStringWithLength(*codeLen)}
		courses = append(courses, c)
	}
	if err = d.AddCourseMany(courses); err != nil {
		log.Fatal(err)
	}

	var names []string
	for i := 0; i < *sampleSize; i++ {
		names = append(names, faker.Person().Name())
	}
	if err = d.AddProfessorMany(names); err != nil {
		log.Fatal(err)
	}

	var professors []*db.Professor
	professors, err = d.GetLastProfessors()
	if err != nil {
		log.Fatal(err)
	}

	if len(professors) > *sampleSize {
		professors = professors[:*sampleSize]
	}

	for i, p := range professors {
		randScores := [3]float32{float32(rand.Intn(maxRandScore)), float32(rand.Intn(maxRandScore)), float32(rand.Intn(maxRandScore))}
		err = d.GradeCourseProfessor(p.UUID, courses[i].Code, faker.Person().Name(), randScores)
		if err != nil {
			log.Fatal(err)
		}
	}
}
