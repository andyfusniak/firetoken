# Firetoken

A command line utility to sign-in to a Firebase project given an email and password. Intended to be used whilst developing applications that need to call Firebase.

The utlity will return a JWT and details about the claims for this user.

## Usage

### Grab a JWT
```
$ firetoken -w AIzaSyBS6cml-rrPpIrSgf6Mqc3ZeSI_B50sdxg
? Email: andy@andyfusniak.com
? Password: ********
Display Name: John Doe
Email: andy@andyfusniak.com

IDToken:
eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk3ZmNiY2EzNjhmZTc3ODA4ODMwYzgxMDAxMjFlYzdiZGUyMmNmMGUiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vdGVzdC1zcHljYW1lcmFjY3R2IiwibmFtZSI6IkpvaG4gRG9lIiwiY3V1aWQiOiJhODg1YTlhYi01YjIyLTRmZTQtODgyYi1jNTA4YjFhMDM0N2IiLCJyb2xlIjoiY3VzdG9tZXIiLCJhdWQiOiJ0ZXN0LXNweWNhbWVyYWNjdHYiLCJhdXRoX3RpbWUiOjE1NTAwMzU3OTIsInVzZXJfaWQiOiJhakxEWWF1ZHpmT0s4Vkk4N05vSG5MTWZWRkIyIiwic3ViIjoiYWpMRFlhdWR6Zk9LOFZJODdOb0huTE1mVkZCMiIsImlhdCI6MTU1MDAzNTc5MiwiZXhwIjoxNTUwMDM5MzkyLCJlbWFpbCI6ImFuZHlAYW5keWZ1c25pYWsuY29tIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7ImVtYWlsIjpbImFuZHlAYW5keWZ1c25pYWsuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoicGFzc3dvcmQifX0.M0Ve7-aq3lSN_yjGkD9vaTsUiFhwnmESL6ongb8b4ZkcsiHctsXZYivwYKltO2UdhHblCE0BKnXY5goFwNS98uuq0xtu0waTVdkgO396VfW1HQOVuExWa8VjwYTPu_z8sV2doKlkfaLdFJd0cXVYJ2EF0pa7FAgEL4Zr3y4khmFsT1Vl-dwgVDF1vaqTFntJxlxVvJiPIGxgdtu6CYwranyKwaYi6mzJYOtiiorYAte2GX_rUm4JBuTxVEkBnxi3ZXO4V2JeL6iKrmYJluC8T1aBcFeLx6dihTaWRTZwVUEtAcV8Jys7VNfUJPGJKMACm5ds4sZJ_U9Zjfy_cqnHJA

Claims:
CUUID: a885a9ab-5b22-4fe4-882b-c508b1a0347b Role: customer
```

### Show version
```
$ firetoken -v
v0.1.0
```

## Build

```
go build -ldflags "-X main.version=v0.1.0" -o firetoken
``
