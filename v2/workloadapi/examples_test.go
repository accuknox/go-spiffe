package workloadapi_test

import (
	"context"

	"github.com/vishnusomank/go-spiffe/v2/spiffeid"
	"github.com/vishnusomank/go-spiffe/v2/svid/jwtsvid"
	"github.com/vishnusomank/go-spiffe/v2/workloadapi"
)

func ExampleFetchX509SVID() {
	svid, err := workloadapi.FetchX509SVID(context.TODO(), map[string]string{})
	if err != nil {
		// TODO: error handling
	}

	// TODO: use the X509-SVID
	svid = svid
}

func ExampleFetchJWTSVID() {
	serverID, err := spiffeid.FromString("spiffe://example.org/server")
	if err != nil {
		// TODO: error handling
	}

	svid, err := workloadapi.FetchJWTSVID(context.TODO(), jwtsvid.Params{
		Audience: serverID.String(),
	}, map[string]string{})
	if err != nil {
		// TODO: error handling
	}

	// TODO: use the JWT-SVID
	svid = svid
}

func ExampleValidateJWTSVID() {
	serverID, err := spiffeid.FromString("spiffe://example.org/server")
	if err != nil {
		// TODO: error handling
	}

	token := "TODO"
	svid, err := workloadapi.ValidateJWTSVID(context.TODO(), token, serverID.String(), map[string]string{})
	if err != nil {
		// TODO: error handling
	}

	// TODO: use the JWT-SVID
	svid = svid
}
