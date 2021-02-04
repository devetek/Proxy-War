#!/bin/bash

mkcert "*.domain.com"

mkcert -install

mv _wildcard.domain.com-key.pem domain.com-key.pem
mv _wildcard.domain.com.pem domain.com.pem