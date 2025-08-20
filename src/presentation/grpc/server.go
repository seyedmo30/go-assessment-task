package grpc

import (
	"assessment/di"
	"assessment/pkg"
	"assessment/presentation/grpc/pbs"
	"assessment/presentation/grpc/servers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var grpcServer *grpc.Server

func Start() {
	listener, err := net.Listen("tcp", pkg.Config.GrpcServerAddress)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Panic("failed to make listener for gRPC server")
	}

	grpcServer = grpc.NewServer()
	/* register GRPC servers */
	pbs.RegisterPlanningServiceServer(grpcServer, servers.NewPlanningServer(di.PlanningService()))

	logrus.WithFields(logrus.Fields{
		"grpc_started_at": pkg.Config.GrpcServerAddress,
	}).Info("gRPC server started")

	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Panic("failed to serve gRPC server")
		}
	}()
}

func Stop() {
	// stopping gracefully
	if grpcServer != nil {
		grpcServer.Stop()
	}
}
