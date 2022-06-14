package automated_instance_backup

import (
	"encoding/json"
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"io/ioutil"
	"net/http"
	"os"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	orgaId := os.Getenv("ORGANIZATION_ID")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	client, err := scw.NewClient(
		scw.WithDefaultOrganizationID(orgaId),
		scw.WithAuth(accessKey, secretKey),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
		return
	}
	instanceApi := instance.NewAPI(client)

	type Body struct {
		Server string
		Zone   string
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
		return
	}

	var server Body

	err = json.Unmarshal(body, &server)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
		return
	}

	request := instance.ServerActionRequest{
		Zone:     scw.Zone(server.Zone),
		ServerID: server.Server,
		Action:   instance.ServerActionBackup,
	}
	_, err = instanceApi.ServerAction(&request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := fmt.Fprint(w, err.Error())
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string("done"))

}
