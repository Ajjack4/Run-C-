package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Compile the C++ program
	compileCmd := exec.Command("g++", "main.cpp", "-o", "main")
	compileCmd.Stderr = os.Stderr
	compileCmd.Stdout = os.Stdout
	err := compileCmd.Run() // Run the compilation process
	if err != nil {
		fmt.Println("Compilation Error:", err)
		return
	}
	fmt.Println("Compilation Successful!")

	// Run the compiled C++ executable
	runCmd := exec.Command("./main")
	runCmd.Stderr = os.Stderr
	runCmd.Stdout = os.Stdout
	runCmd.Stdin = os.Stdin // Ensure input is passed
	err = runCmd.Run()      // Run the executable
	if err != nil {
		fmt.Println("Execution Error:", err)
		return
	}

	// Check for input before exiting
	fmt.Println("Press Enter to exit...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Waits for user input before exiting
}
