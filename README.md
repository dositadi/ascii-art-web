![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)

[![Go Report Card](https://goreportcard.com/badge/github.com/dositadi/ascii-art-web)](https://goreportcard.com/report/github.com/dositadi/ascii-art)

# ASCII-Art Web
A robust Go-based web application that transforms user-inputted text into stylized ASCII art using various banners. This project demonstrates a clear separation of concerns, utilizing Go interfaces for high modularity and testability.

## Features
Web-Based Interface: Easy-to-use GUI for inputting text and selecting banners.

* Multiple Banners: Supports standard, shadow, and thinkertoy styles.

* Modular Architecture: Uses a Processor pattern and interfaces to decouple the web server logic from the ASCII generation engine.

* Real-time Processing: Efficiently reads banner files line-by-line to construct characters dynamically.

## Project Structure
The codebase is organized into logical layers:

- `internal/handlers`: Manages HTTP requests/responses and template rendering.

- `internal/processor`: The core engine that handles file I/O and the algorithm for "stitching" ASCII characters together.

- `pkg/models`: Defines the data structures used across the application.

- `web/`: Contains HTML templates (`index.html`, `response.html`).

## Getting Started
### Prerequisites
Go 1.20 or higher installed.

### Installation
1. Clone the repository:

```bash
git clone https://github.com/dositadi/ascii-art-web.git
cd ascii-art-web
```
2. Ensure your banner files are located in the asset/ directory:

- `asset/standard.txt`

- `asset/shadow.txt`

- `asset/thinker_toy.txt`

### Running the Application
Execute the following command to start the server:

```bash
go run cmd/main.go
```

The server will start on `http://localhost:8081`.

## Technical Implementation
### 1. Character Mapping Logic
The engine uses a map-based system to find the starting line of each character within the banner files. Each character in the banner files typically occupies 8 lines plus a separator.

### 2. The Stitching Process
Unlike simple vertical printing, WriteArtOut uses a strings.Builder to concatenate character slices horizontally.

- Step A: Split input text by `\n`.

- Step B: Fetch the 8-line block for each character in a word.

- Step C: Iterate through lines 1-8, appending the corresponding line from every character in the word before moving to the next line.

## Technical Deep Dive
### The Interface Pattern
The project uses the AsciiInterface to define how text should be processed. This allows the web handlers to remain "agnostic" of the underlying processing logic.

```go
type AsciiInterface interface {
    Post(request m.Ascii) (string, *m.Error)
}
```

### The Processing Logic
When a request is received:

- Parsing: The input text is split by newline characters.

- Mapping: Each character is mapped to its specific line index in the selected banner file.

- Building: The DesignAllCharacters function reads the required 8 lines for each character and stores them in a 3D slice.

- Stitching: WriteArtOut iterates through the store to concatenate the strings horizontally and vertically to form the final art.

## Usage
1. Navigate to `http://localhost:8081` in your browser.

2. Type your desired text into the text area.

3. Select a banner style (Standard, Shadow, or Thinkertoy).

4. Click Submit to view your generated ASCII art.


## Contribution
This project was developed as part of a learning journey into Go's backend capabilities. Feel free to fork the repository and submit pull requests for new banner styles or optimization tweaks.

**Author**:    $Ositadinma Divine .A. (dositadi)$

## License

This is the [MIT](LICENSE) License