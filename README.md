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
   ```json
   {
     "expression": "2+2*2"
   }
   ```
3. Приложение разбирает и вычисляет выражение, а затем возвращает результат в поле `result` или ошибку в поле `error`.

## Типы ответов

- **200 OK**  
  Приложение вернёт:
  ```json
  {
    "result": "6"
  }
  ```
  Если выражение корректно (например, `2+2*2`).

- **422 Unprocessable Entity**  
  Приложение вернёт:
  ```json
  {
    "error": "Expression is not valid"
  }
  ```
  Если выражение некорректно (лишние символы, неправильный синтаксис, деление на ноль и т.д.).

- **500 Internal Server Error**  
  Приложение вернёт:
  ```json
  {
    "error": "Internal server error"
  }
  ```
  Если внутри сервиса произошла непредвиденная ошибка.

## Примеры использования (curl)

Ниже несколько примеров запроса с помощью `curl` для **Windows** (cmd/PowerShell). Если вы используете **Git Bash/WSL** или **Linux**, можно заменить двойные кавычки на одинарные и убрать экранирование `\"`.

### 1. Успешный запрос (результат `6`):
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2+2*2\"}"
```
Ожидаемый ответ:
```json
{
  "result": "6"
}
```

### 2. Ошибка 422 (деление на ноль):
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2/0\"}"
```
Ожидаемый ответ:
```json
{
  "error": "Expression is not valid"
}
```

### 3. Ошибка 422 (некорректные символы):
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
  --header "Content-Type: application/json" ^
  --data "{\"expression\":\"2+a\"}"
```
Ожидаемый ответ:
```json
{
  "error": "Expression is not valid"
}
```

### 4. Ошибка 500 (внутренняя ошибка):
При сбое внутри самого приложения ответ будет:
```json
{
  "error": "Internal server error"
}
```

## Пример с `Invoke-RestMethod` (PowerShell)

Если не хочется возиться с кавычками в `curl` на Windows, можно использовать встроенный способ PowerShell:
```powershell
Invoke-RestMethod `
  -Uri "http://localhost:8080/api/v1/calculate" `
  -Method POST `
  -Body '{"expression": "2+2*2"}' `
  -ContentType "application/json"
```
Ответ придёт в виде объекта PowerShell, содержащего поля JSON.

## Как запустить проект

1. **Установите Go** (версии 1.18 или новее).
2. **Склонируйте репозиторий** (или скачайте архив) из GitHub:
   ```bash
   git clone https://github.com/<ВАШ-ЛОГИН>/<ИМЯ-РЕПОЗИТОРИЯ>.git
   ```
3. **Перейдите в папку** с проектом:
   ```bash
   cd <ИМЯ-РЕПОЗИТОРИЯ>
   ```
4. **Запустите** приложение:
   ```bash
   go run main.go
   ```
5. Сервис стартует на порту `8080`. Теперь можно отправлять **POST**-запросы на:
   ```
   http://localhost:8080/api/v1/calculate
   ```

## Дополнительно

- Если у вас **Git Bash** или **WSL**, то можно писать в стиле Linux без экранирования:
  ```bash
  curl --location 'http://localhost:8080/api/v1/calculate' \
    --header 'Content-Type: application/json' \
    --data '{
      "expression": "2+2*2"
    }'
  ```
  В таком случае пример выше с `^` и экранированными кавычками не нужен.
