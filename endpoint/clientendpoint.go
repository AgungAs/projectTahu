package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.bluebird.id/mini/Agama/server"
)

func (ce AgamaEndpoint) AddAgamaService(ctx context.Context, agama sv.Agama) error {
	_, err := ce.AddAgamaEndpoint(ctx, agama)
	return err
}

func (ce AgamaEndpoint) ReadAgamaByKeteranganService(ctx context.Context, Keterangan string) (sv.Agamas, error) {
	req := sv.Agama{Keterangan: Keterangan}
	fmt.Println(req)
	resp, err := ce.ReadAgamaByKeteranganEndpoint(ctx, req)
	fmt.Println("service respon", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cuss := resp.(sv.Agamas)
	return cuss, err
}

func (ce AgamaEndpoint) ReadAgamaService(ctx context.Context) (sv.Agamas, error) {
	resp, err := ce.ReadAgamaEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Agamas), err
}

func (ce AgamaEndpoint) UpdateAgamaService(ctx context.Context, cus sv.Agama) error {
	_, err := ce.UpdateAgamaEndpoint(ctx, cus)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
