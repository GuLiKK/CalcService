# CalcService
Проект «CalcService» — это простой веб-сервис на Go, который вычисляет арифметические выражения. Он принимает POST-запрос с JSON-данными, где ключ "expression" хранит строку-выражение, а в ответ возвращает результат вычисления или ошибку.

Развёрнутый сервис слушает запросы на эндпоинте /api/v1/calculate.

Входные данные (JSON) должны содержать поле expression. Например:
{ "expression": "2+2*2" }

Возможные ответы:

Код 200 OK, если вычисление завершилось успешно. Тогда сервис вернёт JSON вида:
{ "result": "6" }
Код 422 Unprocessable Entity, если пользователь ввёл некорректное выражение, сделал деление на ноль или в тексте оказались лишние символы. Тогда в ответ придёт:
{ "error": "Expression is not valid" }
Код 500 Internal Server Error, если в процессе обработки возникла какая-то внутренняя ошибка. Тогда в ответ придёт:
{ "error": "Internal server error" }
Ниже несколько примеров использования с помощью curl (в формате, работающем на Linux или в Git Bash/WSL под Windows):

Пример успешного вычисления (получим ответ 6):
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'

Пример ошибки 422 из-за деления на ноль:
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2/0"}'

Пример ошибки 422 из-за недопустимого символа:
curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+a"}'

Пример 500 Internal Server Error воспроизвести сложно, но если внутренняя ошибка действительно произойдёт, в ответе будет:
{ "error": "Internal server error" }

Инструкция по запуску сервиса:

Установите Go (версии 1.18 или новее).
Склонируйте репозиторий: git clone https://github.com/ВАШ-ЛОГИН/ВАШ-РЕПО.git
Перейдите в папку проекта: cd ВАШ-РЕПО
Запустите сервер командой go run ./cmd/calc_service/...
Сервис по умолчанию слушает на http://localhost:8080. Отправляйте POST-запросы на http://localhost:8080/api/v1/calculate с JSON-телом, содержащим поле "expression".
