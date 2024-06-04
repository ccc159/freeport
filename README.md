# Freeport

A simple CLI tool to kill the processes that are using a specific port.

## Installation

### Mac - Homebrew

`brew install freeport`

### Linux - apt

`sudo apt install freeport`

### Windows - winget

`winget install freeport`

### Build locally

If you want to build the project locally, you can run the following commands:

```bash
git clone https://github.com/ccc159/freeport.git
cd freeport
chmod +x build_local.sh
./build_local.sh
```

Then you'll find the binary in the `build` folder.

## Usage

To kill the process that is using a specific port, run:
`freeport <port>`

Check the version:
`freeport --version`
