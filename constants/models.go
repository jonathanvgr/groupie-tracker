package constants

// Structs basées sur les données de l'API

type Artist struct {
	Id uint
	Image string
	Name string
	Members []string
	CreationDate uint
	FirstAlbum string
	Locations string
	ConcertDates string
	Relations string
}

type Locations struct {
	Id uint
	Locations []string
	Dates string
}

type Dates struct {
	Id uint
	Dates []string
}

type Relations struct {
	Id uint
	DatesLocations map[string][]string
}

// Format des donnees a envoyer a une page web
type PageData struct {
	Data any
}