#!/bin/bash

## Define package name
PKG_NAME='proctel'

## Build the main
go build main.go

## Rename main as a package name
mv -v main "${PKG_NAME}"

## Install the package in the rtt directory
#sudo mkdir -p "${module_path}/bin"
#sudo cp -vf "${PKG_NAME}" "${pkg_path}"