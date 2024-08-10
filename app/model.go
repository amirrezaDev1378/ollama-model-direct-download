package app

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	DefaultRegistry       = "registry.ollama.ai"
	DefaultNamespace      = "library"
	DefaultTag            = "latest"
	DefaultProtocolScheme = "https"
)

type ModelPath struct {
	ProtocolScheme string
	Registry       string
	Namespace      string
	Repository     string
	Tag            string
}

func ParseModelPath(name string) ModelPath {
	mp := ModelPath{
		ProtocolScheme: DefaultProtocolScheme,
		Registry:       DefaultRegistry,
		Namespace:      DefaultNamespace,
		Repository:     "",
		Tag:            DefaultTag,
	}

	before, after, found := strings.Cut(name, "://")
	if found {
		mp.ProtocolScheme = before
		name = after
	}

	name = strings.ReplaceAll(name, string(os.PathSeparator), "/")
	parts := strings.Split(name, "/")
	switch len(parts) {
	case 3:
		mp.Registry = parts[0]
		mp.Namespace = parts[1]
		mp.Repository = parts[2]
	case 2:
		mp.Namespace = parts[0]
		mp.Repository = parts[1]
	case 1:
		mp.Repository = parts[0]
	}

	if repo, tag, found := strings.Cut(mp.Repository, ":"); found {
		mp.Repository = repo
		mp.Tag = tag
	}

	return mp
}

func (mp ModelPath) GetNamespaceRepository() string {
	return fmt.Sprintf("%s/%s", mp.Namespace, mp.Repository)
}

func (mp ModelPath) BaseURL() *url.URL {
	return &url.URL{
		Scheme: mp.ProtocolScheme,
		Host:   mp.Registry,
	}
}
