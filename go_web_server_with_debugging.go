
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func startLab(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Lab setup is in progress...")
	fmt.Println("Debug: Start Lab button clicked.")
	cmd := exec.Command("/bin/bash", "-c", "./setup_lab_with_debugging.sh")
	cmd.Start()
	fmt.Println("Debug: Bash script executed.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/start-lab", startLab)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Debug: Web server started at http://localhost:4099/")
	log.Fatal(http.ListenAndServe(":4099", nil))
}
