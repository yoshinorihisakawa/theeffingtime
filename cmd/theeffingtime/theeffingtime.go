package main

import (
	"time"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	version string
)

var responseTemplate = template.Must(template.New("main").Parse(`
<html>
<head>
</head>
<body>
The current date is: {{.Date}} </br>
UTC: {{.UTC}} </br>
Eastern: {{.Eastern}} </br>
Pacific: {{.Pacific}} </br>
</body>
</html>
`))

func generatedTemplate(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	data := map[string]interface{}{
		"Date": now.Date(),
		"UTC": now.UTC(),
		"Eastern": now.In(time.LoadLocation("America/New_York")),
		"Pacific": now.In(time.LoadLocation("America/Los_Angeles"))
	}
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
