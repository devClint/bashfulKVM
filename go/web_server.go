                                                                                // v0.008
package main
                                                                                //20 tabs | 80 spaces
import (
	"bytes"                                                                        // Importing the bytes package for capturing output
	"fmt"                                                                          // Importing the fmt package for formatting
	"log"                                                                          // Importing the log package for logging
	"net/http"                                                                     // Importing the net/http package for HTTP server
	"os/exec"                                                                      // Importing the os/exec package for executing shell commands
)

func startLab(w http.ResponseWriter, r *http.Request) {                         // Define a function to handle the '/start-lab' HTTP route
	var out bytes.Buffer                                                           // Initialize a buffer to capture command output
	cmd := exec.Command("/bin/bash", "bash/setup_lab.sh")                          // Create a new command
	cmd.Stdout = &out                                                              // Redirect the standard output to the buffer
	err := cmd.Run()                                                               // Run the command and wait for it to complete
	if err != nil {                                                                // Check if the command execution returned an error
		fmt.Println("Debug: Bash script failed to execute.")                          // Print a debug message
		w.Write([]byte("Lab setup failed."))                                          // Write a message to the HTTP response
		return                                                                        // Stop further execution
	}
	fmt.Printf("Debug: Bash script executed.\n%s", out.String())                   // Print the captured output from the shell script
	w.Write([]byte("Lab setup is in progress..."))                                 // Write a message to the HTTP response
}

func main() {                                                                   // Define the main function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {            // Register a function to handle the root URL path
		http.ServeFile(w, r, "gui/index.html")                                        // Serve the index.html file
	})
	http.HandleFunc("/start-lab", startLab)                                        // Register the startLab function to handle the '/start-lab' URL path

	fmt.Println("Debug: Web server started at http://localhost:4099/")             // Print a debug message indicating the server has started
	log.Fatal(http.ListenAndServe(":4099", nil))                                   // Start the HTTP server and listen on port 4099
}
