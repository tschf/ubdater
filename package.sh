#!/bin/bash
#build.sh: Build the program for multiple platforms for releases.
#Move into the dist folder.
#A list of possible platforms can be seen: https://golang.org/doc/install/source#environment
#Also, referenced:
#- https://dave.cheney.net/2015/08/22/cross-compilation-with-go-1-5
#- https://github.com/golang/go/wiki/WindowsCrossCompiling
set -e
APP_NAME=ubdater
BASEDIR=$GOPATH/src/github.com/tschf/$APP_NAME/build

rm -rf ${BASEDIR}

echo "Building for OS X (darwin) 32 bit"
export GOOS=darwin
export GOARCH=386
mkdir -p ${BASEDIR}/osx_32
go build -o "$BASEDIR/osx_32/$APP_NAME"
echo "Done"

echo "Building for OS X (darwin) 64 bit"
export GOOS=darwin
export GOARCH=amd64
mkdir -p ${BASEDIR}/osx_64
go build -o "$BASEDIR/osx_64/$APP_NAME"
echo "Done"

echo "Building for Linux 32 bit"
export GOOS=linux
export GOARCH=386
mkdir -p ${BASEDIR}/linux_32
go build -o "$BASEDIR/linux_32/$APP_NAME"
echo "Done"

echo "Building for Linux 64 bit"
export GOOS=linux
export GOARCH=amd64
mkdir -p ${BASEDIR}/linux_64
go build -o "$BASEDIR/linux_64/$APP_NAME"
echo "Done"

echo "Building for Windows 32 bit"
export GOOS=windows
export GOARCH=386
mkdir -p ${BASEDIR}/win_32
go build -o "$BASEDIR/win_32/$APP_NAME.exe"
echo "Done"

echo "Building for Windows 64 bit"
export GOOS=windows
export GOARCH=amd64
mkdir -p ${BASEDIR}/win_64
go build -o "$BASEDIR/win_64/$APP_NAME.exe"
echo "Done"


#Next, make archives of these for each platform
echo "Creating archives"
cd ${BASEDIR}
tar cvzf linux.tgz linux_32/ linux_64/ --remove-files
zip -mr windows.zip win_32/ win_64/ 
tar cvzf osx.tgz osx_32/ osx_64/ --remove-files

echo "All finished. Now, go release!"

