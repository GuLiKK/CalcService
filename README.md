# CalcService

**CalcService** — это веб-приложение на Go, которое принимает арифметические выражения в формате JSON и возвращает результат их вычисления. Кроме того, оно обрабатывает различные ошибки: от деления на ноль до некорректных символов в выражении.

## Как это работает

1. Приложение слушает запросы на порту `8080` по адресу:
```
http://localhost:8080
```
2. Главный эндпоинт:
```
/api/v1/calculate
```
Сюда нужно отправлять **POST**-запрос в формате JSON, например:
```
{
  "expression": "2+2*2"
}
```
3. Приложение разбирает и вычисляет выражение, а затем возвращает результат в поле result или ошибку в поле error.
Типы ответов
```
200 OK
```
Приложение вернёт:
```
{
  "result": "6"
}
```
Если выражение корректно (например, 2+2*2).
```
422 Unprocessable Entity
```
Приложение вернёт:
```
{
  "error": "Expression is not valid"
}
```
Если выражение некорректно (лишние символы, неправильный синтаксис, деление на ноль и т.д.).
```
500 Internal Server Error
```
Приложение вернёт:
```
{
  "error": "Internal server error"
}
```
Если внутри сервиса произошла непредвиденная ошибка.

Примеры использования (curl)
Ниже приведены примеры для Windows (cmd/PowerShell). Если вы используете Git Bash/WSL или Linux, можно заменить двойные кавычки на одинарные и убрать экранирование \".

Успешный запрос (результат 6):
```
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2+2*2\"}"
```
Ожидаемый ответ:
```
{
  "result": "6"
}
```
Ошибка 422 (деление на ноль):
```
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2/0\"}"
```
Ожидаемый ответ:
```
{
  "error": "Expression is not valid"
}
```
Ошибка 422 (некорректные символы):
```
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2+a\"}"
```
Ожидаемый ответ:
```
{
  "error": "Expression is not valid"
}
```
Ошибка 500 (внутренняя ошибка сервера):
При сбое внутри самого приложения ответ будет:
```
{
  "error": "Internal server error"
}
```
Как запустить проект
1. Склонируйте репозиторий (или скачайте архив) из GitHub:
```
git clone https://github.com/GuLiKK/CalcService.git
```
2. Перейдите в папку с проектом:
```
cd CalcService
```
3. Запустите приложение:
```
go run ./cmd/main.go
```
4. Сервис стартует на порту 8080. Теперь можно отправлять POST-запросы на:
```
http://localhost:8080/api/v1/calculate
```
Дополнительно
Если в PowerShell вы не хотите возиться с экранированием, можете использовать:
```
Invoke-RestMethod `
  -Uri "http://localhost:8080/api/v1/calculate" `
  -Method POST `
  -Body '{"expression": "2+2*2"}' `
  -ContentType "application/json"
```
Если у вас Git Bash или WSL, то можно писать в стиле Linux:
```
curl --location 'http://localhost:8080/api/v1/calculate' \
  --header 'Content-Type: application/json' \
  --data '{
    "expression": "2+2*2"
  }'
```
Теперь вы можете тестировать работу сервиса, отправляя различные выражения и проверяя результаты.
