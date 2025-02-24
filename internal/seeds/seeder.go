package seeds

import (
	"log"
	"reflect"

	"github.com/babyfaceeasy/crims/internal/db"
	"github.com/babyfaceeasy/crims/internal/repository"
)

type Seed struct {
	repo *repository.Repository
}

// Execute will executes the given seeder method
func Execute(seedMethodNames ...string) {
	DB, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatal("Error in connecting to DB - ", err)
	}

	s := Seed{repository.NewRepository(DB)}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Running", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "was successful")
}
