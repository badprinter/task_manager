# Task Manager (Go + PostgreSQL)  
Микросервис для управления задачами с REST API.  

## 📌 Функционал  
- Создание, редактирование, удаление задач через API.  
- Миграции БД с помощью **Goose**.  
- Конфигурация через `.env` (реализуется).  
- Запуск в Docker-контейнере.  

## 🛠 Технологии  
- **Язык**: Go (Fiber для роутинга).  
- **База данных**: PostgreSQL (Docker).  
- **Миграции**: Goose.  
- **Инфраструктура**: Docker Compose.  

## Запуск
**postgresql** запускается через *docker compose*.
Сама программа **app** запускается через *go build* (TODO).
Файл конфигурации _internal/config/config.go_ (TODO .env).

## Скриншот
![](readme/db_image.png)
