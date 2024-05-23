#!/usr/bin/env bash

set -euo pipefail

project_dir=$(pwd)

cd "${project_dir}/third_party/proto/" && \
buf generate && \
cd ../ && \
cp -r github.com/sentinel-official/hub/v12/third_party/* ./ && \
rm -rf github.com/

cd "${project_dir}/proto/" && \
buf generate && \
cd ../ && \
cp -r github.com/sentinel-official/hub/v12/* ./ && \
rm -rf github.com/
