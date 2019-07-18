package ServerStatus

import(
   "net/http"
)

type SrvStatus {
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

