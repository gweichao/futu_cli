@echo off

set CURDIR=%~dp0
set MAINPATH=E:\futu\api

:start
cls
echo.请输入版本号，比如：v3.13，参考 “readme-下载和创建pb go.txt”
set /p input_ver=
if "%input_ver%"=="" echo.你输入空的版本号！ &goto :start

rem        if not exist "%MAINPATH%\github.com" mkdir "%MAINPATH%\github.com"
rem        echo.开始从github下载版本
rem        cd /d "%MAINPATH%\github.com"
rem        mkdir "%input_ver%"
rem        cd %input_ver%
rem        git clone -b %input_ver% https://github.com/FutunnOpen/py-futu-api.git

        echo.拷贝版本文件，生成go文件
        mkdir "%MAINPATH%\%input_ver%\pbproto"
        copy %MAINPATH%\github.com\%input_ver%\py-futu-api\futu\common\pb\*.proto  %MAINPATH%\%input_ver%\pbproto\
        copy %MAINPATH%\ref\makego.bat %MAINPATH%\%input_ver%\
        cd /d %MAINPATH%\%input_ver%\
        call makego.bat
        echo.文件参见 %MAINPATH%\%input_ver%\，将其拷贝到目标工程进一步使用
rem if "%str%"=="" (echo yes) else echo no
rem if not exist "!input_source!" echo.你输入路径不存在！！&goto :start

pause
