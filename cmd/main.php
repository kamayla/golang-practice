<?php

$startTime = microtime(true);

for ($i = 0; $i < 3000000; $i++) {
    echo "{$i}処理\n";
}

$executeTime = microtime(true) - $startTime;

echo "{$executeTime}秒";