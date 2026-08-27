package main

import (
	"context"
	"fmt"

	"github.com/amirrezaDev1378/ollama-model-direct-download/app"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type GetLinkResponse struct {
	ManifestLink  string   `json:"manifestLink"`
	DownloadLinks []string `json:"downloadLinks"`
}

// GetModelDownloadLinks returns download links for a given model
func (a *App) GetModelDownloadLinks(modelName string) (*GetLinkResponse, error) {
	parsedModelPath := app.ParseModelPath(modelName)

	manifest, manifestLink, err := app.GetManifest(a.ctx, parsedModelPath, app.DefaultRegistryConfig)
	if err != nil {
		return nil, err
	}

	var layers []*app.Layer
	layers = append(layers, manifest.Layers...)
	layers = append(layers, manifest.Config)

	var downloadLinks []string

	for _, layer := range layers {
		config := app.DownloadLinkConfig{
			ModelPath: parsedModelPath,
			Digest:    layer.Digest,
			RegOpts:   app.DefaultRegistryConfig,
		}
		link := config.GetDownloadLink()
		downloadLinks = append(downloadLinks, link)
	}

	return &GetLinkResponse{
		ManifestLink:  manifestLink,
		DownloadLinks: downloadLinks,
	}, nil
}

// InstallModel installs a downloaded model
func (a *App) InstallModel(modelName string, blobsPath string) error {
	hasPermission, err := app.HasElevatedPermissions()
	if err != nil {
		return err
	}
	if !hasPermission {
		return fmt.Errorf("no elevated permissions, please run as administrator/root")
	}

	err = app.VerifyDownloadedModel(modelName, blobsPath)
	if err != nil {
		return err
	}

	err = app.InstallModel(modelName, blobsPath, true, func(progress float64, message string) {
		runtime.EventsEmit(a.ctx, "install_progress", map[string]interface{}{
			"progress": progress,
			"message":  message,
		})
	})
	if err != nil {
		return err
	}

	return nil
}
