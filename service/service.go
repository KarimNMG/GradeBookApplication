package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/KarimNMG/GradeBookApplication/registery"
)

func Start(
	ctx context.Context,
	host,
	port string,
	req registery.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, req.ServiceName, host, port)
	err := registery.RegisterService(req)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(
	ctx context.Context,
	serviceName registery.ServiceName,
	host,
	port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registery.ShuttDownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Printf("Error: %v", err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
