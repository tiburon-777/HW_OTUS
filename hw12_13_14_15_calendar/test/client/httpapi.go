package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tiburon-777/HW_OTUS/hw12_13_14_15_calendar/pkg/api/public"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

type HTTPAPI struct{
	Name string
	BaseURL string
}

func (h HTTPAPI) GetName() string {
	return h.Name
}

func (h HTTPAPI) Create(req *public.CreateReq) (*public.CreateRsp, error) {
	jreq,err:= json.Marshal(req)
	if err != nil {
		return &public.CreateRsp{}, err
	}
	res, body, err := apiCall("POST",h.BaseURL+"/events", jreq)
	if err != nil {
		return &public.CreateRsp{}, err
	}
	if res.StatusCode!=201 {
		return &public.CreateRsp{}, fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	var createRsp public.CreateRsp
	err = json.Unmarshal(body,&createRsp)
	if err != nil {
		return &public.CreateRsp{}, err
	}
	return &createRsp, nil
}

func (h HTTPAPI) Update(req *public.UpdateReq) error {
	jreq, err:= json.Marshal(req)
	if err != nil {
		return err
	}
	res, _, err := apiCall("PUT",h.BaseURL+"/events/"+strconv.Itoa(int(req.ID)), jreq)
	if err != nil {
		return err
	}
	if res.StatusCode!=200 {
		return fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	return nil
}

func (h HTTPAPI) Delete(req *public.DeleteReq) error {
	jreq, err:= json.Marshal(req)
	if err != nil {
		return err
	}
	res, _, err := apiCall("DELETE",h.BaseURL+"/events/"+strconv.Itoa(int(req.ID)), jreq)
	if err != nil {
		return err
	}
	if res.StatusCode!=200 {
		return fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	return nil
}

func (h HTTPAPI) GetByID(req *public.GetByIDReq) ( *public.GetByIDResp, error) {
	jreq, err:= json.Marshal(req)
	if err != nil {
		return &public.GetByIDResp{}, err
	}
	res, body, err := apiCall("GET",h.BaseURL+"/events/"+strconv.Itoa(int(req.ID)), jreq)
	if err != nil {
		return &public.GetByIDResp{}, err
	}
	if res.StatusCode!=200 {
		return &public.GetByIDResp{}, fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	var getByIDResp public.GetByIDResp
	err = json.Unmarshal(body,&getByIDResp)
	if err != nil {
		return &public.GetByIDResp{}, err
	}
	return &getByIDResp, nil
}

func (h HTTPAPI) List() ( *public.ListResp, error) {
	res, body, err := apiCall("GET",h.BaseURL+"/events", nil)
	if err != nil {
		return &public.ListResp{}, err
	}
	if res.StatusCode!=200 {
		return &public.ListResp{}, fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	var listResp public.ListResp
	err = json.Unmarshal(body,&listResp)
	if err != nil {
		return &public.ListResp{}, err
	}
	return &listResp, nil
}

func (h HTTPAPI) GetByDate(req *public.GetByDateReq) ( *public.GetByDateResp, error) {
	jreq, err:= json.Marshal(req)
	if err != nil {
		return &public.GetByDateResp{}, err
	}
	res, body, err := apiCall("GET",h.BaseURL+"/events/"+string(req.Range)+"/"+req.Date.String(), jreq)
	if err != nil {
		return &public.GetByDateResp{}, err
	}
	if res.StatusCode!=200 {
		return &public.GetByDateResp{}, fmt.Errorf("unexpected status code %d",res.StatusCode)
	}
	var getByDateResp public.GetByDateResp
	err = json.Unmarshal(body,&getByDateResp)
	if err != nil {
		return &public.GetByDateResp{}, err
	}
	return &getByDateResp, nil
}

func apiCall(method string, url string, payload []byte) (*http.Response, []byte, error) {
	client := &http.Client{Transport: &http.Transport{ DialContext: (&net.Dialer{ Timeout: 15*time.Second}).DialContext }}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if err = res.Body.Close(); err != nil {
		return nil, nil, err
	}
	return res, body, nil
}
