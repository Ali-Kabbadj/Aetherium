//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("--> Generating sqlc code...")
	runCmd("sqlc", "-f", "configs/sqlc.yaml", "generate")

	fmt.Println("--> Generating gqlgen code...")
	runCmd("go", "run", "-v", "github.com/99designs/gqlgen", "generate", "--config", "configs/gqlgen.yml")



	fmt.Println("--> All code generated.")
}

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}