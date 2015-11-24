package volumecommands

import (
	"net/http"

	"github.com/gluster/glusterd2/errors"
	"github.com/gluster/glusterd2/utils"
	"github.com/gluster/glusterd2/volume"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func volumeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	volname := p["volname"]

	log.Info("In Volume delete API")

	if volume.Exists(volname) {
		utils.SendHTTPError(w, http.StatusBadRequest, errors.ErrVolNotFound.Error())
		return
	}

	e := volume.DeleteVolume(volname)
	if e != nil {
		log.WithFields(log.Fields{"error": e.Error(),
			"volume": volname,
		}).Error("Failed to delete the volume")
		utils.SendHTTPError(w, http.StatusInternalServerError, e.Error())
		return
	}
	utils.SendHTTPResponse(w, http.StatusOK, nil)
}
