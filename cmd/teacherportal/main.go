package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/KarimNMG/GradeBookApplication/log"
	"github.com/KarimNMG/GradeBookApplication/registery"
	"github.com/KarimNMG/GradeBookApplication/service"
	"github.com/KarimNMG/GradeBookApplication/teacherportal"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("%s:%s", host, port)

	var r registery.Registration
	r.ServiceName = registery.TeacherPortal
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registery.ServiceName{
		registery.LogService,
		registery.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		teacherportal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registery.GetProvider(registery.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")
}
