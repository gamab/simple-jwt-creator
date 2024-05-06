#!/bin/bash

# Generate EC key pair
echo "Generating EC key pair..."

echo "Private key: private.ec.key"
openssl ecparam -name prime256v1 -genkey -noout -out private.ec.key

echo "Public key: public.pem"
openssl ec -in private.ec.key -pubout -out public.pem
