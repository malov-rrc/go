**P0 — первая CLI-программа**
Программа считает статистику по результатам автотестов.

Вход — текстовый файл, одна строка = один результат, формат `имя_теста,СТАТУС`:
```
test_login,PASS
test_logout,FAIL
test_password_reset,PASS
test_signup,SKIP
```
Пустые строки пропускаются, невалидные статусы подсчитываются

Запуск: `go run main.go results.txt` (путь к файлу — первый аргумент через `os.Args`).

Вывод в stdout:
```
Total: 7
PASS: 2 (28.6%)
FAIL: 1 (14.3%)
SKIP: 1 (14.3%)
unknownStatus: 2 (28.6%)
Invalid lines: 1
```