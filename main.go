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
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	sw "github.com/marcomicera/sayhi/go"
)

type logWriter struct {
}

// Custom log format includes ISO date
func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format(time.RFC3339) + " " + string(bytes))
}

// Launches the web server and takes care of graceful shutdown
func serve(ctx context.Context, port int) (err error) {

	// Using custom log format
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// Starting the web server
	router := sw.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Printf("Server started")
	fmt.Printf("You can override env. variables by editing the \"config.env\" file.\n"+
		"If you defined \"TEST_ENV_VAR\", you will see its value here: %s\n", os.Getenv("TEST_ENV_VAR"))

	<-ctx.Done()

	log.Printf("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown Failed: %+s", err)
	}

	log.Printf("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func main() {

	// Web service port flag
	port := flag.Int("port", 8080, "Web service port")
	flag.Parse()

	// Channel listening to OS signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Context to handle OS interrupt signals
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("System call: %+v", oscall)
		cancel()
	}()

	// Create and run the server with the reference of the context
	if err := serve(ctx, *port); err != nil {
		log.Printf("Failed to serve: %+v\n", err)
	}
}
