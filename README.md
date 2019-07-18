## Server 狀態檢查程式

### API
```
/healthz
```

### How to use
```
import "github.com/asccclass/serverstatus"

m := serverstatus.NewServerStatus()
router.HandleFunc("/healthz", m.Healthz).Methods("GET")
```
