#!/bin/bash

# Verifica o sistema operacional
OS=$(uname -s)

# Verifica a arquitetura
ARCH=$(uname -m)

# Verifica se a versão foi fornecida como argumento
if [ $# -eq 0 ]; then
    echo "Por favor, forneça a versão do subzero como argumento."
    exit 1
fi

# Obtém a versão a partir do primeiro argumento
VERSION=$1

# Define o URL do release
URL="https://github.com/subzero-cli/subzero/releases/download/v$VERSION/"

# Define o diretório temporário para baixar os arquivos
TMP_DIR=$(mktemp -d)

# Baixa o arquivo apropriado para o sistema operacional e arquitetura
case "$OS" in
    "Linux")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_linux_amd64.tar.gz" ;;
            "i386") FILE="subzero_${VERSION}_linux_386.tar.gz" ;;
            "aarch64") FILE="subzero_${VERSION}_linux_arm64.tar.gz" ;;
            *) echo "Arquitetura não suportada: $ARCH"; exit 1 ;;
        esac
        ;;
    "Darwin")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_darwin_amd64.tar.gz" ;;
            "arm64") FILE="subzero_${VERSION}_darwin_arm64.tar.gz" ;;
            *) echo "Arquitetura não suportada: $ARCH"; exit 1 ;;
        esac
        ;;
    "WindowsNT")
        case "$ARCH" in
            "x86_64") FILE="subzero_${VERSION}_windows_amd64.tar.gz" ;;
            "i386") FILE="subzero_${VERSION}_windows_386.tar.gz" ;;
            "arm64") FILE="subzero_${VERSION}_windows_arm64.tar.gz" ;;
            *) echo "Arquitetura não suportada: $ARCH"; exit 1 ;;
        esac
        ;;
    *)
        echo "Sistema operacional não suportado: $OS"
        exit 1
        ;;
esac

# Baixa o arquivo
wget "$URL$FILE" -P "$TMP_DIR"

# Extrai o conteúdo do arquivo
tar -xzvf "$TMP_DIR/$FILE" -C "$TMP_DIR"

# Copia o binário para /usr/bin
sudo cp "$TMP_DIR/subzero" /usr/bin

# Limpa o diretório temporário
rm -rf "$TMP_DIR"

echo "Instalação concluída com sucesso!"
