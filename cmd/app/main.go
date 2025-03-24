package main

import (
	"log"

	"github.com/badprinter/task_manager/internal/app"
)

func main() {
	log.Println("Starting app...")
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	defer app.Close()
	app.Start()
	log.Println("App done.")
}
