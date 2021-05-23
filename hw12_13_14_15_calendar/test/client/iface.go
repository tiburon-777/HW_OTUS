package client

import "github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"

type Interface interface {
	GetName() string
	Create(req *public.CreateReq) (*public.CreateRsp, error)
	Update(req *public.UpdateReq) error
	Delete(req *public.DeleteReq) error
	GetByID(req *public.GetByIDReq) (*public.GetByIDResp, error)
	List() (*public.ListResp, error)
	GetByDate(req *public.GetByDateReq) (*public.GetByDateResp, error)
}
