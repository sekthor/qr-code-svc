# qr-code-svc

A simple REST service that takes a string (plain url encoded or base64) and returns it qr encoded as a png.

## Endpoints

```http
GET {host}/qr/str/{raw-string}
GET {host}/qr/b64/{b64-string}
```