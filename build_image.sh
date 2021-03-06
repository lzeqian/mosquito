#!/bin/bash
go build -x
docker build -t liaomin789/mosquito:1.0.0 .
