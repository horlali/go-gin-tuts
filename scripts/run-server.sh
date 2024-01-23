#!/bin/sh

set -e

# Go to root folder
cd $(dirname $0)/..

CompileDaemon -command="./go-gin-tuts ."
