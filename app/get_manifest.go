package app

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Manifest struct {
	SchemaVersion int      `json:"schemaVersion"`
	MediaType     string   `json:"mediaType"`
	Config        *Layer   `json:"config"`
	Layers        []*Layer `json:"layers"`

	filepath string
	fi       os.FileInfo
	digest   string
}

type Layer struct {
	MediaType string `json:"mediaType"`
	Digest    string `json:"digest"`
	Size      int64  `json:"size"`
	From      string `json:"from,omitempty"`
	status    string
}

func GetManifest(ctx context.Context, modelPath ModelPath, regOpts *registryOptions) (*Manifest, string, error) {
	requestURL := modelPath.BaseURL().JoinPath("v2", modelPath.GetNamespaceRepository(), "manifests", modelPath.Tag)

	headers := make(http.Header)
	headers.Set("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	resp, err := makeRequest(ctx, http.MethodGet, requestURL, headers, nil, regOpts)
	if err != nil {
		return nil, "", err
	}

	if !strings.HasPrefix(resp.Status, "2") {
		body, _ := io.ReadAll(resp.Body)
		log.Println(string(body))
		log.Println(requestURL.String())
		log.Println(headers)
		return nil, "", errors.New("error getting manifest: " + resp.Status)
	}
	defer resp.Body.Close()

	var m *Manifest
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, "", err
	}

	return m, requestURL.String(), err
}
