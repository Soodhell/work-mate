# Task Manager API

HTTP API для управления длительными I/O bound задачами, написанный на Go.

## Немного о проекте

Тз довольно размыто (этим и интересно), поэтому решил сделать такое рещение данного проекта

## URL 

> GET [localhost://tasks/](localhost://tasks/) (получить всё)  
> GET [localhost://tasks/{id}](localhost://tasks/{id}) (получить опереденную таску)  
> POST [localhost://tasks/](localhost://tasks/) (добавить таску)  
> DELETE [localhost://tasks/{id}](localhost://tasks/{id}) (удалить таску)  


## Запуск сервера

1. Убедитесь, что у вас установлен Go (версия 1.22 или выше)
2. Клонируйте репозиторий
3. Перейдите в директорию проекта
4. Запустите сервер:

```bash
go run cmd/adpp/main.go
```
## docker
Сборка докер файла
```bash
docker build -t task-manager .
docker run -d -p 8080:8080 --name task-manager-container task-manager

