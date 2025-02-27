## Server 狀態檢查程式

### API
```
/healthz
```

### How to use
```
import "github.com/asccclass/serverstatus"

// health check
systemName := os.Getenv("SystemName")
h := serverstatus.NewServerStatus(systemName)
router.HandleFunc("/healthz", h.Healthz).Methods("GET")
```
