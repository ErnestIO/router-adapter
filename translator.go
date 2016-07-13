package main

import (
	"encoding/json"
)

type builderEvent struct {
	Uuid               string `json:"_uuid"`
	BatchID            string `json:"_batch_id"`
	Service            string `json:"service"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	ClientName         string `json:"client_name"`
	DatacenterName     string `json:"datacenter_name"`
	DatacenterPassword string `json:"datacenter_password"`
	DatacenterRegion   string `json:"datacenter_region"`
	DatacenterType     string `json:"datacenter_type"`
	DatacenterUsername string `json:"datacenter_username"`
	ExternalNetwork    string `json:"external_network"`
	VCloudURL          string `json:"vcloud_url"`
	VseURL             string `json:"vse_url"`
	IP                 string `json:"ip"`
	Created            bool   `json:"created"`
	Status             string `json:"status"`
	ErrorCode          string `json:"error_code"`
	ErrorMessage       string `json:"error_message"`
}

type vcloudEvent struct {
	Uuid               string `json:"_uuid"`
	BatchID            string `json:"_batch_id"`
	Service            string `json:"service_id"`
	Type               string `json:"type"`
	RouterName         string `json:"router_name"`
	RouterType         string `json:"router_type"`
	RouterIP           string `json:"router_ip"`
	ClientName         string `json:"client_name"`
	DatacenterName     string `json:"datacenter_name"`
	DatacenterUsername string `json:"datacenter_username"`
	DatacenterPassword string `json:"datacenter_password"`
	DatacenterRegion   string `json:"datacenter_region"`
	DatacenterType     string `json:"datacenter_type"`
	ExternalNetwork    string `json:"external_network"`
	VCloudURL          string `json:"vcloud_url"`
	VseURL             string `json:"vse_url"`
	Status             string `json:"status"`
}

type Translator struct{}

func (t Translator) BuilderToConnector(j []byte) []byte {
	var input builderEvent
	var output vcloudEvent

	json.Unmarshal(j, &input)

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.RouterName = input.Name
	output.Service = input.Service
	output.RouterType = input.Type
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterType = input.DatacenterType
	output.ExternalNetwork = input.ExternalNetwork
	output.VCloudURL = input.VCloudURL
	output.VseURL = input.VseURL
	output.Status = input.Status

	body, _ := json.Marshal(output)
	return body
}

func (t Translator) ConnectorToBuilder(j []byte) []byte {
	var input vcloudEvent
	var output builderEvent

	json.Unmarshal(j, &input)

	output.Uuid = input.Uuid
	output.BatchID = input.BatchID
	output.Name = input.RouterName
	output.Service = input.Service
	output.Type = input.RouterType
	output.IP = input.RouterIP
	output.ClientName = input.ClientName
	output.DatacenterName = input.DatacenterName
	output.DatacenterUsername = input.DatacenterUsername
	output.DatacenterPassword = input.DatacenterPassword
	output.DatacenterRegion = input.DatacenterRegion
	output.DatacenterType = input.DatacenterType
	output.ExternalNetwork = input.ExternalNetwork
	output.VCloudURL = input.VCloudURL
	output.VseURL = input.VseURL
	output.Status = input.Status

	body, _ := json.Marshal(output)
	return body
}
