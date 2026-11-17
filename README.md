# 3026 Years Checker

输入年份，检查它距离 3026 年有多远。同一个功能用四种语言实现。

## 组成

| 文件 | 语言 | 说明 |
| --- | --- | --- |
| `server.go` | Go | 静态站 + `/api/check` 接口，默认 `127.0.0.1:8086` |
| `web/index.html` | HTML/JS | 页面本体，调用接口 |
| `server.php` | PHP | 同功能接口，`php -S 127.0.0.1:8087 server.php` |
| `checker.rs` | Rust | 命令行版本，`rustc -O checker.rs -o checker` |
| `run.sh` | Bash | 统一启动：`./run.sh [go\|php\|rust]` |

## 用法

```bash
./run.sh            # Go 版本，打开 http://127.0.0.1:8086
./run.sh php        # PHP 版本
./run.sh rust       # 编译 Rust 命令行版本
```

接口：

```bash
curl "http://127.0.0.1:8086/api/check?year=2026"
{"year":2026,"delta":-1000,"text":"2026 在 3026 年之前 1000 年"}
```

## 验证状态

- Go：`go build` + 接口实测通过
- Bash：`bash -n` 语法检查通过
- PHP / Rust：本机没装 `php`、`rustc`，代码未实测
