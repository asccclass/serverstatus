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

// check health
func(srv *SrvStatus) Healthz(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "{\"status\": \"ok\"}")
}

