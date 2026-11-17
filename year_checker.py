#!/usr/bin/env python3
"""判断一个年份相对 3026 年的位置。

    python3 year_checker.py 3026
"""

from __future__ import annotations

import sys

TARGET = 3026


def check(year: int) -> str:
    delta = year - TARGET
    if delta == 0:
        return f"{year} 就是 {TARGET} 年"
    side = "之后" if delta > 0 else "之前"
    return f"{year} 在 {TARGET} 年{side} {abs(delta)} 年"


def main() -> None:
    if len(sys.argv) != 2:
        sys.exit("用法: python3 year_checker.py <年份>")
    try:
        year = int(sys.argv[1])
    except ValueError:
        sys.exit("请给一个整数年份")
    print(check(year))


if __name__ == "__main__":
    main()
