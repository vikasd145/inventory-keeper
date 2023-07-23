# inventory-keeper
- Certificate generated using opensll self signed certificate
```bash
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
- https server running on `8443` and http server running on port `80`
