#!/bin/bash

# Stop on error
set -e

APP_NAME="excel-protector"
OUTPUT_DIR="build"

# Check if platform and architecture arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <platform> <architecture>"
    echo "Example: $0 linux amd64"
    exit 1
fi

GOOS=$1
GOARCH=$2

# Clean previous build for this target
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# Set the output file name based on the platform and architecture
output_name="$OUTPUT_DIR/${APP_NAME}-${GOOS}-${GOARCH}"

# Add .exe extension for Windows builds
if [ "$GOOS" == "windows" ]; then
    output_name+='.exe'
fi

# Build for the specified platform and architecture
echo "Building for $GOOS/$GOARCH..."
GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_name" src/main.go

echo "Build completed successfully! Output file: $output_name"
