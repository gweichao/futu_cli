#!/bin/sh

read -p "input comment:" COMMETSTR
echo "";echo "";
# echo $COMMETSTR
# read -p "Press any key to continue." var
# exit 0

# @echo off
#:start
#cls
#echo .«Î ‰»Î◊¢ Õ√Ë ˆ
#set /p COMMETSTR=
#rem echo "%COMMETSTR%"
#rem goto :exitlabel

unset SSH_ASKPASS
git config --global http.sslVerify false

git remote -v
git add .
git commit -m "$COMMETSTR"
git push 

# rem origin

#:exitlabel
# pause
echo "";echo "";
read -p "Press enter key to exit " var
echo "";echo "";
