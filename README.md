# qr-code-svc

A simple REST service that takes a string (plain url encoded or base64) and returns it qr encoded as a png.

## Endpoints

### URL encoded string

#### Request

```
GET {host}/qr/str/{raw-string}
```

#### Response

| httpcode | content type     |
|:--------:|:-----------------|
| 200      | image/png        |
| 400      | application/json |

### Base64 encoded string

#### Request

```
GET {host}/qr/b64/{b64-string}
```

#### Response

| httpcode | content type     |
|:--------:|:-----------------|
| 200      | image/png        |
| 400      | application/json |
