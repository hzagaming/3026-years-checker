#!/bin/bash
# 3026 Years Checker 启动脚本
set -e
cd "$(dirname "$0")"

port="${PORT:-8086}"

case "${1:-go}" in
  go)
    echo "Go 版本 -> http://127.0.0.1:${port}"
    exec go run server.go
    ;;
  php)
    if ! command -v php >/dev/null 2>&1; then
      echo "未安装 php" >&2
      exit 1
    fi
    echo "PHP 版本 -> http://127.0.0.1:${port}"
    exec php -S "127.0.0.1:${port}" server.php
    ;;
  rust)
    if ! command -v rustc >/dev/null 2>&1; then
      echo "未安装 rustc" >&2
      exit 1
    fi
    rustc -O checker.rs -o checker
    echo "已编译 checker，用法: ./checker 2026"
    ;;
  *)
    echo "用法: ./run.sh [go|php|rust]" >&2
    exit 1
    ;;
esac
