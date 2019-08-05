package grpc

import (
	"context"
	"errors"
	pbuf_generics "github.com/ropenttd/tsubasa/generics/api/protobuf"
	pbuf_service "github.com/ropenttd/tsubasa/service.build.docker/api/protobuf"
	"github.com/ropenttd/tsubasa/service.build.docker/pkg/builder"
)

type BuildDockerHandler struct {
}

func (s BuildDockerHandler) Build(ctx context.Context, r *pbuf_service.BuildRequest) (response *pbuf_generics.StatusResponse, err error) {
	// Do request validation
	if r.Target.RepoName == "" {
		return nil, errors.New("Target.RepoName cannot be nil")
	} else if r.Target.ImageVersion == "" {
		return nil, errors.New("Target.ImageVersion cannot be nil")
	} else if r.Source.Repo == "" {
		return nil, errors.New("Source.Repo cannot be nil")
	}

	buildConfig := builder.CreatedockerBuildConfig(r.Target.RepoName, r.Target.ImageVersion, r.Source.Repo, r.Source.GetBuildArgs(), r.Target.GetTag())
	// The HTTP request context is passed down so that if the request is cancelled, that is propagated to the builder.
	err = builder.Build(ctx, builder.BuildClientConfig, buildConfig)
	if err != nil {
		return nil, err
	}
	// Now push it.
	err = builder.Push(ctx, builder.BuildClientConfig, buildConfig)
	if err != nil {
		return nil, err
	}
	// Everything was awesome, return a response
	return &pbuf_generics.StatusResponse{Status: pbuf_generics.StatusCode_SUCC}, nil
}
