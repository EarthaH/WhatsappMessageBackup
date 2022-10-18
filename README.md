# WhatsAppAPI
Golang API to parse whatsapp message jsons to a DB

## Requirements
- go 1.19+
Get started with Go [Here](https://go.dev/learn/)

## Build and Run
Build the main.go file into an executable called `main`
```
go build cmd/main/main.go
```
Now you can run `main` which will make the API available at `http://localhost:8080`
```
./main
```

## Usage

### Check Health `/health/`
This endpoint can be used to confirm the service is running.
```
curl http://localhost:8080/health/
```

### Save New Message `/message/`
Use this endpoint to POST new messages to the service in JSON form. The messages will then be parsed from JSON and saved into a DB for later use.
Returns HTTP 201 on success.
```
curl -X POST http://localhost:8080/message/ -H 'Content-Type: application/json' -d '[ { ... : ... } ]'
```
