
@echo off
:start
cls
echo .������ע������
set /p COMMETSTR=

rem echo "%COMMETSTR%"
rem goto :exitlabel
git remote -v
git add .
git commit -m "%COMMETSTR%"
git push 
rem origin


:exitlabel
pause
