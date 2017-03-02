#!/usr/bin/env bash

goagen main -d github.com/couchbase/cbbootstrap/design -o ./controllers

rm controllers/main.go

goagen app -d github.com/couchbase/cbbootstrap/design -o ./goa
goagen client -d github.com/couchbase/cbbootstrap/design -o ./goa
goagen swagger -d github.com/couchbase/cbbootstrap/design -o ./goa