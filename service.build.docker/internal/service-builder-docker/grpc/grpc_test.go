package grpc

import (
	"context"
	api_v1 "github.com/ropenttd/tsubasa/service.build.docker/api/protobuf"
	"github.com/ropenttd/tsubasa/service.build.docker/pkg/builder"
	"testing"
)

func initTestEnvironment(auth_user string, auth_pass string) error {
	var err error
	builder.BuildClientConfig, err = builder.CreatedockerConfig("unix:///var/run/docker.sock", auth_user, auth_pass)
	return err
}

func TestBuildRequestBadAuth(t *testing.T) {
	s := BuildDockerHandler{}
	// No auth credentials for this test.
	err := initTestEnvironment("", "")
	if err != nil {
		t.Error(err)
	}

	requestBody := api_v1.BuildRequest{
		Source: &api_v1.BuildRequest_BuildRequestSourceData{
			Repo: "https://github.com/ropenttd/docker_openttd.git",
			BuildArgs: map[string]string{
				"OPENTTD_VERSION": "1.9.1",
			},
		},
		// We don't have the necessary credentials to push to this target.
		Target: &api_v1.BuildRequest_BuildRequestTargetData{
			RepoName:     "redditopenttd/openttd",
			ImageVersion: "ohnothetestfailed",
			Tag: []string{
				"unit_test",
			},
		},
	}

	response, err := s.Build(context.Background(), &requestBody)
	if response != nil {
		t.Errorf("Got a response when we did not expect one: %s", response)
	}

	// Check the status code is what we expect.
	if err.Error() != "denied: requested access to the resource is denied" {
		t.Errorf("handler returned wrong error: got %v want %v",
			err.Error(), "denied: requested access to the resource is denied")
	}
}
