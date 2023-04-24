package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/KarimNMG/GradeBookApplication/log"
	"github.com/KarimNMG/GradeBookApplication/registery"
	"github.com/KarimNMG/GradeBookApplication/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	var r registery.Registration
	r.ServiceName = registery.LogService
	r.ServiceURL = serviceAddress
	r.RequiredServices = make([]registery.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandler,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
