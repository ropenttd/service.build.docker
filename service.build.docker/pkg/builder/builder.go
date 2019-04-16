package builder

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/ropenttd/tsubasa/generics/pkg/helpers"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func Build(ctx context.Context, buildClient *dockerConfig, buildConfig *dockerBuildConfig) error {
	// Populate the targetTags with Docker format repo/image:tag formatting from the buildTags.
	targetTags := []string{buildConfig.repoName + ":" + buildConfig.repoVersion}
	for _, tag := range buildConfig.tags {
		targetTags = append(targetTags, buildConfig.repoName+":"+tag)
	}

	// Get a new buildClient for this build.
	bc, err := buildClient.getClient()
	if err != nil {
		return err
	}

	// Prepare the build options.
	buildOptions := types.ImageBuildOptions{
		Dockerfile:    "Dockerfile",
		RemoteContext: buildConfig.buildSource,
		Tags:          targetTags,
		// NoCache: true,
		BuildArgs: *helpers.MarshalStringStringMapToPoint(&buildConfig.buildArgs),
	}

	// Issue the build.
	log.Printf("👷 Dispatching build %s:%s", buildConfig.repoName, buildConfig.repoVersion)
	state, err := bc.ImageBuild(ctx, strings.NewReader(""), buildOptions)

	// Buffer the output to a scanner.
	scanner := bufio.NewScanner(state.Body)

	for scanner.Scan() {
		var stream dockerStreamMessage
		json.Unmarshal([]byte(scanner.Text()), &stream)
		log.Printf("build debug: %s", stream.Stream)
		// fmt.Printf(scanner.Text())
	}
	if err != nil {
		buildConfig.state = BUILD_STATUS_ERR_GENERAL
		log.Printf("⚠️ Failed to build image: %s", err)
		return err
	} else {
		// Succ
		buildConfig.state = BUILD_STATUS_SUCC
		log.Printf("✅ Built image %s successfully.", buildConfig.repoName)
		return nil
	}
}

func Push(ctx context.Context, buildClient *dockerConfig, buildConfig *dockerBuildConfig) error {
	if buildConfig.state != BUILD_STATUS_SUCC {
		return errors.New("build has not or did not complete successfully, cannot push")
	}

	AuthString, err := buildClient.authBase64()

	if err != nil {
		return err
	}

	// Get a new buildClient for this build.
	bc, err := buildClient.getClient()
	if err != nil {
		return err
	}

	// Build the image push options.
	opt := types.ImagePushOptions{
		All:           true,
		RegistryAuth:  AuthString,                // RegistryAuth is the base64 encoded credentials for the registry
		PrivilegeFunc: buildClient.privilegeFunc, // Reserved for future use
		Platform:      buildClient.platform,
	}

	// This is done synchronously instead of async because the first push will contain all of our layers
	// subsequent pushes (which are just retags) will just reuse the layers already uploaded
	for _, tag := range buildConfig.tags {
		state, err := bc.ImagePush(ctx, buildConfig.repoName+":"+tag, opt)
		if err != nil {
			log.Printf("⚠️ Failed to push image: %s", err)
			return err
		}
		// Buffer the output to a scanner.
		defer state.Close()
		scanner := bufio.NewScanner(state)

		for scanner.Scan() {
			var stream dockerImagePushMessage
			json.Unmarshal([]byte(scanner.Text()), &stream)
			if stream.Status != "" {
				log.Printf("push debug: %s", stream.Status)
			}
			if stream.Error != "" {
				// An error occurred, break the loop
				err = errors.New(stream.Error)
				break
			}

		}
		if err != nil {
			log.Printf("⚠️ Failed to push image: %s", err)
			return err
		} else {
			log.Printf("✅ Pushed image %s successfully.", buildConfig.repoName+":"+tag)
		}
	}

	return err
}
