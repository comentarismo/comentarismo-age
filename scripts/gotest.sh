#!/usr/bin/env bash

LEARNAGE=true godep go test -v $(go list ./age | grep -v /vendor/);
godep go test -v $(go list ./lang | grep -v /vendor/);
godep go test -v $(go list ./server | grep -v /vendor/);
