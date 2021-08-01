package facade

import (
	"go-chat/src/domain/service"
	"net/http"
)

func UpgradeToWSConnection(w http.ResponseWriter, r *http.Request) {
	service.UpgradeToWSConnection(w, r)
}
