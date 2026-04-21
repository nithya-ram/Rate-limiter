

# Rate Limiter API

**5 req/min per user** 

## Run
```
go run main.go
```

## Test
```bash
# Request
curl -X POST http://localhost:8080/request -d '{"user_id":"user1","payload":"data"}'

# Stats  
curl http://localhost:8080/status
```

**Concurrent safe + Sliding window** 
