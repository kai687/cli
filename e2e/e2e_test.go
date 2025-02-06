//go:build e2e

package e2e_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/algolia/cli/pkg/cmd/root"
	"github.com/cli/go-internal/testscript"
)

// algolia runs the root command of the Algolia CLI
func algolia() int {
	return int(root.Execute())
}

// TestMain sets the executable program so that we don't depend on the compiled binary
func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"algolia": algolia,
	}))
}

// testEnvironment stores the environment variables we need to setup for the tests
type testEnvironment struct {
	AppID  string
	ApiKey string
}

// getEnv reads the environment variables and prints errors for missing ones
func (e *testEnvironment) getEnv() error {
	env := map[string]string{}

	required := []string{
		// The CLI testing Algolia app
		"ALGOLIA_APPLICATION_ID",
		// API key with sufficient permissions to run all tests
		"ALGOLIA_API_KEY",
	}

	var missing []string

	for _, envVar := range required {
		val, ok := os.LookupEnv(envVar)
		if val == "" || !ok {
			missing = append(missing, envVar)
			continue
		}

		env[envVar] = val
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing environment variables: %s", strings.Join(missing, ", "))
	}

	e.AppID = env["ALGOLIA_APPLICATION_ID"]
	e.ApiKey = env["ALGOLIA_API_KEY"]

	return nil
}

// setupEnv sets up the environment variables for the test
func setupEnv(testEnv testEnvironment) func(ts *testscript.Env) error {
	return func(ts *testscript.Env) error {
		ts.Setenv("ALGOLIA_APPLICATION_ID", testEnv.AppID)
		ts.Setenv("ALGOLIA_API_KEY", testEnv.ApiKey)

		return nil
	}
}

// Run the test scripts from the directory `testscripts`
func TestCommands(t *testing.T) {
	var testEnv testEnvironment
	if err := testEnv.getEnv(); err != nil {
		t.Fatal(err)
	}

	t.Parallel()

	testscript.Run(t, testscript.Params{
		Dir:   "testscripts",
		Setup: setupEnv(testEnv),
	})
}
