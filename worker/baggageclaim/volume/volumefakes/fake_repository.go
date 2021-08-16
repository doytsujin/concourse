// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	"context"
	"io"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim/volume"
)

type FakeRepository struct {
	CreateVolumeStub        func(context.Context, string, volume.Strategy, volume.Properties, bool) (volume.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 volume.Strategy
		arg4 volume.Properties
		arg5 bool
	}
	createVolumeReturns struct {
		result1 volume.Volume
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 error
	}
	DestroyVolumeStub        func(context.Context, string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	destroyVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVolumeAndDescendantsStub        func(context.Context, string) error
	destroyVolumeAndDescendantsMutex       sync.RWMutex
	destroyVolumeAndDescendantsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	destroyVolumeAndDescendantsReturns struct {
		result1 error
	}
	destroyVolumeAndDescendantsReturnsOnCall map[int]struct {
		result1 error
	}
	GetPrivilegedStub        func(context.Context, string) (bool, error)
	getPrivilegedMutex       sync.RWMutex
	getPrivilegedArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getPrivilegedReturns struct {
		result1 bool
		result2 error
	}
	getPrivilegedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetVolumeStub        func(context.Context, string) (volume.Volume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getVolumeReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	ListVolumesStub        func(context.Context, volume.Properties) (volume.Volumes, []string, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 context.Context
		arg2 volume.Properties
	}
	listVolumesReturns struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}
	SetPrivilegedStub        func(context.Context, string, bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(context.Context, string, string, string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, string, string, io.Reader) (bool, error)
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 io.Reader
	}
	streamInReturns struct {
		result1 bool
		result2 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	StreamOutStub        func(context.Context, string, string, string, io.Writer) error
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 io.Writer
	}
	streamOutReturns struct {
		result1 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 error
	}
	StreamP2pOutStub        func(context.Context, string, string, string, string) error
	streamP2pOutMutex       sync.RWMutex
	streamP2pOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	streamP2pOutReturns struct {
		result1 error
	}
	streamP2pOutReturnsOnCall map[int]struct {
		result1 error
	}
	VolumeParentStub        func(context.Context, string) (volume.Volume, bool, error)
	volumeParentMutex       sync.RWMutex
	volumeParentArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	volumeParentReturns struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	volumeParentReturnsOnCall map[int]struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) CreateVolume(arg1 context.Context, arg2 string, arg3 volume.Strategy, arg4 volume.Properties, arg5 bool) (volume.Volume, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 volume.Strategy
		arg4 volume.Properties
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("CreateVolume", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createVolumeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeRepository) CreateVolumeCalls(stub func(context.Context, string, volume.Strategy, volume.Properties, bool) (volume.Volume, error)) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *FakeRepository) CreateVolumeArgsForCall(i int) (context.Context, string, volume.Strategy, volume.Properties, bool) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRepository) CreateVolumeReturns(result1 volume.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 volume.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateVolumeReturnsOnCall(i int, result1 volume.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) DestroyVolume(arg1 context.Context, arg2 string) error {
	fake.destroyVolumeMutex.Lock()
	ret, specificReturn := fake.destroyVolumeReturnsOnCall[len(fake.destroyVolumeArgsForCall)]
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DestroyVolume", []interface{}{arg1, arg2})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyVolumeReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeRepository) DestroyVolumeCalls(stub func(context.Context, string) error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = stub
}

func (fake *FakeRepository) DestroyVolumeArgsForCall(i int) (context.Context, string) {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	argsForCall := fake.destroyVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) DestroyVolumeReturns(result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeReturnsOnCall(i int, result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	if fake.destroyVolumeReturnsOnCall == nil {
		fake.destroyVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeAndDescendants(arg1 context.Context, arg2 string) error {
	fake.destroyVolumeAndDescendantsMutex.Lock()
	ret, specificReturn := fake.destroyVolumeAndDescendantsReturnsOnCall[len(fake.destroyVolumeAndDescendantsArgsForCall)]
	fake.destroyVolumeAndDescendantsArgsForCall = append(fake.destroyVolumeAndDescendantsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DestroyVolumeAndDescendants", []interface{}{arg1, arg2})
	fake.destroyVolumeAndDescendantsMutex.Unlock()
	if fake.DestroyVolumeAndDescendantsStub != nil {
		return fake.DestroyVolumeAndDescendantsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyVolumeAndDescendantsReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsCallCount() int {
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	return len(fake.destroyVolumeAndDescendantsArgsForCall)
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsCalls(stub func(context.Context, string) error) {
	fake.destroyVolumeAndDescendantsMutex.Lock()
	defer fake.destroyVolumeAndDescendantsMutex.Unlock()
	fake.DestroyVolumeAndDescendantsStub = stub
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsArgsForCall(i int) (context.Context, string) {
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	argsForCall := fake.destroyVolumeAndDescendantsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsReturns(result1 error) {
	fake.destroyVolumeAndDescendantsMutex.Lock()
	defer fake.destroyVolumeAndDescendantsMutex.Unlock()
	fake.DestroyVolumeAndDescendantsStub = nil
	fake.destroyVolumeAndDescendantsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) DestroyVolumeAndDescendantsReturnsOnCall(i int, result1 error) {
	fake.destroyVolumeAndDescendantsMutex.Lock()
	defer fake.destroyVolumeAndDescendantsMutex.Unlock()
	fake.DestroyVolumeAndDescendantsStub = nil
	if fake.destroyVolumeAndDescendantsReturnsOnCall == nil {
		fake.destroyVolumeAndDescendantsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeAndDescendantsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) GetPrivileged(arg1 context.Context, arg2 string) (bool, error) {
	fake.getPrivilegedMutex.Lock()
	ret, specificReturn := fake.getPrivilegedReturnsOnCall[len(fake.getPrivilegedArgsForCall)]
	fake.getPrivilegedArgsForCall = append(fake.getPrivilegedArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetPrivileged", []interface{}{arg1, arg2})
	fake.getPrivilegedMutex.Unlock()
	if fake.GetPrivilegedStub != nil {
		return fake.GetPrivilegedStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPrivilegedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetPrivilegedCallCount() int {
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	return len(fake.getPrivilegedArgsForCall)
}

func (fake *FakeRepository) GetPrivilegedCalls(stub func(context.Context, string) (bool, error)) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = stub
}

func (fake *FakeRepository) GetPrivilegedArgsForCall(i int) (context.Context, string) {
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	argsForCall := fake.getPrivilegedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetPrivilegedReturns(result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	fake.getPrivilegedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetPrivilegedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	if fake.getPrivilegedReturnsOnCall == nil {
		fake.getPrivilegedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.getPrivilegedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetVolume(arg1 context.Context, arg2 string) (volume.Volume, bool, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetVolume", []interface{}{arg1, arg2})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getVolumeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRepository) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeRepository) GetVolumeCalls(stub func(context.Context, string) (volume.Volume, bool, error)) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = stub
}

func (fake *FakeRepository) GetVolumeArgsForCall(i int) (context.Context, string) {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	argsForCall := fake.getVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetVolumeReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) GetVolumeReturnsOnCall(i int, result1 volume.Volume, result2 bool, result3 error) {
	fake.getVolumeMutex.Lock()
	defer fake.getVolumeMutex.Unlock()
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 bool
			result3 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) ListVolumes(arg1 context.Context, arg2 volume.Properties) (volume.Volumes, []string, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 context.Context
		arg2 volume.Properties
	}{arg1, arg2})
	fake.recordInvocation("ListVolumes", []interface{}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.listVolumesReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRepository) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeRepository) ListVolumesCalls(stub func(context.Context, volume.Properties) (volume.Volumes, []string, error)) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = stub
}

func (fake *FakeRepository) ListVolumesArgsForCall(i int) (context.Context, volume.Properties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	argsForCall := fake.listVolumesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) ListVolumesReturns(result1 volume.Volumes, result2 []string, result3 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) ListVolumesReturnsOnCall(i int, result1 volume.Volumes, result2 []string, result3 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 volume.Volumes
			result2 []string
			result3 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 volume.Volumes
		result2 []string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) SetPrivileged(arg1 context.Context, arg2 string, arg3 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetPrivileged", []interface{}{arg1, arg2, arg3})
	fake.setPrivilegedMutex.Unlock()
	if fake.SetPrivilegedStub != nil {
		return fake.SetPrivilegedStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPrivilegedReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeRepository) SetPrivilegedCalls(stub func(context.Context, string, bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeRepository) SetPrivilegedArgsForCall(i int) (context.Context, string, bool) {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRepository) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetProperty(arg1 context.Context, arg2 string, arg3 string, arg4 string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetProperty", []interface{}{arg1, arg2, arg3, arg4})
	fake.setPropertyMutex.Unlock()
	if fake.SetPropertyStub != nil {
		return fake.SetPropertyStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setPropertyReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeRepository) SetPropertyCalls(stub func(context.Context, string, string, string) error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = stub
}

func (fake *FakeRepository) SetPropertyArgsForCall(i int) (context.Context, string, string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	argsForCall := fake.setPropertyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeRepository) SetPropertyReturns(result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamIn(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 io.Reader) (bool, error) {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 io.Reader
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.streamInReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeRepository) StreamInCalls(stub func(context.Context, string, string, string, io.Reader) (bool, error)) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeRepository) StreamInArgsForCall(i int) (context.Context, string, string, string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRepository) StreamInReturns(result1 bool, result2 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) StreamInReturnsOnCall(i int, result1 bool, result2 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) StreamOut(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 io.Writer) error {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 io.Writer
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("StreamOut", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamOutReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeRepository) StreamOutCalls(stub func(context.Context, string, string, string, io.Writer) error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeRepository) StreamOutArgsForCall(i int) (context.Context, string, string, string, io.Writer) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRepository) StreamOutReturns(result1 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamOutReturnsOnCall(i int, result1 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamP2pOut(arg1 context.Context, arg2 string, arg3 string, arg4 string, arg5 string) error {
	fake.streamP2pOutMutex.Lock()
	ret, specificReturn := fake.streamP2pOutReturnsOnCall[len(fake.streamP2pOutArgsForCall)]
	fake.streamP2pOutArgsForCall = append(fake.streamP2pOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("StreamP2pOut", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.streamP2pOutMutex.Unlock()
	if fake.StreamP2pOutStub != nil {
		return fake.StreamP2pOutStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.streamP2pOutReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) StreamP2pOutCallCount() int {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	return len(fake.streamP2pOutArgsForCall)
}

func (fake *FakeRepository) StreamP2pOutCalls(stub func(context.Context, string, string, string, string) error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = stub
}

func (fake *FakeRepository) StreamP2pOutArgsForCall(i int) (context.Context, string, string, string, string) {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	argsForCall := fake.streamP2pOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeRepository) StreamP2pOutReturns(result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	fake.streamP2pOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) StreamP2pOutReturnsOnCall(i int, result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	if fake.streamP2pOutReturnsOnCall == nil {
		fake.streamP2pOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamP2pOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) VolumeParent(arg1 context.Context, arg2 string) (volume.Volume, bool, error) {
	fake.volumeParentMutex.Lock()
	ret, specificReturn := fake.volumeParentReturnsOnCall[len(fake.volumeParentArgsForCall)]
	fake.volumeParentArgsForCall = append(fake.volumeParentArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("VolumeParent", []interface{}{arg1, arg2})
	fake.volumeParentMutex.Unlock()
	if fake.VolumeParentStub != nil {
		return fake.VolumeParentStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.volumeParentReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRepository) VolumeParentCallCount() int {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	return len(fake.volumeParentArgsForCall)
}

func (fake *FakeRepository) VolumeParentCalls(stub func(context.Context, string) (volume.Volume, bool, error)) {
	fake.volumeParentMutex.Lock()
	defer fake.volumeParentMutex.Unlock()
	fake.VolumeParentStub = stub
}

func (fake *FakeRepository) VolumeParentArgsForCall(i int) (context.Context, string) {
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	argsForCall := fake.volumeParentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) VolumeParentReturns(result1 volume.Volume, result2 bool, result3 error) {
	fake.volumeParentMutex.Lock()
	defer fake.volumeParentMutex.Unlock()
	fake.VolumeParentStub = nil
	fake.volumeParentReturns = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) VolumeParentReturnsOnCall(i int, result1 volume.Volume, result2 bool, result3 error) {
	fake.volumeParentMutex.Lock()
	defer fake.volumeParentMutex.Unlock()
	fake.VolumeParentStub = nil
	if fake.volumeParentReturnsOnCall == nil {
		fake.volumeParentReturnsOnCall = make(map[int]struct {
			result1 volume.Volume
			result2 bool
			result3 error
		})
	}
	fake.volumeParentReturnsOnCall[i] = struct {
		result1 volume.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	fake.destroyVolumeAndDescendantsMutex.RLock()
	defer fake.destroyVolumeAndDescendantsMutex.RUnlock()
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	fake.volumeParentMutex.RLock()
	defer fake.volumeParentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ volume.Repository = new(FakeRepository)
