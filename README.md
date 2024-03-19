# Профильное задание. Стажировка VK
> Профильное задание на позицию "Программист-разработчик" для стажировки VK.
Реализация REST API сервиса, который засчитывает задания для пользователя. 

## Используемые технологии
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Swagger](https://img.shields.io/badge/-Swagger-%23Clojure?style=for-the-badge&logo=swagger&logoColor=white)

## Реализованные требования к проекту
Сервис имеет 5 апи-метода:
- `[post]../api/users/` для создания пользователя.
- `[get]../api/users/:id/balance` для получения баланса пользователя и истории выполненых им заданий.
- `[post]../api/quests/` для создания квестов.
- `[get]../api/quests/` для получения квестов.
- `[post]../task-progres/` для завершения задания и начисления баланса пользователя.

> [!NOTE]
> В сервисе реализована система квестов (наборов заданий). Квест может включать в себя как одно задание, так и множество, но минимум одно задание обязательно. Баллы начисляются как за выполнение отдельных заданий, так и за прохождение квеста в целом при выполнении всех заданий в нем. Задания можно делать как одноразовыми, так и многоразовыми.

### Другое
* Предоставлена спецификация на API в формате Swagger 2.0.
* Логирование. В лог попадает базовая информация об обрабатываемых запросах и ошибках.
* Dockerfile для сборки образа.
* docker-compose для запуска окружения с работающим приложением и СУБД.

## Локальный запуск проекта

> 1. Запуск проекта на локальном сервере производиться командой 
```
docker-compose up -d
```
> 2. Для миграции используйте команду:
```
migrate -path ./schema -database 'postgresql://root:root@localhost:5432/task_tracking_db?sslmode=disable' up
```
> 3. Swagger на локальном сервере доступен по URL:
```
http://127.0.0.1:8080/swagger/index.html
```

