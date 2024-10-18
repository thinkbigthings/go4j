package main

import (
	"github.com/rollout/rox-go/v5/core/context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFM(t *testing.T) {

	SetupFM()

	// use a context to get the feature flags
	// this could be info that comes from a request like a user id or user role

	// dev environment flag should have jyEnableTutorial set to true if username is Jason
	// so this test should return "Tutorial is enabled" for username Jason,
	// and "Tutorial is disabled" for any other username or if the context is nil
	contextProps := make(map[string]interface{})
	contextProps["username"] = "Jason"
	contextProps["isAdmin"] = false
	currentContext := context.NewContext(contextProps)

	// TODO look for context.Background() check golang context vs fm V5 context
	message := GetTutorialMessage(currentContext)

	assert.Equal(t, "Tutorial is disabled", message)

}
