
package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/start-lab", startLab)
	log.Fatal(http.ListenAndServe(":4099", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
		<body>
			<h1>Welcome to Data Center Lab</h1>
			<form action="/start-lab" method="post">
				<input type="submit" value="Start Lab">
			</form>
		</body>
	</html>`)
}

func startLab(w http.ResponseWriter, r *http.Request) {
	// Execute the Bash script to set up the lab
	cmd := exec.Command("/bin/bash", "setup_lab.sh")
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to start the lab:", err)
		fmt.Fprintf(w, "Failed to start the lab: %s", err)
		return
	}

	fmt.Fprintf(w, "Lab setup is in progress...")
}
