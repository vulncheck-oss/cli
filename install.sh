#!/bin/bash

# Fetch the latest version from GitHub
echo "Fetching latest version..."
VERSION=$(curl -s https://api.github.com/repos/vulncheck-oss/cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' | sed 's/v//')

if [ -z "$VERSION" ]; then
    echo "Failed to fetch the latest version."
    exit 1
fi

echo "Latest version: $VERSION"

# Detect the operating system and architecture
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macOS"
    INSTALL_DIR="/usr/local/bin"
    if [[ $(uname -m) == "arm64" ]]; then
        ARCH="arm64"
    else
        ARCH="amd64"
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="Linux"
    ARCH="x86_64"
    INSTALL_DIR="/usr/local/bin"
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    OS="Windows"
    ARCH="x86_64"
    INSTALL_DIR="/c/Windows/System32"
else
    echo "Unsupported operating system"
    exit 1
fi

# Set the appropriate filename based on the OS and architecture
if [[ "$OS" == "macOS" ]]; then
    FILENAME="vulncheck_${VERSION}_macOS_${ARCH}.zip"
elif [[ "$OS" == "Linux" ]]; then
    FILENAME="vulncheck_${VERSION}_Linux_${ARCH}.tar.gz"
elif [[ "$OS" == "Windows" ]]; then
    FILENAME="vulncheck_${VERSION}_Windows_${ARCH}.zip"
fi

# Download URL
URL="https://github.com/vulncheck-oss/cli/releases/download/v${VERSION}/${FILENAME}"

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Download the file
echo "Downloading $FILENAME..."
if [[ "$OS" == "Windows" ]]; then
    powershell -Command "Invoke-WebRequest -Uri $URL -OutFile $FILENAME"
else
    curl -LO "$URL"
fi

# Extract the contents
echo "Extracting..."
if [[ "$FILENAME" == *.zip ]]; then
    unzip "$FILENAME"
elif [[ "$FILENAME" == *.tar.gz ]]; then
    tar -xzf "$FILENAME"
fi

# Get the name of the extracted folder
EXTRACTED_FOLDER="${FILENAME%.*}"
if [[ "$FILENAME" == *.tar.gz ]]; then
    EXTRACTED_FOLDER="${EXTRACTED_FOLDER%.tar}"
fi

# Explain sudo usage
if [[ "$OS" != "Windows" ]]; then
    echo "To install vulncheck system-wide, we need to copy the binary to $INSTALL_DIR."
    echo "This requires administrator privileges."
    echo "You will be prompted for your password to perform this action."
    echo "This allows all users on this system to use the vulncheck command."
    echo "Press Ctrl+C to cancel the installation if you don't want to proceed."
fi


# Copy the binary to the install directory
echo "Installing vulncheck to $INSTALL_DIR..."
if [[ "$OS" == "Windows" ]]; then
    mv "$EXTRACTED_FOLDER/bin/vulncheck.exe" "$INSTALL_DIR"
else
    sudo mv "$EXTRACTED_FOLDER/bin/vulncheck" "$INSTALL_DIR"
fi

# Clean up
cd ..
rm -rf "$TEMP_DIR"

echo "Installation complete. You can now use 'vulncheck' command."

if [[ "$OS" != "Windows" ]]; then
    echo "The vulncheck binary has been installed to $INSTALL_DIR, which is already in your system's PATH."
    echo "You can run it by typing 'vulncheck' in your terminal."
else
    echo "The vulncheck binary has been installed to $INSTALL_DIR."
    echo "This directory should already be in your system's PATH."
    echo "You can run it by typing 'vulncheck' in your command prompt or PowerShell."
fi
