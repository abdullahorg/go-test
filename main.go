package main;
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(CreateGreeting(name)))
}

func CreateGreeting(name string) string {
	if name == "" {
		name = "Guest"
	}
	return "Hello changed test, " + name + "\n"
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        hello := "Hello from Amnic!!"
        golang := "Golang Version :"
        golangversion := runtime.Version()
        config_key := "TEST_INT_VALUE"
        secret_key := "SAMPLE_SECRET"
    

        res := fmt.Sprintf("%s \n %s %s \n ", hello, golang, golangversion)
        fmt.Fprintf(w, res)


        valuesecret, ok := os.LookupEnv(secret_key)
	    if ok {
            // secrets := os.Getenv("SAMPLE_SECRET")
            secrets := valuesecret
            fmt.Fprintf(w, "secrets: %s\n",secrets)
	    } else {
            fmt.Fprintf(w, "Secrets is not present \n")
	    }

        valueconfig, okk := os.LookupEnv(config_key)
	    if okk {
            // config := os.Getenv("TEST_INT_VALUE")
            config := valueconfig
            fmt.Fprintf(w, "Configs: %s\n",config)
	    } else {
            fmt.Fprintf(w, "Config is not present")
	    }
        
    })
    fmt.Printf("Server running (port=8080), route: http://localhost:8080/")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
