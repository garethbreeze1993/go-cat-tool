# go-cat-tool

## Overview
`go-cat-tool` is a command-line tool built in Go that is my version of the built in cat tool. This was done as part of this coding challenge Build Your Own cat Tool 
from https://codingchallenges.fyi/challenges/challenge-cat/

Follow the instructions below to clone the repository and install the executable on your local machine.

## Installation

To install and run `go-cat-tool` on your machine, follow these steps:

### 1. Clone the repository
First, clone the GitHub repository to your local machine using the following command:

```bash
git clone https://github.com/garethbreeze1993/go-cat-tool.git
```
### 2. Navigate to the project directory

Move into the cloned directory:
```bash
cd go-cat-tool
```

### 3. Install the executable

Run the following command to install the go-cat-tool executable:

```bash
go install
```

This will compile the source code and place the binary into your $GOPATH/bin directory, making the tool available globally from your terminal.
### 4. Run the tool

Once installed, you can run the tool from anywhere in your terminal:

```bash
go-cat-tool [options] [arguments]
````
### Options:

--number, -n: Whether user wants all lines numbered (default: false)

--nonblank, -b: Whether user wants only non-blank lines numbered (default: false)

**Note:** You can use either the `--number` or `--nonblank` flag, or omit both if you don't want any lines numbered.
### Arguments:

[file paths...]: One or more paths to files on the user's computer.

For example, to display all lines with line numbers from a file:

```bash
go-cat-tool -n /path/to/file
```
Or, to number only non-blank lines:
```bash
go-cat-tool -b /path/to/file
```
Finally to just display all lines without line numbers use this

```bash
go-cat-tool /path/to/file
```

### Prerequisites

Go must be installed on your machine. You can download it from the official Go website.
Ensure your $GOPATH/bin is added to your system's PATH.

### Additional Information

For more details on how to use go-cat-tool, check out the Usage section below or run:

```bash
go-cat-tool --help
```