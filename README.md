# dupefinder

A fast and simple command-line tool to scan directories and find **duplicate files** based on content, not names.

## Features

- Scans directories recursively
- Detects duplicate files using SHA256 hashing
- Optionally include/exclude file extensions
- Export results to JSON with `--output`
- Colourful terminal output for easier scanning

##  Installation

As a prerequisite, you should have the [latest Go version](https://go.dev/dl/) installed.
```bash
git clone https://github.com/anthophobiac/dupefinder.git
cd dupefinder
go install
```

## Usage

```bash
dupefinder scan [path] [flags]
```

### Flags:

| Flag            | Description                                               |
|-----------------|-----------------------------------------------------------|
| `--include-ext` | Only scan files with these extensions (e.g. `.jpg,.txt`)  |
| `--exclude-ext` | Exclude files with these extensions                       |
| `--output, -o`  | Write duplicate results to a JSON file                    |

## Usage Examples

### Scan a directory and show duplicate groups:

```bash
dupefinder scan ./Downloads
```

### Only include `.mp3` and `.flac` files:

```bash
dupefinder scan ./Music --include-ext=.mp3,.flac
```

### Exclude `.log` and `.tmp` files:

```bash
dupefinder scan ./Logs --exclude-ext=.log,.tmp
```

### Save results to a JSON file (no console output):

```bash
dupefinder scan ./Documents --output=dupes.json
```

## Sample Output

```bash
Hashing files  [======================>] 100%
Duplicate group (e3b0c442...):
   ./photos/image1.jpg
   ./backup/image1_copy.jpg

Duplicate group (9a0364b9...):
   ./docs/report.pdf
   ./archive/old_report.pdf
```
