package builder

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"runtime"
)

const (
	BUILD_STATUS_SUCC        = 0
	BUILD_STATUS_PENDING     = 1
	BUILD_STATUS_ERR_GENERAL = 100
)

var BuildClientConfig *dockerConfig

type dockerConfig struct {
	// dockerConfig implements the buildConfig interface.
	hostname string            // hostname is the Docker host URI.
	platform string            // The platform to tag completed images as.
	auth     *types.AuthConfig // The base64 authentication string to authenticate to registries as.
}

// Supplies the authentication string to the Docker API (or an error if it's nil).
// Not currently used, see authBase64
func (conf dockerConfig) privilegeFunc() (string, error) {
	return "", errors.New("not currently used")
}

// Supplies a base64 encoded authentication string based on the config's auth field.
func (conf dockerConfig) authBase64() (b64 string, err error) {
	authBytes, err := json.Marshal(conf.auth)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(authBytes), nil
}

func (conf dockerConfig) getClient() (*client.Client, error) {
	return client.NewClientWithOpts(
		client.WithHost(conf.hostname),
		client.WithVersion("v1.39"),
		client.WithHTTPHeaders(map[string]string{"User-Agent": "/r/openttd/tsubasa/service.build.docker"}))
}

func CreatedockerConfig(hostname string, auth_user string, auth_password string) (config *dockerConfig, err error) {
	config = &dockerConfig{hostname: hostname}

	// TODO is this the right way to format Docker platforms for multi-arch images?
	// in theory yes because docker itself uses golang? need to review the docker build pipeline code
	config.platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	// add the authentication details to the config
	config.auth = &types.AuthConfig{
		Username: auth_user,
		Password: auth_password,
	}

	return config, nil

}

type dockerBuildConfig struct {
	// dockerBuildConfig is a configuration struct for one docker build.

	repoName    string            // The Docker repository name.
	repoVersion string            // The unique version to tag this build as on the registry.
	buildSource string            // The source to build this image from (i.e a git URL).
	buildArgs   map[string]string // The build arguments to apply to the build.
	tags        []string          // The tags for this build.

	state int // The status of this build.
}

func CreatedockerBuildConfig(repoName string, repoVersion string, buildSource string, buildArgs map[string]string, tags []string) *dockerBuildConfig {
	return &dockerBuildConfig{
		repoName:    repoName,
		repoVersion: repoVersion,
		buildSource: buildSource,
		buildArgs:   buildArgs,
		tags:        tags,
		state:       BUILD_STATUS_PENDING,
	}
}

type dockerStreamMessage struct {
	// dockerStreamMessage is used to unmarshal JSON from a Docker io.Reader.
	Stream string `json:"stream"` // A Stream text entry.
}

type dockerImagePushMessage struct {
	// dockerImagePushMessage is used to unmarshal JSON from a Docker io.Reader provided by client.ImagePush().
	Status string `json:"status"` // The current status.
	Error  string `json:"error"`  // Set if an error is flagged.

}
