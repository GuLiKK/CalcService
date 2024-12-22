# CalcService

Это простое веб-приложение на Go, которое умеет вычислять арифметические выражения из JSON-запроса и возвращать результат.

## Как пользоваться

1. Приложение запущено на порту `8080`, отвечает на адрес:
   ```
   http://localhost:8080/api/v1/calculate
   ```
2. Нужно отправлять POST-запрос с JSON-объектом вида:
   ```json
   {
     "expression": "2+2*2"
   }
   ```
3. Сервис вернёт либо:
   ```json
   {
     "result": "6"
   }
   ```
   если всё в порядке, либо:
   ```json
   {
     "error": "Expression is not valid"
   }
   ```
   (код 422), если выражение неправильное, или:
   ```json
   {
     "error": "Internal server error"
   }
   ```
   (код 500), если что-то сломалось внутри.

## Примеры (curl)

### Успешный запрос (6)
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
     --header "Content-Type: application/json" ^
     --data "{\"expression\":\"2+2*2\"}"
```

### Деление на ноль (422)
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
     --header "Content-Type: application/json" ^
     --data "{\"expression\":\"2/0\"}"
```

### Некорректные символы (422)
```bash
curl --location "http://localhost:8080/api/v1/calculate" ^
     --header "Content-Type: application/json" ^
     --data "{\"expression\":\"2+a\"}"
```

## Запуск

1. Склонируйте репозиторий или скачайте файлы.
2. В папке с проектом выполните:
   ```
   go run main.go
   ```
3. Всё, приложение работает на `http://localhost:8080`. Посылайте POST-запросы на `/api/v1/calculate`.

## Если вы на Windows в PowerShell

Можно использовать `Invoke-RestMethod`:
```powershell
Invoke-RestMethod `
  -Uri "http://localhost:8080/api/v1/calculate" `
  -Method POST `
  -Body '{"expression": "2+2*2"}' `
  -ContentType "application/json"
```

Или запустить Git Bash/WSL и использовать команды, как в Linux.  
