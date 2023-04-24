package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/KarimNMG/GradeBookApplication/registery"
)

func main() {
	registery.SetUpRegisteryService()

	http.Handle("/services", &registery.RegisteryService{})

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var srv http.Server
	srv.Addr = registery.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registery Service started. Press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down registery service")
}
