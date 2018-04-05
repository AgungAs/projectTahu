package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Agama.bluebird.id"
	OnAdd     Status = 1
)

type Agama struct {
	IDAgama   string
	NamaAgama string
	Status    int32
	CreateBy  string
	//CreatedOn string
	//UpdatedOn string
	UpdateBy   string
	Keterangan string
}
type Agamas []Agama

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddAgama(Agama) error
	ReadAgamaByKeterangan(string) (Agamas, error)
	ReadAgama() (Agamas, error)
	UpdateAgama(Agama) error
}

type AgamaService interface {
	AddAgamaService(context.Context, Agama) error
	ReadAgamaByKeteranganService(context.Context, string) (Agamas, error)
	ReadAgamaService(context.Context) (Agamas, error)
	UpdateAgamaService(context.Context, Agama) error
}
