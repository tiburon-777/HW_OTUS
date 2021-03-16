package client

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"time"
)

type GRPCAPI struct{
	Name string
	Ctx context.Context
	Host string
	Port string
}

func (h GRPCAPI) GetName() string {
	return h.Name
}

func (h GRPCAPI) Create(req *public.CreateReq) (*public.CreateRsp, error) {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return nil, err
	}
	resp, err := cliGRPC.Create(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func (h GRPCAPI) Update(req *public.UpdateReq) error {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return err
	}
	_, err = cliGRPC.Update(ctx, req)
	if err != nil{
		return err
	}
	return nil
}

func (h GRPCAPI) Delete(req *public.DeleteReq) error {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return err
	}
	_, err = cliGRPC.Delete(ctx, req)
	if err != nil{
		return err
	}
	return nil
}

func (h GRPCAPI) GetByID(req *public.GetByIDReq) (*public.GetByIDResp, error) {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return nil, err
	}
	resp, err := cliGRPC.GetByID(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func (h GRPCAPI) List() (*public.ListResp, error) {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return nil, err
	}
	resp, err := cliGRPC.List(ctx, &empty.Empty{})
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func (h GRPCAPI) GetByDate(req *public.GetByDateReq) (*public.GetByDateResp, error) {
	ctx, cliGRPC, err := getCli(h)
	if err != nil{
		return nil, err
	}
	resp, err := cliGRPC.GetByDate(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

func getCli(h GRPCAPI) (context.Context, public.GrpcClient, error) {
	ctx,_ := context.WithTimeout(h.Ctx, 15*time.Second)
	cliGRPC, err := public.NewClient(ctx,h.Host, h.Port)
	return ctx, cliGRPC, err
}