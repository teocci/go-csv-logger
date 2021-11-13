#!/bin/bash
## Define package name
PKG_NAME='proctel'

## Define local path
LOCAL_PATH='/usr/local'

## Define module
MODULE_NAME='rtt'

## Define installation path
module_path="${LOCAL_PATH}/${MODULE_NAME}"

## Define package path
pkg_path="${module_path}/bin/${PKG_NAME}"

## Git pull
git pull

## Build the main
go build main.go

## Rename main as a package name
mv -v main "${PKG_NAME}"

## Install the package in the rtt directory
sudo mkdir -p "${module_path}/bin"
sudo cp -vf "${PKG_NAME}" "${pkg_path}"
