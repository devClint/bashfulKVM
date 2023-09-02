
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func startLab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Start Lab button clicked.")
	cmd := exec.Command("/bin/bash", "-c", "./test_script.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Command output: %s\n", out)
	fmt.Fprintf(w, "Lab setup is in progress...")
}

func main() {
	http.HandleFunc("/start", startLab)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Debug: Web server started at http://localhost:4099/")
	log.Fatal(http.ListenAndServe(":4099", nil))
}
