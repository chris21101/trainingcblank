package registrierung

//Registrierung ist das zentrale Domänenobjekt des Registrierungsservice
type Registrierung struct {
	ID                    string
	Firstname             string
	Lastname              string
	Email                 string
	Firma                 string
	Schulungscode         string
	Datum                 string
	DatenschutzAkzeptiert bool
	Confirmed             bool
	//Adresse, etc nicht dargestellt
}
