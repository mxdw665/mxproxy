#!/bin/bash

PROJECT_NAME="mxproxy"
MAIN_GO="main.go"
OUTPUT_DIR="build"

mkdir -p "$OUTPUT_DIR"

build_for() {
    local GOOS=$1
    local GOARCH=$2
    local EXT=$3
    local OUTPUT_FILE="${OUTPUT_DIR}/${PROJECT_NAME}_${GOOS}_${GOARCH}${EXT}"

    echo "正在打包 GOOS=${GOOS} GOARCH=${GOARCH}..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_FILE" "$MAIN_GO"
    
    if [ $? -eq 0 ]; then
        echo "打包成功: $OUTPUT_FILE"
    else
        echo "打包失败: GOOS=${GOOS} GOARCH=${GOARCH}"
    fi
}

build_for "linux" "amd64" ""
build_for "linux" "386" ""
build_for "linux" "arm64" ""
build_for "darwin" "amd64" ""
build_for "darwin" "arm64" ""
build_for "windows" "amd64" ".exe"
build_for "windows" "386" ".exe"

echo "打包完成。所有文件都已放在 ${OUTPUT_DIR} 目录中。"