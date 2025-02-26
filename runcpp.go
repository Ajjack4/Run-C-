package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	fmt.Println("Current Directory:", dir)

	// Compile the C++ file (use absolute path if needed)
	compileCmd := exec.Command("g++", dir+"/main.cpp", "-o", "main")
	compileCmd.Stderr = os.Stderr
	compileCmd.Stdout = os.Stdout
	err = compileCmd.Run()
	if err != nil {
		fmt.Println("Compilation Error:", err)
		return
	}
	fmt.Println("Compilation Successful!")

	// Run the compiled C++ executable
	runCmd := exec.Command("./main")
	runCmd.Stderr = os.Stderr
	runCmd.Stdout = os.Stdout
	err = runCmd.Run()
	if err != nil {
		fmt.Println("Execution Error:", err)
		return
	}
}
