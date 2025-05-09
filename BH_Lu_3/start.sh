#!/bin/bash

# Запуск всех трёх приложений параллельно
./runBH&
./runDemon &
./runOrchestrator &

# Ожидание завершения всех процессов
wait