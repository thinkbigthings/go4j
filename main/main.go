package main

// Import CloudBees platform SDK
import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rollout/rox-go/v5/core/context"
	"os"
)
import fmserver "github.com/rollout/rox-go/v5/server"

// Flags Create Rox flags in the Flags container class
type Flags struct {
	jyEnableTutorial fmserver.RoxFlag
	jyUseDebugMode   fmserver.RoxFlag
}

var flags = &Flags{
	// Define the feature flags
	jyEnableTutorial: fmserver.NewRoxFlag(false),
	jyUseDebugMode:   fmserver.NewRoxFlag(false),
}

var rox *fmserver.Rox

func main() {
}

func SetupFM() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiKey := os.Getenv("FM_DEV_SDK_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	} else {
		fmt.Println("API key is set: " + apiKey)
	}

	options := fmserver.NewRoxOptions(fmserver.RoxOptionsBuilder{})

	rox = fmserver.NewRox()

	// Register the flags container with the CloudBees platform
	rox.RegisterWithEmptyNamespace(flags)

	// Set up the feature management environment key
	<-rox.Setup(apiKey, options)
}

// GetTutorialMessage Pass in nil to run without context
func GetTutorialMessage(ctx context.Context) string {

	if flags.jyEnableTutorial.IsEnabled(ctx) {
		return "Tutorial is enabled"
	} else {
		return "Tutorial is disabled"
	}

}
