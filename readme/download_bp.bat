@echo off

set CURDIR=%~dp0
set MAINPATH=E:\futu\api

:start
cls
echo.������汾�ţ����磺v3.13���ο� ��readme-���غʹ���pb go.txt��
set /p input_ver=
if "%input_ver%"=="" echo.������յİ汾�ţ� &goto :start

        if not exist "%MAINPATH%\github.com" mkdir "%MAINPATH%\github.com"
        echo.��ʼ��github���ذ汾
        cd /d "%MAINPATH%\github.com"
        mkdir "%input_ver%"
        cd %input_ver%
        git clone -b %input_ver% https://github.com/FutunnOpen/py-futu-api.git
        echo.�ļ��μ� %MAINPATH%\github.com\%input_ver%

rem       echo.�����汾�ļ�������go�ļ�
rem       mkdir "%MAINPATH%\%input_ver%\pb.proto"
rem       copy %MAINPATH%\github.com\%input_ver%\py-futu-api\futu\common\pb\*.proto  %MAINPATH%\%input_ver%\pb.proto\
rem       copy %MAINPATH%\ref\makego.bat %MAINPATH%\%input_ver%\
rem       cd /d %MAINPATH%\%input_ver%\
rem       call makego.bat
rem       echo.�ļ��μ� %MAINPATH%\%input_ver%\

rem if "%str%"=="" (echo yes) else echo no
rem if not exist "!input_source!" echo.������·�������ڣ���&goto :start

pause
