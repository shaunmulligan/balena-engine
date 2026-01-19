//go:build no_buildkit

package buildkit

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/backend"
	"github.com/docker/docker/builder"
	"github.com/docker/docker/builder/builder-next/exporter"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/daemon/images"
	"github.com/docker/docker/libnetwork"
	"github.com/docker/docker/pkg/idtools"

	"github.com/containerd/containerd/remotes/docker"
	"github.com/moby/buildkit/session"
	"google.golang.org/grpc"
)

// ErrBuildKitDisabled is returned when BuildKit operations are attempted but BuildKit is disabled
var ErrBuildKitDisabled = errors.New("buildkit is disabled in this build")

// Opt is option struct required for creating the builder
type Opt struct {
	SessionManager      *session.Manager
	Root                string
	EngineID            string
	Dist                images.DistributionServices
	ImageTagger         interface{}
	NetworkController   *libnetwork.Controller
	DefaultCgroupParent string
	RegistryHosts       docker.RegistryHosts
	BuilderConfig       config.BuilderConfig
	Rootless            bool
	IdentityMapping     idtools.IdentityMapping
	DNSConfig           config.DNSConfig
	ApparmorProfile     string
	UseSnapshotter      bool
	Snapshotter         string
	ContainerdAddress   string
	ContainerdNamespace string
	Callbacks           exporter.BuildkitCallbacks
}

// Builder is a stub when BuildKit is disabled
type Builder struct{}

// New returns nil when BuildKit is disabled
func New(ctx context.Context, opt Opt) (*Builder, error) {
	return nil, nil
}

// Build returns an error when BuildKit is disabled
func (b *Builder) Build(ctx context.Context, opt backend.BuildConfig) (*builder.Result, error) {
	return nil, ErrBuildKitDisabled
}

// Prune returns an error when BuildKit is disabled
func (b *Builder) Prune(ctx context.Context, opts types.BuildCachePruneOptions) (int64, []string, error) {
	return 0, nil, ErrBuildKitDisabled
}

// DiskUsage returns an error when BuildKit is disabled
func (b *Builder) DiskUsage(ctx context.Context) ([]*types.BuildCache, error) {
	return nil, ErrBuildKitDisabled
}

// Cancel is a no-op when BuildKit is disabled
func (b *Builder) Cancel(ctx context.Context, id string) error {
	return nil
}

// RegisterGRPC is a no-op when BuildKit is disabled
func (b *Builder) RegisterGRPC(s *grpc.Server) {
}

// Close is a no-op when BuildKit is disabled
func (b *Builder) Close() error {
	return nil
}
