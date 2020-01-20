@echo off

set CURDIR=%~dp0
set MAINPATH=E:\futu\api

:start
cls
echo.请输入版本号，比如：v3.13，参考 “readme-下载和创建pb go.txt”
set /p input_ver=
if "%input_ver%"=="" echo.你输入空的版本号！ &goto :start

        if not exist "%MAINPATH%\github.com" mkdir "%MAINPATH%\github.com"
        echo.开始从github下载版本
        cd /d "%MAINPATH%\github.com"
        mkdir "%input_ver%"
        cd %input_ver%
        git clone -b %input_ver% https://github.com/FutunnOpen/py-futu-api.git
        echo.文件参见 %MAINPATH%\github.com\%input_ver%

rem       echo.拷贝版本文件，生成go文件
rem       mkdir "%MAINPATH%\%input_ver%\pb.proto"
rem       copy %MAINPATH%\github.com\%input_ver%\py-futu-api\futu\common\pb\*.proto  %MAINPATH%\%input_ver%\pb.proto\
rem       copy %MAINPATH%\ref\makego.bat %MAINPATH%\%input_ver%\
rem       cd /d %MAINPATH%\%input_ver%\
rem       call makego.bat
rem       echo.文件参见 %MAINPATH%\%input_ver%\

rem if "%str%"=="" (echo yes) else echo no
rem if not exist "!input_source!" echo.你输入路径不存在！！&goto :start

pause
