package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	Port string
	Log  *log.Logger
	Err  *log.Logger
}

func main() {

	ilog := log.New(os.Stdout, "\033[32mINFO:\t\033[0m", log.Ldate)
	elog := log.New(os.Stdout, "\033[31mError:\t\033[0m", log.Ldate)

	app := application{
		Port: "8080",
		Log:  ilog,
		Err:  elog,
	}

	fmt.Printf("Server run on %s\n", app.Port)
	err := http.ListenAndServe(":"+app.Port, app.routes())
	if err != nil {
		app.Err.Println(err)
	}

}
