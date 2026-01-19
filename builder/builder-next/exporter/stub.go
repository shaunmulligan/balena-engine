//go:build no_buildkit

package exporter

import (
	"context"

	"github.com/distribution/reference"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const Moby = "moby"

// BuildkitCallbacks contains callbacks that buildkit uses to notify the daemon
// about events.
type BuildkitCallbacks struct {
	// Exported is called when an image is exported by buildkit.
	Exported func(ctx context.Context, id string, desc ocispec.Descriptor)

	// Named is a callback that is called when an image is created in the
	// containerd image store by buildkit.
	Named func(ctx context.Context, ref reference.NamedTagged, desc ocispec.Descriptor)
}
