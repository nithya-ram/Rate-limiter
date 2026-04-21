

# Rate Limiter API

**5 req/min per user** 

## How to Run
1. Install Go dependencies:
   ```bash
   go mod tidy
   ```
2. Run the project:
   ```bash
   go run main.go
   ```

3. Server will run on:
   ```bash
   http://localhost:8080
## Test
```bash
# Request
curl -X POST http://localhost:8080/request -d '{"user_id":"user1","payload":"data"}'

# Stats  
curl http://localhost:8080/status
```

**Concurrent safe + Sliding window** 
