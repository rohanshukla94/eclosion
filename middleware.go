package eclosion

import "net/http"

func (ecl *Eclosion) SessionLoad(next http.Handler) http.Handler {

	ecl.InfoLog.Print("SessionLoad called")
	return ecl.Session.LoadAndSave(next)
}
