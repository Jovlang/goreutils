#!/bin/sh

utils="rev pwd mv mkdir rmdir ls"
target="/usr/local/bin"

case $1 in
("")
    8g shared.go || exit 1
    for i in $utils; do
        8g $i.go || exit 1
        8l -o $i $i.8 || exit 1
        rm $i.8
    done
    ;;
(install)
    cp $utils $target
    ;;
(clean)
    rm -f shared.8 $utils
    ;;
(*)
    echo "usage: $0 [install|clean]" >&2
    exit 1
    ;;
esac
