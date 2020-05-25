#!/usr/bin/env sh

if [ -z "$1" ]; then
  modversion patch++
elif [ "$1" = "M" ]; then
  modversion major++
elif [ "$1" = "m" ]; then
  modversion minor++
elif [ "$1" = "p" ]; then
  modversion patch++
fi

if [ -f VERSION ]; then
  V=$(cat VERSION)
fi

git tag "$V"
git push --tags
