package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/start-lab", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Debug: Start Lab button clicked.")
		cmd := exec.Command("./setup_lab_with_progress_and_debugging.sh")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Debug: Bash script failed to execute.")
			fmt.Fprintf(w, "Lab setup failed.")
			return
		}
		fmt.Println("Debug: Bash script executed.")
		fmt.Fprintf(w, "Lab setup is in progress...")
	})

	fmt.Println("Debug: Web server started at http://localhost:4099/")
	http.ListenAndServe(":4099", nil)
}
