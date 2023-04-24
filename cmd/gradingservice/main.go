package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/KarimNMG/GradeBookApplication/grades"
	"github.com/KarimNMG/GradeBookApplication/log"
	"github.com/KarimNMG/GradeBookApplication/registery"
	"github.com/KarimNMG/GradeBookApplication/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("%s:%s", host, port)

	var r registery.Registration
	r.ServiceName = registery.GradingService
	r.ServiceURL = serviceAddress

	r.RequiredServices = []registery.ServiceName{
		registery.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registery.GetProvider(registery.LogService); err == nil {
		fmt.Println("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
