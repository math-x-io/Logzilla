# Logzilla

![](img/logo.png)

Logzilla is an open-source log analysis tool designed for security teams and system administrators. It leverages a structured approach to analyze log data and generate detailed Markdown reports using the Ollama API for report generation.

## Purpose and Goals

- **Comprehensive Analysis:** Logzilla processes raw log files, extracts key events, and produces a structured report that outlines significant patterns and potential threats.
- **Security Insights:** The report includes sections such as Analysis Summary, Relevant Log Excerpts, Detailed Attack Analysis, Potential Impact, and Security Recommendations, making it easier for security professionals to understand the context and significance of log data.
- **Actionable Recommendations:** By highlighting vulnerabilities and suggesting security measures, Logzilla provides clear and focused guidance.
- **Ease of Use:** With a simple command-line interface, users can provide any log file and get a quick, professional analysis suitable for further investigation or audits.

## How It Works

1. **Log File Input:** Users pass the path of a log file via a command-line flag.

2. **Ollama API-based Analysis:** Logzilla reads the log file content and sends it to a locally running Ollama API. The API performs in-depth analysis and generates a detailed Markdown report.

3. **Markdown Report Generation:** The detailed report covers key aspects of potential security incidents.

4. **Output:** The tool outputs the report to stdout and saves it in an `output` directory for further review.

## Requirements and Setup

1. **Build the Application:**
    - Ensure Go is installed on your system.
    - Build the executable using `go build`.
    
2. **Ollama API Setup:**
    - Install and run the Ollama API server on your local machine.
    - Configure the API endpoint if necessary. Logzilla relies on the Ollama API to generate its Markdown reports.

3. **Usage:**
    - Run Logzilla with a specified log file:
      ```
      ./logzilla -log=/path/to/your/logfile
      ```

4. **Review the Report:**
    - The generated report will display on your terminal and be saved to the `output/report.md` file.

## Contributing

Contributions to improve the prompt and accuracy of log analysis are welcome. Feel free to submit pull requests or open issues on the project's repository.