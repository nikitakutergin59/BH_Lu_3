@echo off
start cmd /k "go run cmd\runBH\main.go"
start cmd /k "go run cmd\runOrchestrator\main.go"
start cmd /k "go run cmd\runDemon\main.go"
pause
