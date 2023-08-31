
package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/start-lab", StartLab)

	fmt.Println("Server starting on port 4099")
	http.ListenAndServe(":4099", nil)
}

func StartLab(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("bash", "/mnt/data/bashfulKVM-v04/setup_lab_with_progress.sh")
	cmd.Start()

	fmt.Fprintf(w, "Lab setup is in progress...")
}
