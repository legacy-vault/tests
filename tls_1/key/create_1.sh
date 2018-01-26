#!/bin/bash

openssl req -config 1.cfg -new -x509 -sha256 -newkey rsa:4096 -nodes \
    -keyout 1/localhost.key -days 365 -out 1/localhost.crt