package main

import (
	"fmt"
	"log"
	"os"

	"rogchap.com/v8go"
)

func main() {

	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <controller-path> <module-path> <hash>", os.Args[0])
	}

	controllerPath := os.Args[1]
	modulePath := os.Args[2]
	hash := os.Args[3]

	iso := v8go.NewIsolate()
	defer iso.Dispose()

	ctx := v8go.NewContext(iso)
	defer ctx.Close()

	logFnTemplate := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		msg := info.Args()[0].String()
		fmt.Printf("[JS Log]: %s\n", msg)
		return nil
	})

	global := ctx.Global()
	global.Set("log", logFnTemplate.GetFunction(ctx))

	_, err := ctx.RunScript(fmt.Sprintf(`const jsProgramPath = "%s";`, modulePath), "modulePath.js")
	if err != nil {
		log.Fatalf("Error setting jsProgramPath: %v", err)
	}

	_, err = ctx.RunScript(fmt.Sprintf(`const hash = "%s";`, hash), "hash.js")
	if err != nil {
		log.Fatalf("Error setting hash: %v", err)
	}

	moduleScript, err := os.ReadFile(modulePath)
	if err != nil {
		log.Fatalf("Failed to read module.js file %s: %v", modulePath, err)
	}

	_, err = ctx.RunScript(string(moduleScript), modulePath)
	if err != nil {
		log.Fatalf("Failed to execute module.js: %v", err)
	}

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
