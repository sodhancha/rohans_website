package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sodhancha/rohans_website/router"
)

//TODO
/*
* Make a webserver - DONE
* Make a static file handler - DONE
* Make a Routes handler - DONE
* Make a DB Connector - DONE
* Make a dynamic page which renders SQL output - DONE
* Do a CRUD - DONE
* Make an admin panel where the dynamic content can be updated from - DONE
* Build a RESTFul API for the Backend Content - Skipping and sending to different project with learnings
* Render the API first content in the dynamic web pages  - Skipping and sending to different project with learnings
 */

func main() {
	fmt.Println("Starting Rohans Website")

	//Remember that Go Lang requires exported functions to have Capital Case Names
	router.RoutesHandler()
	router.BuildServer()

}
