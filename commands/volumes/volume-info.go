package volumecommands

import (
	"net/http"

	"github.com/gluster/glusterd2/errors"
	"github.com/gluster/glusterd2/utils"
	"github.com/gluster/glusterd2/volume"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func volumeInfoHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	volname := p["volname"]

	log.Debug("In Volume info API")

	vol, e := volume.GetVolume(volname)
	if e != nil {
		utils.SendHTTPError(w, http.StatusNotFound, errors.ErrVolNotFound.Error())
	} else {

		utils.SendHTTPResponse(w, http.StatusOK, vol)
	}
}
