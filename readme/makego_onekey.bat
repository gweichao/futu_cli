@echo off

set CURDIR=%~dp0
set MAINPATH=E:\futu\api

:start
cls
echo.������汾�ţ����磺v3.13���ο� ��readme-���غʹ���pb go.txt��
set /p input_ver=
if "%input_ver%"=="" echo.������յİ汾�ţ� &goto :start

rem        if not exist "%MAINPATH%\github.com" mkdir "%MAINPATH%\github.com"
rem        echo.��ʼ��github���ذ汾
rem        cd /d "%MAINPATH%\github.com"
rem        mkdir "%input_ver%"
rem        cd %input_ver%
rem        git clone -b %input_ver% https://github.com/FutunnOpen/py-futu-api.git

        echo.�����汾�ļ�������go�ļ�
        mkdir "%MAINPATH%\%input_ver%\pbproto"
        copy %MAINPATH%\github.com\%input_ver%\py-futu-api\futu\common\pb\*.proto  %MAINPATH%\%input_ver%\pbproto\
        copy %MAINPATH%\ref\makego.bat %MAINPATH%\%input_ver%\
        cd /d %MAINPATH%\%input_ver%\
        call makego.bat
        echo.�ļ��μ� %MAINPATH%\%input_ver%\�����俽����Ŀ�깤�̽�һ��ʹ��
rem if "%str%"=="" (echo yes) else echo no
rem if not exist "!input_source!" echo.������·�������ڣ���&goto :start

pause
