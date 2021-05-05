#!/usr/bin/env bash

error_prebuild=1
error_build=2
error_package=3
error_test=4

if [ x"" = x"$PROJECT_ROOT" ]; then
  # Configure for regular build
  PROJECT_ROOT=$PWD
  BUILD_PATH=$PWD/build
  DIST_PATH=$PWD/dist
else
  # Configure for Docker build
  BUILD_PATH=/build
  DIST_PATH=/dist
fi

for p in linux win mac; do
  mkdir -p $BUILD_PATH/$p
done

go run tbx.go dev build info
if [[ $? == 0 ]]; then
  echo "Build information created."
else
  exit $error_prebuild
fi

function build_and_package() {
    platform_alias=$1
    goos=$2
    goarch=$3
    bin_path=$4
    bin_linux=$5

    echo Building: $platform_alias [$goos][$goarch]
    CGO_ENABLED=0 GOOS=$goos GOARCH=$goarch go build -o $bin_path github.com/watermint/toolbox
    if [[ $? == 0 ]]; then
      echo "The binary created: $bin_path"
    else
      exit $error_build
    fi

    $5 dev build package -build-path $bin_path -dest-path $DIST_PATH -deploy-path /watermint-toolbox-build -platform $platform_alias
    if [[ $? == 0 ]]; then
      echo "The binary packaged"
    else
      exit $error_package
    fi
}

LINUX_BIN=$BUILD_PATH/linux/tbx
build_and_package linux linux   amd64 $BUILD_PATH/linux/tbx   $LINUX_BIN
build_and_package win   windows amd64 $BUILD_PATH/win/tbx.exe $LINUX_BIN
build_and_package mac    darwin amd64 $BUILD_PATH/mac/tbx     $LINUX_BIN
