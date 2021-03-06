#!/usr/bin/env bash
set -e
if [ -f env.bash ]
then . ./env.bash
else
    echo 1>&2 "! $0 must be run from the root directory"
    exit 1
fi

xcd() {
    echo
    cd $1
    echo --- cd $1
}

mk() {
    d=$PWD
    xcd $1
    gomake install
    cd "$d"
}

tst() {
    d=$PWD
    xcd $1
    gotest
    cd "$d"
}

for req in $PKG_REQS
do goinstall $req
done

for pkg in $PKGS
do mk pkg/$pkg
done

for cmd in $CMDS
do mk cmd/$cmd
done

if [ "$1" == "test" ]; then
    for pkg in $PKGS
    do tst pkg/$pkg
    done
fi
