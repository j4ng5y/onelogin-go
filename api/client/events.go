package client

type GetEventTypesRequest struct {}
type GetEventTypesResponse struct {}
type GetEventsRequest struct {}
type GetEventsResponse struct {}
type GetEventRequest struct {}
type GetEventResponse struct {}
type CreateEventRequest struct {}
type CreateEventResponse struct {}

func (C *Client) GetEventTypes(req *GetEventTypesRequest) (*GetEventTypesResponse, error) {}

func (C *Client) GetEvents(req *GetEventsRequest) (*GetEventsResponse, error) {}

func (C *Client) GetEvent(req *GetEventRequest) (*GetEventResponse, error) {}

func (C *Client) CreateEvent(req *CreateEventRequest) (*CreateEventResponse, error) {}