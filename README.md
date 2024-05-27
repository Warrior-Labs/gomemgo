# GOMEMGO

This is the client library for the Memgo in-memory story.

## Installing

```bash
go get github.com/warrior-labs/gomemgo
```

## Connecting

### Insecure Connection (ideal for local development)

```go
client, err := gomemgo.NewClient("localhost:8000", nil)
if err != nil {
  log.Fatalln(err)
}
```

### Secure Connection (ideal for deployment)
```go
client, err := gomemgo.NewClient("mysecurehost.com:8000",
&gomemgo.MemgoSecureOptions{
  X509CertificatePath: "path/to/cert.pem",
  X509KeyPath          "path/to.key.pem",
})
if err != nil {
  log.Fatalln(err)
}
```
