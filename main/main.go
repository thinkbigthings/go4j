package main

// Import CloudBees platform SDK
import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)
import fmserver "github.com/rollout/rox-go/v5/server"

// Flags Create Rox flags in the Flags container class
type Flags struct {
	EnableTutorial fmserver.RoxFlag
}

var flags = &Flags{
	// Define the feature flags
	EnableTutorial: fmserver.NewRoxFlag(false),
}

var rox *fmserver.Rox

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	apiKey := os.Getenv("FM_DEV_SDK_KEY")
	if apiKey == "" {
		fmt.Println("API key is not set")
		return
	}

	options := fmserver.NewRoxOptions(fmserver.RoxOptionsBuilder{})

	rox := fmserver.NewRox()

	// Register the flags container with the CloudBees platform
	rox.RegisterWithEmptyNamespace(flags)

	// Setup the feature management environment key
	<-rox.Setup(apiKey, options)

	// Boolean flag example
	fmt.Println("EnableTutorial's value is ", flags.EnableTutorial.IsEnabled(nil))
}
