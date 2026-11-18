module client-server

go 1.26.0

replace github.com/accuknox/go-spiffe/v2 => ../../

require (
	github.com/accuknox/go-spiffe/v2 v2.2.0
	google.golang.org/grpc v1.85.0-dev.0.20260825072537-93e31b48545e
	google.golang.org/grpc/examples v0.0.0-20261118000413-70d9d2ff7934
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/go-jose/go-jose/v4 v4.1.6 // indirect
	github.com/zeebo/errs v1.4.0 // indirect
	golang.org/x/net v0.63.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260921155816-b14227669459 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)
