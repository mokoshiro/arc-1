package client

import (
	"context"
	"encoding/json"

	"net/http"

	"bytes"

	"io/ioutil"

	"errors"

	"github.com/Bo0km4n/arc/internal/util"
	"github.com/Bo0km4n/arc/pkg/txm/executor"
	"github.com/Bo0km4n/arc/pkg/txm/executor/schema"
)

type ExecutorClient interface {
	StorePeer(ctx context.Context, req *executor.PreparePutPeerRequest) (*executor.PreparePutPeerResponse, error)
	UpdatePeerLocation(ctx context.Context, req *executor.UpdatePeerRequest) (*executor.UpdatePeerLocationResponse, error)
	DeletePeer(ctx context.Context, req *executor.DeletePeerRequest) (*executor.DeletePeerResponse, error)
	SelectPeer(ctx context.Context, peerID string) (*schema.Peer, error)
}

type executorClient struct {
	client *http.Client
	host   string
}

func NewExecutorClient(httpclient *http.Client, host string) ExecutorClient {
	return &executorClient{
		host:   host,
		client: httpclient,
	}
}

func (e *executorClient) StorePeer(ctx context.Context, req *executor.PreparePutPeerRequest) (*executor.PreparePutPeerResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	r, err := http.NewRequest("POST", e.host+"/api/peer", buf)
	if err != nil {
		return nil, err
	}
	httpresp, err := e.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()
	if !util.CheckHttpStatusCode(httpresp) {
		return nil, errors.New("Failed store peer on executor server")
	}

	resp := &executor.PreparePutPeerResponse{}
	respBody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *executorClient) UpdatePeerLocation(ctx context.Context, req *executor.UpdatePeerRequest) (*executor.UpdatePeerLocationResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	r, err := http.NewRequest("PUT", e.host+"/api/peer", buf)
	if err != nil {
		return nil, err
	}
	httpresp, err := e.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()
	if !util.CheckHttpStatusCode(httpresp) {
		return nil, errors.New("Failed store peer on executor server")
	}

	resp := &executor.UpdatePeerLocationResponse{}
	respBody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *executorClient) DeletePeer(ctx context.Context, req *executor.DeletePeerRequest) (*executor.DeletePeerResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	r, err := http.NewRequest("DELETE", e.host+"/api/peer", buf)
	if err != nil {
		return nil, err
	}
	httpresp, err := e.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()
	if !util.CheckHttpStatusCode(httpresp) {
		return nil, errors.New("Failed delete peer on executor server")
	}

	resp := &executor.DeletePeerResponse{}
	respBody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *executorClient) SelectPeer(ctx context.Context, peerID string) (*schema.Peer, error) {
	r, err := http.NewRequest("GET", e.host+"/api/peer?peer_id="+peerID, nil)
	if err != nil {
		return nil, err
	}
	httpresp, err := e.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer httpresp.Body.Close()
	if !util.CheckHttpStatusCode(httpresp) {
		return nil, errors.New("Failed delete peer on executor server")
	}

	resp := &schema.Peer{}
	respBody, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(respBody, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
