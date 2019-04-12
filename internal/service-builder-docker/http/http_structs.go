package http

import "errors"

type apiBuildRequest struct {
	Source apiBuildRequestSourceData `json:"source"`
	Target apiBuildRequestTargetData `json:"target"`
}

type apiBuildRequestSourceData struct {
	Source         string             `json:"repo"`       // MANDATORY The source is the location of the Dockerfile to build (i.e a Git repository)
	BuildArguments map[string]*string `json:"build_args"` // Arguments to pass to the build environment.
}

type apiBuildRequestTargetData struct {
	RepositoryName string   `json:"repo_name"`     // MANDATORY The repositoryName is the name of the Docker Registry image to upload to.
	ImageVersion   string   `json:"image_version"` // MANDATORY The imageVersion is the version that this package should be tagged with.
	Tags           []string `json:"tags"`          // tags is an array of extra tags to tag this build as.
}

// Validator for above request
// surely there is a better way of doing this?????
func (r *apiBuildRequest) Validate() error {
	switch {
	case r.Source.Source == "":
		return errors.New("element repo in element source cannot be null")
	case r.Target.RepositoryName == "":
		return errors.New("element repo_name in element target cannot be null")
	case r.Target.ImageVersion == "":
		return errors.New("element image_version in element target cannot be null")
	default:
		// Nothing wrong here officer
		return nil
	}
}

type apiGenericResponse struct {
	// apiGenericResponse is a generic status response.
	Status int `json:"status"`
}

type apiError struct {
	Error   error
	Message string
	Code    int // The HTTP status code to return.
}

type apiErrorJSONObject struct {
	Status int    `json:"status"`
	Error  string `json:"message"`
}
