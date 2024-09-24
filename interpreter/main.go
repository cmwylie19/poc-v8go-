package main

import (
	"fmt"
	"log"
	"os"

	"rogchap.com/v8go"
)

func main() {
	// Ensure we have the correct number of arguments
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <controller-path> <module-path> <hash>", os.Args[0])
	}

	// Get the paths and hash from arguments
	controllerPath := os.Args[1]
	modulePath := os.Args[2]
	hash := os.Args[3]

	// Create a new V8 isolate (execution context)
	iso := v8go.NewIsolate()
	defer iso.Dispose()

	// Create a new V8 context for running JavaScript
	ctx := v8go.NewContext(iso)
	defer ctx.Close()

	// Define a Go function that handles log messages from JavaScript
	logFnTemplate := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		msg := info.Args()[0].String()
		fmt.Printf("[JS Log]: %s\n", msg)
		return nil
	})

	// Attach the log function to the global object
	global := ctx.Global()
	global.Set("log", logFnTemplate.GetFunction(ctx))

	// Inject the modulePath and hash as global variables
	_, err := ctx.RunScript(fmt.Sprintf(`const jsProgramPath = "%s";`, modulePath), "modulePath.js")
	if err != nil {
		log.Fatalf("Error setting jsProgramPath: %v", err)
	}

	_, err = ctx.RunScript(fmt.Sprintf(`const hash = "%s";`, hash), "hash.js")
	if err != nil {
		log.Fatalf("Error setting hash: %v", err)
	}

	// Load and execute the module.js file (this defines `hello`)
	moduleScript, err := os.ReadFile(modulePath)
	if err != nil {
		log.Fatalf("Failed to read module.js file %s: %v", modulePath, err)
	}

	_, err = ctx.RunScript(string(moduleScript), modulePath)
	if err != nil {
		log.Fatalf("Failed to execute module.js: %v", err)
	}

	// Read and execute the controller.js file
	controllerScript, err := os.ReadFile(controllerPath)
	if err != nil {
		log.Fatalf("Failed to read controller.js file %s: %v", controllerPath, err)
	}

	_, err = ctx.RunScript(string(controllerScript), controllerPath)
	if err != nil {
		log.Fatalf("Failed to execute controller.js: %v", err)
	}

	fmt.Println("Controller script executed successfully")
}
