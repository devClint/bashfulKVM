
package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "gui/index.html")
	})

	http.HandleFunc("/start-lab", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Debug: Start Lab button clicked.")
		cmd := exec.Command("bash", "bash/setup_lab_with_progress_and_debugging_and_network_fix.sh")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Debug: Bash script failed to execute.")
			w.Write([]byte("Lab setup failed."))
		} else {
			fmt.Println("Debug: Bash script executed.")
			w.Write([]byte("Lab setup is in progress..."))
		}
	})

	fmt.Println("Debug: Web server started at http://localhost:4099/")
	log.Fatal(http.ListenAndServe(":4099", nil))
}
