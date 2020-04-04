#!/bin/sh

# @echo off
# unable to read askpass response from '/usr/libexec/openssh/gnome-ssh-askpass
unset SSH_ASKPASS
git config --global http.sslVerify false
#:start
#cls
#rem echo "%COMMETSTR%"
#rem goto :exitlabel

git remote -v
git pull origin

#:exitlabel
#pause

read -p "Press enter key to exit " var
