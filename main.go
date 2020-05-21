/*
 * sayhi web service
 *
 * A polite web service that always says hello
 *
 * API version: 1.0.0
 * Contact: marco.micera+sayhi@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	sw "github.com/marcomicera/sayhi/go"
)

type logWriter struct {
}

// Custom log format includes ISO date
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format(time.RFC3339) + " " + string(bytes))
}

func main() {

	// Using custom log format
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	log.Printf("Server started")
	fmt.Printf("You can override env. variables by editing the \"config.env\" file.\n"+
		"If you defined \"TEST_ENV_VAR\", you will see its value here: %s\n", os.Getenv("TEST_ENV_VAR"))

	router := sw.NewRouter()

	port := flag.Int("port", 8080, "Web service port")
	flag.Parse()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), router))
}
