package workloadapi

import (
	"context"

	"github.com/vishnusomank/go-spiffe/v2/bundle/jwtbundle"
	"github.com/vishnusomank/go-spiffe/v2/bundle/x509bundle"
	"github.com/vishnusomank/go-spiffe/v2/svid/jwtsvid"
	"github.com/vishnusomank/go-spiffe/v2/svid/x509svid"
)

// FetchX509SVID fetches the default X509-SVID, i.e. the first in the list
// returned by the Workload API.
func FetchX509SVID(ctx context.Context, meta map[string]string, options ...ClientOption) (*x509svid.SVID, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchX509SVID(ctx, meta)
}

// FetchX509SVIDs fetches all X509-SVIDs.
func FetchX509SVIDs(ctx context.Context, meta map[string]string, options ...ClientOption) ([]*x509svid.SVID, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchX509SVIDs(ctx, meta)
}

// FetchX509Bundle fetches the X.509 bundles.
func FetchX509Bundles(ctx context.Context, meta map[string]string, options ...ClientOption) (*x509bundle.Set, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchX509Bundles(ctx, meta)
}

// FetchX509Context fetches the X.509 context, which contains both X509-SVIDs
// and X.509 bundles.
func FetchX509Context(ctx context.Context, meta map[string]string, options ...ClientOption) (*X509Context, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchX509Context(ctx)
}

// WatchX509Context watches for updates to the X.509 context.
func WatchX509Context(ctx context.Context, watcher X509ContextWatcher, meta map[string]string, options ...ClientOption) error {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return err
	}
	defer c.Close()
	return c.WatchX509Context(ctx, watcher)
}

// FetchJWTSVID fetches a JWT-SVID.
func FetchJWTSVID(ctx context.Context, params jwtsvid.Params, meta map[string]string, options ...ClientOption) (*jwtsvid.SVID, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchJWTSVID(ctx, params)
}

// FetchJWTSVID fetches all JWT-SVIDs.
func FetchJWTSVIDs(ctx context.Context, params jwtsvid.Params, meta map[string]string, options ...ClientOption) ([]*jwtsvid.SVID, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchJWTSVIDs(ctx, params)
}

// FetchJWTBundles fetches the JWT bundles for JWT-SVID validation, keyed
// by a SPIFFE ID of the trust domain to which they belong.
func FetchJWTBundles(ctx context.Context, meta map[string]string, options ...ClientOption) (*jwtbundle.Set, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.FetchJWTBundles(ctx)
}

// WatchJWTBundles watches for changes to the JWT bundles.
func WatchJWTBundles(ctx context.Context, watcher JWTBundleWatcher, meta map[string]string, options ...ClientOption) error {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return err
	}
	defer c.Close()
	return c.WatchJWTBundles(ctx, watcher)
}

// WatchX509Bundles watches for changes to the X.509 bundles.
func WatchX509Bundles(ctx context.Context, watcher X509BundleWatcher, meta map[string]string, options ...ClientOption) error {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return err
	}
	defer c.Close()
	return c.WatchX509Bundles(ctx, watcher)
}

// ValidateJWTSVID validates the JWT-SVID token. The parsed and validated
// JWT-SVID is returned.
func ValidateJWTSVID(ctx context.Context, token, audience string, meta map[string]string, options ...ClientOption) (*jwtsvid.SVID, error) {
	c, err := New(ctx, meta, options...)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	return c.ValidateJWTSVID(ctx, token, audience)
}
