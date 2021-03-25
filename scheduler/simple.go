package main

import (
	"example.com/go-http-demo/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan (c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit (r engine.Request) {
	// 不希望卡死,添加goruntine来添加
	go func() { s.workerChan <- r }()
}