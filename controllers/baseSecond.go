package controllers

import (
	"kereta/usecase"
	"net/http"
)

type BaseControllerSecond struct {
	us usecase.UsecaseInterfaceSecond
}

type ControllerSecond interface {
	GetData(w http.ResponseWriter, r *http.Request)
	UpdateData(w http.ResponseWriter, r *http.Request)
	DeleteDataDetailKereta(w http.ResponseWriter, r *http.Request)
}

func NewControllerSecond(us usecase.UsecaseInterfaceSecond) ControllerSecond {
	return &BaseControllerSecond{us: us}
}
