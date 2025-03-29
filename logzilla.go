package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	logPath := flag.String("log", "", "Path to the log file")
	flag.Parse()

	if *logPath == "" {
		fmt.Println("Usage: logzilla -log=/path/to/logfile")
		os.Exit(1)
	}

	// Read the log file content
	content, err := os.ReadFile(*logPath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	// Construct the prompt for the API call
	prompt := `
You're a cybersecurity expert and a seasoned system administrator. Analyze the following log data and generate a comprehensive and structured report in Markdown format. The report must include the following sections:

1: Analysis Summary: A synthesis of the most significant events observed in the logs.

2. Relevant Log Excerpts: Include specific portions of the logs that justify your conclusions, accompanied by clear explanations.

3. Detailed Attack Analysis:

- Identify and describe the type of attack (e.g., SQL injection, brute force, DDoS, etc.).

- Specify any vulnerabilities or CVEs that were exploited, if they can be inferred.

- Provide links to reports or online resources that describe this type of attack or vulnerabilities in detail.

4. Potential Impact: Describe the risks that this attack might pose to the system, data, or users.

5. Security Recommendations: List clear, specific, and actionable measures to prevent this type of attack in the future (e.g., software updates, configuration changes, enhanced monitoring, etc.).

Generate all your responses in English. Be technical and precise in your explanations, and ensure the report is professional and directly usable by a security team or system administrator.

The log data is as follows:
` + string(content)

	// Prepare the JSON payload.
	payload := map[string]interface{}{
		"prompt": prompt,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling JSON payload: %v", err)
	}

	// Call the local Ollama API (update the API URL if necessary)
	apiURL := "http://localhost:11434/api"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending HTTP request to the API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Fatalf("API responded with status %v: %s", resp.Status, string(body))
	}

	// Assume the API returns a JSON with a "report" field containing the markdown report.
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		log.Fatalf("Error decoding API response: %v", err)
	}

	report, ok := responseData["report"].(string)
	if !ok {
		log.Fatalf("Unexpected response format: 'report' field missing")
	}

	// Print the generated report to stdout.
	fmt.Println(report)

	// Create the output directory if it does not exist.
	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatalf("Unable to create output directory: %v", err)
	}

	// Save the report in the output directory.
	outputFilePath := "output/report.md"
	if err := os.WriteFile(outputFilePath, []byte(report), 0644); err != nil {
		log.Fatalf("Error writing report to file: %v", err)
	}

	fmt.Println("Report saved to", outputFilePath)
}
