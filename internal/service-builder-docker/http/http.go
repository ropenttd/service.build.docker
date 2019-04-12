package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ropenttd/service.build.docker/pkg/builder"
	"log"
	"net/http"
)

func BuildHandler(w http.ResponseWriter, r *http.Request) *apiError {
	switch r.Method {
	case "POST":
		// Decode the JSON in the body
		d := json.NewDecoder(r.Body)
		p := &apiBuildRequest{}
		err := d.Decode(p)
		if err != nil {
			return &apiError{err, "JSON Decode Failed", http.StatusBadRequest}
		}
		if err = p.Validate(); err != nil {
			return &apiError{err, "Request Validation Failed", http.StatusBadRequest}
		}
		buildConfig := builder.CreatedockerBuildConfig(p.Target.RepositoryName, p.Target.ImageVersion, p.Source.Source, p.Source.BuildArguments, p.Target.Tags)
		// The HTTP request context is passed down so that if the request is cancelled, that is propagated to the builder.
		err = builder.Build(r.Context(), builder.BuildClientConfig, buildConfig)
		if err != nil {
			return &apiError{err, "Build Failed", http.StatusInternalServerError}
		}
		// Now push it.
		err = builder.Push(r.Context(), builder.BuildClientConfig, buildConfig)
		if err != nil {
			return &apiError{err, "Push Failed", http.StatusInternalServerError}
		}
		// Everything was awesome, return a response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		responseBody, err := json.Marshal(&apiGenericResponse{Status: 0})
		if err != nil {
			return &apiError{err, "Internal Error", http.StatusInternalServerError}
		}
		w.Write(responseBody)
		return nil
	default:
		return &apiError{
			errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed,
		}
	}
}

// Utility

type ApiHandler func(http.ResponseWriter, *http.Request) *apiError

// Handle the HTTP request.
// TODO add extra logging
func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		errorMessage, err := json.Marshal(&apiErrorJSONObject{Error: fmt.Sprintf("%s: %s", e.Message, e.Error), Status: e.Code})
		if err != nil {
			log.Printf("⁉️ Failed to marshal JSON error: %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(e.Code)
		w.Write(errorMessage)
	}
}
