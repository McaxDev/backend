package main

import "context"

var Servers map[string]Server

type Server interface {
	GetInfo() *SrvInfo
	SendCmd(ctx context.Context, cmd string) (string, error)
}

type SrvInfo struct {
	ID     string
	Name   string
	Path   string
	Backup struct {
		Enable    bool
		Frequency string
		Limit     int
		WorldPath string
	}
}
