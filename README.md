# inventory-keeper
- Certificate generated using opensll self signed certificate
```bash
openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
- https server running on `8443` and http server running on port `80`
- unit test case written as an example in `internal/appliance/model_test.go`

## Possible Improvements
1. Test coverage should be 80%
2. The front end must make calls via HTTPS on port 8443, which should only be accessed from the public for specific HTTP calls. The HTTP port, on the other hand, should remain exposed only for internal testing purposes and should not be accessible to external traffic.
3. certificate should not be commited on github they should be stored in k8 secret. and certificate should be taken from trusted sites self signed certificate is not trusted by google. But communication is encrypted.
4. Storing and fetching value should be transactional