@echo off
start cmd /k "go run cmd\runBH\BHmain.go"
start cmd /k "go run cmd\runOrchestrator\Omain.go"
start cmd /k "go run cmd\runDemon\Dmain.go"
pause
