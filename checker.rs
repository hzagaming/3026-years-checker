// 3026 Years Checker — Rust 命令行版本
// 用法: rustc -O checker.rs -o checker && ./checker 2026
use std::env;
use std::process;

const TARGET: i64 = 3026;

fn main() {
    let arg = match env::args().nth(1) {
        Some(v) => v,
        None => {
            eprintln!("用法: checker <年份>");
            process::exit(1);
        }
    };
    let year: i64 = match arg.trim().parse() {
        Ok(v) => v,
        Err(_) => {
            eprintln!("请给一个整数年份");
            process::exit(1);
        }
    };
    let delta = year - TARGET;
    if delta == 0 {
        println!("{} 就是 {} 年", year, TARGET);
    } else if delta > 0 {
        println!("{} 在 {} 年之后 {} 年", year, TARGET, delta);
    } else {
        println!("{} 在 {} 年之前 {} 年", year, TARGET, -delta);
    }
}
