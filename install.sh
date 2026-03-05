#!/bin/bash

INSTALL_MODE=""

# Parse command line arguments
for arg in "$@"
do
    case $arg in
        --version=*)
        SPECIFIED_VERSION="${arg#*=}"
        shift # Remove --version= from processing
        ;;
        --sudo)
        INSTALL_MODE="sudo"
        shift
        ;;
        --non-sudo)
        INSTALL_MODE="non-sudo"
        shift
        ;;
        --help|-h)
        echo "Usage: $0 [OPTIONS]"
        echo ""
        echo "Options:"
        echo "  --sudo          Install system-wide (requires sudo)"
        echo "  --non-sudo      Install for current user only (no sudo required)"
        echo "  --version=X.Y.Z Install specific version instead of latest"
        echo "  --help, -h      Show this help message"
        exit 0
        ;;
    esac
done

if [ -n "$SPECIFIED_VERSION" ]; then
    VERSION="$SPECIFIED_VERSION"
    echo "Using specified version: $VERSION"
else
    # Fetch the latest version from GitHub
    echo "Fetching latest version..."
    VERSION=$(curl -s https://api.github.com/repos/vulncheck-oss/cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' | sed 's/v//')

    if [ -z "$VERSION" ]; then
        echo "Failed to fetch the latest version. Please check your internet connection and try again."
        exit 1
    fi

    echo "Latest version: $VERSION"
fi

# Detect the operating system and architecture
if [[ "$OSTYPE" == "darwin"* ]]; then
    OS="macOS"
    DEFAULT_INSTALL_DIR="/usr/local/bin"
    LOCAL_INSTALL_DIR="$HOME/.local/bin"
    if [[ $(uname -m) == "arm64" ]]; then
        ARCH="arm64"
    else
        ARCH="amd64"
    fi
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    OS="Linux"
    ARCH="amd64"
    DEFAULT_INSTALL_DIR="/usr/local/bin"
    LOCAL_INSTALL_DIR="$HOME/.local/bin"
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    OS="Windows"
    ARCH="amd64"
    DEFAULT_INSTALL_DIR="/c/Windows/System32"
    LOCAL_INSTALL_DIR="$USERPROFILE/bin"
else
    echo "Unsupported operating system"
    exit 1
fi

# Ask user for installation preference (skip for Windows as it doesn't use sudo)
if [[ "$OS" != "Windows" ]]; then
    if [[ -z "$INSTALL_MODE" ]]; then
        echo ""
        echo "Installation Options:"
        echo "1) System-wide installation (requires sudo) - installs to $DEFAULT_INSTALL_DIR"
        echo "2) Local user installation (no sudo required) - installs to $LOCAL_INSTALL_DIR"
        echo ""
        echo "To skip this prompt in the future, use --sudo or --non-sudo flags"
        echo ""
        
        # Read from /dev/tty to work even when script is piped
        if [[ -t 0 ]] || [[ -e /dev/tty ]]; then
            echo -n "Please select an option (1 or 2): "
            read install_choice < /dev/tty
        else
            # Fallback if /dev/tty is not available
            echo "Unable to read user input. Please use --sudo or --non-sudo flag."
            echo "Example: curl -sSL ... | bash -s -- --non-sudo"
            exit 1
        fi
        
        case $install_choice in
            1)
                INSTALL_DIR="$DEFAULT_INSTALL_DIR"
                SYSTEM_WIDE=true
                echo "Selected: System-wide installation"
                ;;
            2)
                INSTALL_DIR="$LOCAL_INSTALL_DIR"
                SYSTEM_WIDE=false
                echo "Selected: Local user installation"
                ;;
            *)
                echo "Invalid choice. Please run the script again and select 1 or 2."
                exit 1
                ;;
        esac
    else
        case $INSTALL_MODE in
            "sudo")
                INSTALL_DIR="$DEFAULT_INSTALL_DIR"
                SYSTEM_WIDE=true
                echo "Selected: System-wide installation (via --sudo flag)"
                ;;
            "non-sudo")
                INSTALL_DIR="$LOCAL_INSTALL_DIR"
                SYSTEM_WIDE=false
                echo "Selected: Local user installation (via --non-sudo flag)"
                ;;
        esac
    fi
else
    # Windows doesn't need this choice
    INSTALL_DIR="$DEFAULT_INSTALL_DIR"
    SYSTEM_WIDE=true
fi

# Set the appropriate filename based on the OS and architecture
if [[ "$OS" == "macOS" ]]; then
    FILENAME="vulncheck_${VERSION}_macOS_${ARCH}.zip"
elif [[ "$OS" == "Linux" ]]; then
    FILENAME="vulncheck_${VERSION}_linux_${ARCH}.tar.gz"
elif [[ "$OS" == "Windows" ]]; then
    FILENAME="vulncheck_${VERSION}_windows_${ARCH}.zip"
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

# Explain sudo usage for system-wide installation
if [[ "$OS" != "Windows" ]] && [[ "$SYSTEM_WIDE" == true ]]; then
    echo ""
    echo "To install vulncheck system-wide, we need to copy the binary to $INSTALL_DIR."
    echo "This requires administrator privileges."
    echo "You will be prompted for your password to perform this action."
    echo "This allows all users on this system to use the vulncheck command."
    echo "Press Ctrl+C to cancel the installation if you don't want to proceed."
    echo ""
fi

# Ensure the install directory exists
if [[ ! -d "$INSTALL_DIR" ]]; then
    echo "Creating directory $INSTALL_DIR..."
    if [[ "$SYSTEM_WIDE" == true ]] && [[ "$OS" != "Windows" ]]; then
        sudo mkdir -p "$INSTALL_DIR"
    else
        mkdir -p "$INSTALL_DIR"
    fi
fi

# Copy the binary to the install directory
echo "Installing vulncheck to $INSTALL_DIR..."
if [[ "$OS" == "Windows" ]]; then
    mv "$EXTRACTED_FOLDER/bin/vulncheck.exe" "$INSTALL_DIR"
elif [[ "$SYSTEM_WIDE" == true ]]; then
    sudo mv "$EXTRACTED_FOLDER/bin/vulncheck" "$INSTALL_DIR"
else
    mv "$EXTRACTED_FOLDER/bin/vulncheck" "$INSTALL_DIR"
    # Make sure it's executable for local installs
    chmod +x "$INSTALL_DIR/vulncheck"
fi

# Clean up
cd ..
rm -rf "$TEMP_DIR"

echo "Installation complete!"
echo ""

if [[ "$OS" == "Windows" ]]; then
    echo "The vulncheck binary has been installed to $INSTALL_DIR."
    echo "This directory should already be in your system's PATH."
    echo "You can run it by typing 'vulncheck' in your command prompt or PowerShell."
elif [[ "$SYSTEM_WIDE" == true ]]; then
    echo "The vulncheck binary has been installed to $INSTALL_DIR, which is already in your system's PATH."
    echo "You can run it by typing 'vulncheck' in your terminal."
else
    echo "The vulncheck binary has been installed to $INSTALL_DIR."
    echo ""
    # Check if local bin is in PATH
    if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
        echo "⚠️  WARNING: $INSTALL_DIR is not in your PATH."
        echo ""
        echo "To use vulncheck, you need to add this directory to your PATH."
        echo "Add one of the following lines to your shell configuration file:"
        echo ""
        if [[ -f "$HOME/.bashrc" ]] || [[ "$SHELL" == *"bash"* ]]; then
            echo "For bash (~/.bashrc):"
            echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
        fi
        if [[ -f "$HOME/.zshrc" ]] || [[ "$SHELL" == *"zsh"* ]]; then
            echo "For zsh (~/.zshrc):"
            echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
        fi
        echo ""
        echo "After adding the line, reload your shell configuration:"
        echo "  source ~/.bashrc  # or ~/.zshrc"
        echo ""
        echo "Alternatively, you can run vulncheck using its full path:"
        echo "  $INSTALL_DIR/vulncheck"
    else
        echo "You can now use 'vulncheck' command."
    fi
fi
