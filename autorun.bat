@echo off
start cmd /k "cd C:\Users\user\BH_Lu_3\BH_Lu_3\cmd\runBH && set PATH=%PATH%;C:\Program Files\Go\bin && go run BHmain.go"
start cmd /k "cd C:\Users\user\BH_Lu_3\BH_Lu_3\cmd\runOrchestrator && set PATH=%PATH%;C:\Program Files\Go\bin && go run Omain.go"
start cmd /k "cd C:\Users\user\BH_Lu_3\BH_Lu_3\cmd\runDemon && set PATH=%PATH%;C:\Program Files\Go\bin && go run Dmain.go"
pause