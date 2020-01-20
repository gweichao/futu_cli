
@echo off
:start
cls
echo .«Î ‰»Î◊¢ Õ√Ë ˆ
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
