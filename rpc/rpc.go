package rpc

import (
	"net"
	"strings"
)

type reader[T net.Conn] func(T, []byte)
type readWriter[T net.Conn] func(T, func([]byte) []byte)
type procedure[T net.Conn] func([]string) []string

type rpc[T net.Conn] struct {
	procedures map[string]procedure[T]
}

func NewRpc[T net.Conn]() rpc[T] {
	return rpc[T]{
		procedures: map[string]procedure[T]{},
	}
}

func (r *rpc[T]) RegisterProcedure(command string, p procedure[T]) {
	r.procedures[command] = p
}

func (r *rpc[T]) UnregisterProcedure(command string) {
	delete(r.procedures, command)
}

func (r *rpc[T]) responder(request []byte) []byte {
	fp := strings.Split(string(request), " ")
	procedure, payload := r.procedures[fp[0]], fp[1:]
	response := procedure(payload)
	return []byte(strings.Join(response, " "))
}

func (r *rpc[T]) Start(conn T, readWrite readWriter[T], maxBytes int) {
	for {
		readWrite(conn, r.responder)
	}
}
