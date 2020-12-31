package registrierung

import (
	"log"
	"net/http"
	"strconv"
)

// RegistrierungsHandler is ..
type RegistrierungsHandler struct{}

//ServerHTTP is ....
func (rh *RegistrierungsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Printf("Could not parse form because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	var registrierung = &Registrierung{}

	registrierung.Firstname = req.Form.Get("Firstname")
	registrierung.Lastname = req.Form.Get("Lastname")
	registrierung.Email = req.Form.Get("Email")
	registrierung.Firma = req.Form.Get("Firma")
	registrierung.Firma = req.Form.Get("Firma")
	registrierung.Schulungscode = req.Form.Get("Schulungscode")
	registrierung.Datum = req.Form.Get("Datum")

	b, err := strconv.ParseBool(req.Form.Get("DatenschutzAkzeptiert"))

	if err != nil {
		log.Printf("Could not parse value of DatenschutzAkzeptiert because of %v", err)
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}

	registrierung.DatenschutzAkzeptiert = b
	log.Printf("new registration %v", registrierung)
	rw.WriteHeader(http.StatusCreated)
}
