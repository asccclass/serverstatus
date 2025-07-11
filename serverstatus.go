package serverstatus
/*
   function: 系統檢康檢查用
*/

import(
   "fmt"
   "net/http"
)

type SrvStatus struct {
   SystemName	string
}

func NewServerStatus(systemName string)(*SrvStatus) {
   return  &SrvStatus{
      SystemName: systemName,
   }
}

func(app *SrvStatus) AddRouter(router *http.ServeMux) {
   router.Handle("GET /healthz", http.HandlerFunc(app.Healthz))
}

// check health
func(srv *SrvStatus) Healthz(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "{\"systemName\":" + srv.SystemName + ", \"status\": \"ok\"}")
}

