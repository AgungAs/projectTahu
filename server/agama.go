package server

import (
	"context"
)

type agama struct {
	writer ReadWriter
}

func NewAgama(writer ReadWriter) AgamaService {
	return &agama{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (c *agama) AddAgamaService(ctx context.Context, agama Agama) error {
	//fmt.Println("customer")
	err := c.writer.AddAgama(agama)
	if err != nil {
		return err
	}

	return nil
}

func (c *agama) ReadAgamaByKeteranganService(ctx context.Context, mob string) (Agamas, error) {
	cus, err := c.writer.ReadAgamaByKeterangan(mob)
	//fmt.Println(cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *agama) ReadAgamaService(ctx context.Context) (Agamas, error) {
	cus, err := c.writer.ReadAgama()
	//fmt.Println("customer", cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (c *agama) UpdateAgamaService(ctx context.Context, cus Agama) error {
	err := c.writer.UpdateAgama(cus)
	if err != nil {
		return err
	}
	return nil
}
