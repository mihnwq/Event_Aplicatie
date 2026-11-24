@echo off
cd /d "%~dp0"
go run tools\pack.go
if errorlevel 1 exit /b 1
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -trimpath -ldflags "-H=windowsgui -s -w" -o SIDE_CHALLENGE_3000.exe windows\main.go windows\assets.go
if errorlevel 1 exit /b 1
echo Build OK
