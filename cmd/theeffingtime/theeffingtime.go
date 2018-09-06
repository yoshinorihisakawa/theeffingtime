package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	version string
)

var responseTemplate = template.Must(template.New("main").Parse(`
<html>
<head>
</head>
<body>
UTC: {{.UTC}} </br>
Eastern: {{.Eastern}} </br>
Pacific: {{.Pacific}} </br>
</body>
</html>
`))

func generatedTemplate(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	utc := now.UTC()
	etz, eerr := time.LoadLocation("America/New_York")
	ptz, perr := time.LoadLocation("America/Los_Angeles")
	if eerr != nil || perr != nil {
		fmt.Println("Attempted and failed to load TZ Data:")
		fmt.Println(eerr)
		fmt.Println(perr)
	}
	data := map[string]interface{}{
		"UTC":     utc,
		"Eastern": now.In(etz),
		"Pacific": now.In(ptz)}
	outputBuffer := new(bytes.Buffer)
	responseTemplate.Execute(outputBuffer, data)
	fmt.Fprintln(w, outputBuffer)
}

func init() {
	version = "0.0.1"
}

func main() {
	http.HandleFunc("/", generatedTemplate)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
