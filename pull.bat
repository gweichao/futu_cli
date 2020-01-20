
@echo off
:start
cls
rem echo "%COMMETSTR%"
rem goto :exitlabel
git remote -v
git pull origin

:exitlabel
pause
