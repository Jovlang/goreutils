#!/bin/sh

./pwd
echo
./ls
echo
touch foo
./mv -v foo bar
rm bar
echo
echo "This is a line, med vær og vår.." | ./rev
echo
./mkdir -p foo/bar/baz/buh
./ls foo/bar/baz
rm -r foo
