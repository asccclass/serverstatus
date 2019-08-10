package serverstatus
/*
   function: 系統檢康檢查用
*/

import(
   "net/http"
)

type SrvStatus struct {
}

func NewServerStatus()(*SrvStatus) {
   return  &SrvStatus{
   }
}

// check health
func(srv *SrvStatus) Healthz(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
}

