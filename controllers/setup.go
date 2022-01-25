package controllers

import (

)

type interactor struct {
}

type Interactor interface {
	NewHelloControllerInstance() HelloController
}

func NewControllerInteractor() Interactor {
	return &interactor{}
}

func (interactor *interactor) NewHelloControllerInstance() HelloController {
	return NewHelloController()
}
