# Time-To

A simple Go command-line tool for calculating time differences.

## Prerequisites

- Go 1.22.4 or later

## Installation

1. Clone the repository:
```bash
git clone https://github.com/malinatrash/time-to.git
cd time-to
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o time-to ./cmd
```

## Usage

Run the application with the following syntax:

```bash
./time-to -d <date> [-m <measure>]
```

### Parameters

- `-d <date>` (required): Target date in DD.MM.YYYY format
- `-m <measure>` (optional): Unit of measure for calculation (default: "d")
  - `d` - days
  - Other units may be supported depending on implementation

### Examples

Calculate days until December 25, 2024:
```bash
./time-to -d 25.12.2024
```

Calculate with specific measure:
```bash
./time-to -d 01.01.2025 -m d
```

## Development

To run the application directly without building:

```bash
go run ./cmd
```

## Project Structure

```
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   └── date/
│       ├── calculator/  # Date calculation logic
│       └── parser/      # Command-line argument parsing
└── pkg/
    └── checker/         # Error checking utilities
```
