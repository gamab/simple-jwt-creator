# Simple JWT creator

This is a simple program to create a JWT token out of a json payload.

## Generate an ECDSA key pair

This program signs the token with an EC private key.
To generate one, you can run the following command:

```bash
bash gen-ec-key-pair.sh
```

## Create a JWT token

First you'll need to create the payload you want the JWT to embed.
For example:

```json
{
  "sub": "1234567890",
  "name": "John Doe",
  "iat": 1516239022,
  "aud": "myapp"
}
```

Then you can launch the program to sign it using an EC private key:

```bash
go run main.go -private-ec-key-file=./private.ec.key -assertion-file=payload.json
```

The program will return a JWT token.
For example:

```
eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJteWFwcCIsImlhdCI6MTUxNjIzOTAyMiwibmFtZSI6IkpvaG4gRG9lIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.BDuM5unfN-dBCXSkxxR97CaSXocLKaN3hdAJXktMTwepptLXGH6CycVv1exMjaQmPp1n3cHa48vp40-bIMzr_w
```

With your payload and the following header:

```json
{
  "alg": "ES256",
  "typ": "jwt"
}
```
