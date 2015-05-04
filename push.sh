#! /bin/bash -

# add
git add *.go

# tag
git commit -m `date +%F_%T`

#update
#git push

git remote set-url origin  git@github.com:zhaoming200808/go_tutorial.git
