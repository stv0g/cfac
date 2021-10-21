package cccac

type ResponseCurrentStatus struct {
	Changed Status `json:"changed"`
}

type Status struct {
	Status string `json:"status"`
	Time   uint   `json:"time"`
	Type   string `json:"type"`
}
