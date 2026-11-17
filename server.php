<?php
// 3026 Years Checker — PHP 版本接口
// 用法: php -S 127.0.0.1:8087 server.php
declare(strict_types=1);

const TARGET = 3026;

function describe(int $year): array
{
    $delta = $year - TARGET;
    if ($delta === 0) {
        $text = $year . ' 就是 ' . TARGET . ' 年';
    } elseif ($delta > 0) {
        $text = $year . ' 在 ' . TARGET . ' 年之后 ' . $delta . ' 年';
    } else {
        $text = $year . ' 在 ' . TARGET . ' 年之前 ' . (-$delta) . ' 年';
    }
    return ['year' => $year, 'delta' => $delta, 'text' => $text];
}

header('Content-Type: application/json; charset=utf-8');

$path = parse_url($_SERVER['REQUEST_URI'] ?? '/', PHP_URL_PATH);
if ($path === '/api/check') {
    $raw = $_GET['year'] ?? '';
    if (!is_numeric($raw)) {
        echo json_encode(['error' => 'year 需要是整数']);
        return;
    }
    echo json_encode(describe((int) $raw), JSON_UNESCAPED_UNICODE);
    return;
}

readfile(__DIR__ . '/web/index.html');
