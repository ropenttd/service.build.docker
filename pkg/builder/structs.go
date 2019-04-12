package builder

import (
	"fmt"
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

	client   *client.Client
	hostname string // hostname is the Docker host URI.
	platform string // The platform to tag completed images as.
	auth     string // The base64 authentication string to authenticate to registries as.
}

func CreatedockerConfig(hostname string, auth string) (config *dockerConfig, err error) {
	config = &dockerConfig{hostname: hostname}

	config.client, err = client.NewClientWithOpts(
		client.WithHost(hostname),
		client.WithVersion("v1.39"),
		client.WithHTTPHeaders(map[string]string{"User-Agent": "tsubasa"}))
	if err != nil {
		return nil, err
	}

	// TODO is this the right way to format Docker platforms for multi-arch images?
	// in theory yes because docker itself uses golang? need to review the docker build pipeline code
	config.platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	// add the authentication details to the config
	config.auth = auth

	return config, nil

}

type dockerBuildConfig struct {
	// dockerBuildConfig is a configuration struct for one docker build.

	repoName    string             // The Docker repository name.
	repoVersion string             // The unique version to tag this build as on the registry.
	buildSource string             // The source to build this image from (i.e a git URL).
	buildArgs   map[string]*string // The build arguments to apply to the build.
	tags        []string           // The tags for this build.

	state int // The status of this build.
}

func CreatedockerBuildConfig(repoName string, repoVersion string, buildSource string, buildArgs map[string]*string, tags []string) *dockerBuildConfig {
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
