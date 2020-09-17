#!/bin/bash

mkcert "*.tokopedia.com"

mkcert -install

mv _wildcard.tokopedia.com-key.pem tokopedia.com-key.pem
mv _wildcard.tokopedia.com.pem tokopedia.com.pem