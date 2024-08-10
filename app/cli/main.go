package cli

import (
	"context"
	"flag"
	"fmt"
	"github.com/amirrezaDev1378/app"
	"os"
	"slices"
	"strings"
)

const helpCommandText = `
Usage:
get <model-name> 
install --model=<model-name> --blobsPath=<blobs-path> 
help | -h display this help message
`

var availableCommands = []string{"get", "install", "help", "-h", "--help"}

func getSubCommand() (string, error) {
	if len(os.Args) < 2 {
		return "", fmt.Errorf("no command provided, available commands: %v", strings.Join(availableCommands, " ,"))
	}
	if !slices.Contains(availableCommands, os.Args[1]) {
		return "", fmt.Errorf("invalid command provided, available commands: %v", strings.Join(availableCommands, " ,"))
	}

	return os.Args[1], nil
}
func getCliArgs(count int) ([]string, error) {
	if len(os.Args) < 3 {
		return []string{}, fmt.Errorf("invalid usage of app, for help message use -h or --help")
	}
	var args []string
	for i := 0; i < count; i++ {
		args = append(args, os.Args[i+2])
	}
	return args, nil
}
func InitCli() {
	command, err := getSubCommand()
	if err != nil {
		os.Exit(1)
		return
	}
	if command == "-h" || command == "--help" {
		command = "help"
	}

	switch command {
	case "get":
		CommandGet()
		break
	case "install":
		CommandInstall()
		break
	case "help":
		fmt.Print(helpCommandText)
		os.Exit(0)
	default:
		panic("app panic : invalid command passed to cli")
	}
}

func CommandGet() {
	cliArgs, err := getCliArgs(1)
	modelName := cliArgs[0]
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("get direct download link for model :", modelName)

	parsedModelPath := app.ParseModelPath(modelName)

	manifest, manifestLink, err := app.GetManifest(context.TODO(), parsedModelPath, app.DefaultRegistryConfig)
	if err != nil {
		fmt.Println("error getting manifest: ", err)
		os.Exit(1)
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

	fmt.Printf("Manifest download link: %s\n", manifestLink)
	fmt.Println("Download links for layers:")
	for i, link := range downloadLinks {
		fmt.Printf("%d- %s\n", i+1, link)
	}
	fmt.Println("Generated download links for model : ", modelName, "finished successfully.")
	os.Exit(0)
}

func CommandInstall() {
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)
	modelName := installCmd.String("model", "", "model name to install")
	blobsPath := installCmd.String("blobsPath", "", "downloaded blobs path")
	installCmd.Parse(os.Args[2:])

	if *modelName == "" || *blobsPath == "" {
		fmt.Println("model name and blobs path are required")
		os.Exit(1)
	}

	hasPermission, err := app.HasElevatedPermissions()
	if err != nil {
		fmt.Println("error checking permissions: ", err)
		os.Exit(1)
	}
	if !hasPermission {
		fmt.Println("please run the command with elevated permissions")
		os.Exit(1)
	}

	fmt.Println("installing model :", *modelName)

	err = app.VerifyDownloadedModel(*modelName, *blobsPath)
	if err != nil {
		fmt.Println("error verifying model: ", err)
		os.Exit(1)
	}

	err = app.InstallModel(*modelName, *blobsPath)
	if err != nil {
		fmt.Println("error installing model: ", err)
		os.Exit(1)
	}

}
