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
    gomake clean
    cd "$d"
}

rm -rf $GOROOT/pkg/${GOOS}_${GOARCH}/fetch
rm -rf $GOROOT/pkg/${GOOS}_${GOARCH}/fetch.a
rm -rf $GOBIN/fetchd

for pkg in $PKGS
do mk pkg/$pkg
done

for cmd in $CMDS
do mk cmd/$cmd
done
