package connections

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExecuteVQLQuery(query string) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx,"velociraptor","--api_config", "config/api.config.yaml", "query", query, "--format", "jsonl","|", "jq")

	// Run the Python script
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Handle timeout or other errors
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Python script timed out")
		} else {
			fmt.Println("Error executing Python script:", err)
		}
		return
	}

	// Print the output
	fmt.Println("Output from Python script:", string(output))
}