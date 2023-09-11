#!/bin/bash

OS=$(uname -s)

ARCH=$(uname -m)

VERSION=1.0.1

URL="https://github.com/subzero-cli/subzero/releases/download/v$VERSION/"

TMP_DIR=$(mktemp -d)

case "$OS" in
    "Linux")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_linux_amd64.tar.gz" ;;
            "i386") FILE="subzero_${VERSION}_linux_386.tar.gz" ;;
            "aarch64") FILE="subzero_${VERSION}_linux_arm64.tar.gz" ;;
            *) echo "Unsupported arch: $ARCH"; exit 1 ;;
        esac
        ;;
    "Darwin")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_darwin_amd64.tar.gz" ;;
            "arm64") FILE="subzero_${VERSION}_darwin_arm64.tar.gz" ;;
            *) echo "Unsupported arch: $ARCH"; exit 1 ;;
        esac
        ;;
    "WindowsNT")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_windows_amd64.tar.gz" ;;
            "i386") FILE="subzero_${VERSION}_windows_386.tar.gz" ;;
            "arm64") FILE="subzero_${VERSION}_windows_arm64.tar.gz" ;;
            *) echo "Unsupported arch: $ARCH"; exit 1 ;;
        esac
        ;;
    *)
        echo "Failed to automatic match os, please check releases page https://github.com/subzero-cli/subzero/releases: $OS"
        exit 1
        ;;
esac

if command -v curl &> /dev/null
then
    curl -sSL "$URL$FILE" -o "$TMP_DIR/$FILE"
else
    if command -v wget &> /dev/null
    then
        wget -q -O "$TMP_DIR/$FILE" "$URL$FILE"
    else
        echo "Not found 'curl' or 'wget'. Please install."
        exit 1
    fi
fi

tar -xzvf "$TMP_DIR/$FILE" -C "$TMP_DIR"

sudo cp "$TMP_DIR//subzero" /usr/local/bin

rm -rf "$TMP_DIR"

echo "Instalação concluída com sucesso!"

subzero