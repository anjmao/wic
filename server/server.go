package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	grpcServer *grpc.Server
}

var (
	crt = "certs/server.crt"
	key = "certs/server.key"
	ca  = "certs/ca.crt"
)

// New creates new GRPC server instance.
func New() (*Server, error) {
	// Create new logrus entry and replace standart grpc logger to user logrus.
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpclogrus.Option{
		grpclogrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}
	grpclogrus.ReplaceGrpcLogger(logrusEntry)
	unaryLogger := grpclogrus.UnaryServerInterceptor(logrusEntry, opts...)
	streamingLogger := grpclogrus.StreamServerInterceptor(logrusEntry, opts...)

	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, fmt.Errorf("cannot load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority.
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, fmt.Errorf("cannot read ca certificate: %s", err)
	}

	// Append the client certificates from the CA.
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append client certs")
	}

	// Create the TLS credentials.
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS12,
	})

	// Initialize underlying grpc server with middlewares.
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		middleware.WithUnaryServerChain(
			unaryLogger,
			grpcrecovery.UnaryServerInterceptor(),
		),
		middleware.WithStreamServerChain(
			streamingLogger,
			grpcrecovery.StreamServerInterceptor(),
		),
	)
	srv := &Server{grpcServer: grpcServer}
	return srv, nil
}

// Register registers grpc services.
func (s *Server) Register(register func(grpcServer *grpc.Server)) {
	register(s.grpcServer)
}

// Run starts server on given port.
func (s *Server) Run(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("cannot listen on port %d: %v", port, err)
	}
	if err := s.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("cannot server and accept connections: %v", err)
	}
	return nil
}
