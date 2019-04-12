package builder

import (
	"bufio"
	"encoding/json"
	"errors"
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

	// Prepare the build options.
	buildOptions := types.ImageBuildOptions{
		Dockerfile:    "Dockerfile",
		RemoteContext: buildConfig.buildSource,
		Tags:          targetTags,
		// NoCache: true,
		BuildArgs: buildConfig.buildArgs,
	}

	// Issue the build.
	log.Printf("👷 Dispatching build %s:%s", buildConfig.repoName, buildConfig.repoVersion)
	state, err := buildClient.client.ImageBuild(ctx, strings.NewReader(""), buildOptions)

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

	opt := types.ImagePushOptions{
		All:          true,
		RegistryAuth: buildClient.auth, // RegistryAuth is the base64 encoded credentials for the registry
		// PrivilegeFunc is not defined because we have exhausted our authentication options
		Platform: buildClient.platform,
	}
	var err error
	for _, tag := range buildConfig.tags {
		state, err := buildClient.client.ImagePush(ctx, buildConfig.repoName+":"+tag, opt)
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
