package main
import (
	"fmt"
	"net/http"
)
func main() {
    http.HandleFunc("/hello", hellohandler) 

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Glitch !")
}
