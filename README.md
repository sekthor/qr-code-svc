# qr-code-svc

A simple REST service that takes a string (plain url encoded or base64) and returns it qr encoded as a png.

## Endpoints

```
GET {host}/qr/str/{raw-string}
```

| httpcode | content type     |
|:--------:|:-----------------|
| 200      | image/png        |
| 400      | application/json |

GET {host}/qr/b64/{b64-string}