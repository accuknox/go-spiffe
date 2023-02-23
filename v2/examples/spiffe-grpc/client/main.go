package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/accuknox/go-spiffe/v2/spiffegrpc/grpccredentials"
	"github.com/accuknox/go-spiffe/v2/spiffeid"
	"github.com/accuknox/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/accuknox/go-spiffe/v2/workloadapi"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// Workload API socket path
var socketPath = "unix:///tmp/agent.sock"

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {

	saToken, err := os.ReadFile(filepath.Clean("/var/run/secrets/kubernetes.io/serviceaccount/token"))

	if err != nil {
		fmt.Printf("Failed to open serviceaccount token file %v", err)
		return err
	}

	meta := map[string]string{
		"sa_token": string(saToken),
	}

	socketPath = os.Getenv("SPIRE_TCP_ADDR")

	// Create a `workloadapi.X509Source`, it will connect to Workload API using provided socket path
	// If socket path is not defined using `workloadapi.SourceOption`, value from environment variable `SPIFFE_ENDPOINT_SOCKET` is used.
	source, err := workloadapi.NewX509Source(ctx, meta, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)))
	if err != nil {
		return fmt.Errorf("unable to create X509Source: %w", err)
	}
	defer source.Close()

	// Allowed SPIFFE ID
	serverID := spiffeid.RequireFromString(os.Getenv("SPIRE_SERVER_ID"))

	// Dial the server with credentials that do mTLS and verify that presented certificate has SPIFFE ID `spiffe://example.org/server`
	conn, err := grpc.DialContext(ctx, os.Getenv("SERVER_ADDR"), grpc.WithTransportCredentials(
		grpccredentials.MTLSClientCredentials(source, source, tlsconfig.AuthorizeID(serverID)),
	))
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}

	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		return fmt.Errorf("failed issuing RPC to server: %w", err)
	}

	log.Print(reply.Message)
	return nil
}
