FROM golang:1.24.3

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Устанавливаем зависимости для sqlite3
RUN apt-get update && apt-get install -y gcc musl-dev

# Копируем файлы проекта в контейнер
COPY . .

# Загружаем зависимости Go
RUN go mod download

# Сборка первого приложения (rubBH)
RUN go build -o runBH ./cmd/runBH/BHmain.go

# Сборка второго приложения (runDemon)
RUN go build -o runDemon ./cmd/runDemon/Dmain.go

# Сборка третьего приложения (runOrchestrator)
RUN go build -o runOrchestrator ./cmd/runOrchestrator/Omain.go

# Копируем скрипт для запуска всех приложений
COPY start.sh .

# Делаем скрипт исполняемым
RUN chmod +x start.sh

# Указываем запуск скрипта
CMD ["./start.sh"]