//go:build !cl
// +build !cl

package nanopow

import "github.com/bbedward/go-opencl/opencl"

type noneGPUWorker struct{}

func NewWorkerGPU(device opencl.Device) (*noneGPUWorker, error) {
	return NewWorkerGPUThread(0, device)
}

func NewWorkerGPUThread(_ uint64, _ opencl.Device) (*noneGPUWorker, error) {
	return nil, ErrNotSupported
}

func (w *noneGPUWorker) GenerateWork(ctx *Context, root []byte, difficulty uint64) (err error) {
	return ErrNotSupported
}

func GetDevice() (dv opencl.Device, err error) {
	return opencl.Device{}, ErrNotSupported
}
