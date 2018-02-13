package abiquo

import . "github.com/abiquo/opal/core"

type Alarm struct {
	Active            bool   `json:"active"`
	EvaluationPeriods int    `json:"evaluationPeriods"`
	Formula           string `json:"formula"`
	Name              string `json:"name"`
	Period            int    `json:"period"`
	Statistic         string `json:"statistic"`
	Threshold         int    `json:"threshold"`
	DTO
}

type Alert struct {
	Active      bool    `json:"active"`
	Description string  `json:"description"`
	Muted       bool    `json:"muted"`
	Name        string  `json:"name"`
	Alarms      []*Link `json:"alarms"`
	DTO
}

type Device struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	DTO
}

type HardDisk struct {
	Id                 int    `json:"id,omitempty"`
	Label              string `json:"label,omitempty"`
	Bootable           bool   `json:"bootable,omitempty"`
	DiskController     string `json:"diskController,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	Sequence           int    `json:"sequence,omitempty"`
	DiskFormatType     string `json:"diskFormatType,omitempty"`
	DiskFileSize       int    `json:"diskFileSize,omitempty"`
	SizeInMb           int    `json:"sizeInMb,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	DTO
}

type Volume struct {
	Id                 int    `json:"idomitempty"`
	AllowResize        bool   `json:"allowResize"`
	SizeInMB           int    `json:"sizeInMB,omitempty"`
	State              string `json:"state,omitempty"`
	Name               string `json:"name,omitempty"`
	Bootable           bool   `json:"bootable,omitempty"`
	Description        string `json:"description,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty"`
	DiskController     string `json:"diskController"`
	DTO
}
