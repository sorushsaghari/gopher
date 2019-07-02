package controllers

import "../services"

type ControllerConfig func(c *Controller) error

type Controller struct {
	Service services.Service
}

func NewController(cnfgs ...ControllerConfig) (*Controller, error) {
	var c Controller
	for _, cnfgs := range cnfgs {
		if err := cnfgs(&c); err != nil {
			return nil, err
		}
	}
	return &c, nil
}