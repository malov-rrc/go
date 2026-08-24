# Go Practice: 10 Tasks

Структура:
```
task1/main.go  - Конвертер температур
task2/main.go  - Високосный год
task3/main.go  - Таблица умножения
task4/main.go  - Классификатор оценки (1–5)
task5/math.go + math_test.go - Sum/Max/Min/IsPrime + тесты. работаем с []int и int (не float64) — проще для первого знакомства с testing. И раз в дорожной карте месяц 1 явно требует table-driven tests как checkpoint — задачи 5 и 10 должны использовать именно table-driven формат ([]struct{...} + for _, tc := range cases), а не отдельные TestX на каждый кейс. Это тоже фиксирую явно, потому что «тест на каждую» можно прочитать и как «одна функция теста на функцию», и как «table-driven на всё сразу».
task6/main.go  - Rectangle/Circle с Area/Perimeter. в main создать Rectangle{3,4} и Circle{radius:5}, вывести Area/Perimeter каждой с округлением до 2 знаков, например: Rectangle: area=12.00, perimeter=14.00.
task7/main.go  - Swap через указатели
task8/main.go  - Фильтрация (чётные) и сортировка слайса. Источник данных - захардкоженный slice в main, например []int{5,3,8,2,9,4,7,1,6}. Вывод: исходный slice → отфильтрованный (только чётные) → отсортированный по возрастанию, три строки.
task9/main.go  - Частота слов в строке (map). Источник данных - захардкоженная строка-константа (2–3 предложения) в main. Вывод — построчно, отсортировано по алфавиту: слово: количество.
task10/divide.go + divide_test.go - Безопасное деление (error) + тест
```

Запуск:
```bash
go run task1/main.go
go run task2/main.go
# ...
```

Тесты (task5, task10):
```bash
cd task5 && go test ./...
cd task10 && go test ./...
```