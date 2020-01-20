@echo off

set CURDIR=%~dp0
set PROTOC=%GOPATH%\bin\protoc.exe
set PROTO_DIR=%CURDIR%pbproto
set PBGO_DIR=%CURDIR%pbgo
set PLUGIN=gogofaster

echo.
echo,
echo PROTOC:%PROTOC%
echo CURDIR:%CURDIR%
echo PROTO_DIR:%PROTO_DIR%
echo PBGO_DIR:%PBGO_DIR%
echo;
echo;

if exist "%PBGO_DIR%" (
    for /f "delims=" %%i in ('dir /b "%PBGO_DIR%\*"') do (
        echo 删除文件 "%PBGO_DIR%\%%i"
        rmdir /q /s "%PBGO_DIR%\%%i"
    )
    echo 删除结束
    echo.
) else (
    echo 创建文件夹 %PBGO_DIR%
    md "%PBGO_DIR%"
)

rem goto labelexit

for /f "delims=" %%i in ('dir /b "%PROTO_DIR%\*.proto"') do (
    echo 生成 %%i
    if not exist "%PBGO_DIR%\%%~ni" (
        mkdir "%PBGO_DIR%\%%~ni"
    )
    "%PROTOC%" --proto_path="%PROTO_DIR%" --%PLUGIN%_out=paths=source_relative:"%PBGO_DIR%\%%~ni" %%i
rem   protoc --proto_path=./pbproto --go_out=paths=source_relative:./pbgo/GetGlobalState ./pbproto/GetGlobalState.proto
rem    echo %%~ni
rem    goto labelexit
)

rem :labelexit

pause
