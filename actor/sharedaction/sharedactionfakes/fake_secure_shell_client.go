// Code generated by counterfeiter. DO NOT EDIT.
package sharedactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/util/clissh"
)

type FakeSecureShellClient struct {
	ConnectStub        func(username string, passcode string, sshEndpoint string, sshHostKeyFingerprint string, skipHostValidation bool) error
	connectMutex       sync.RWMutex
	connectArgsForCall []struct {
		username              string
		passcode              string
		sshEndpoint           string
		sshHostKeyFingerprint string
		skipHostValidation    bool
	}
	connectReturns struct {
		result1 error
	}
	connectReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	InteractiveSessionStub        func(commands []string, terminalRequest clissh.TTYRequest) error
	interactiveSessionMutex       sync.RWMutex
	interactiveSessionArgsForCall []struct {
		commands        []string
		terminalRequest clissh.TTYRequest
	}
	interactiveSessionReturns struct {
		result1 error
	}
	interactiveSessionReturnsOnCall map[int]struct {
		result1 error
	}
	LocalPortForwardStub        func(localPortForwardSpecs []clissh.LocalPortForward) error
	localPortForwardMutex       sync.RWMutex
	localPortForwardArgsForCall []struct {
		localPortForwardSpecs []clissh.LocalPortForward
	}
	localPortForwardReturns struct {
		result1 error
	}
	localPortForwardReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecureShellClient) Connect(username string, passcode string, sshEndpoint string, sshHostKeyFingerprint string, skipHostValidation bool) error {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct {
		username              string
		passcode              string
		sshEndpoint           string
		sshHostKeyFingerprint string
		skipHostValidation    bool
	}{username, passcode, sshEndpoint, sshHostKeyFingerprint, skipHostValidation})
	fake.recordInvocation("Connect", []interface{}{username, passcode, sshEndpoint, sshHostKeyFingerprint, skipHostValidation})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub(username, passcode, sshEndpoint, sshHostKeyFingerprint, skipHostValidation)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.connectReturns.result1
}

func (fake *FakeSecureShellClient) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeSecureShellClient) ConnectArgsForCall(i int) (string, string, string, string, bool) {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return fake.connectArgsForCall[i].username, fake.connectArgsForCall[i].passcode, fake.connectArgsForCall[i].sshEndpoint, fake.connectArgsForCall[i].sshHostKeyFingerprint, fake.connectArgsForCall[i].skipHostValidation
}

func (fake *FakeSecureShellClient) ConnectReturns(result1 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) ConnectReturnsOnCall(i int, result1 error) {
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeSecureShellClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureShellClient) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) InteractiveSession(commands []string, terminalRequest clissh.TTYRequest) error {
	var commandsCopy []string
	if commands != nil {
		commandsCopy = make([]string, len(commands))
		copy(commandsCopy, commands)
	}
	fake.interactiveSessionMutex.Lock()
	ret, specificReturn := fake.interactiveSessionReturnsOnCall[len(fake.interactiveSessionArgsForCall)]
	fake.interactiveSessionArgsForCall = append(fake.interactiveSessionArgsForCall, struct {
		commands        []string
		terminalRequest clissh.TTYRequest
	}{commandsCopy, terminalRequest})
	fake.recordInvocation("InteractiveSession", []interface{}{commandsCopy, terminalRequest})
	fake.interactiveSessionMutex.Unlock()
	if fake.InteractiveSessionStub != nil {
		return fake.InteractiveSessionStub(commands, terminalRequest)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.interactiveSessionReturns.result1
}

func (fake *FakeSecureShellClient) InteractiveSessionCallCount() int {
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	return len(fake.interactiveSessionArgsForCall)
}

func (fake *FakeSecureShellClient) InteractiveSessionArgsForCall(i int) ([]string, clissh.TTYRequest) {
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	return fake.interactiveSessionArgsForCall[i].commands, fake.interactiveSessionArgsForCall[i].terminalRequest
}

func (fake *FakeSecureShellClient) InteractiveSessionReturns(result1 error) {
	fake.InteractiveSessionStub = nil
	fake.interactiveSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) InteractiveSessionReturnsOnCall(i int, result1 error) {
	fake.InteractiveSessionStub = nil
	if fake.interactiveSessionReturnsOnCall == nil {
		fake.interactiveSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.interactiveSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) LocalPortForward(localPortForwardSpecs []clissh.LocalPortForward) error {
	var localPortForwardSpecsCopy []clissh.LocalPortForward
	if localPortForwardSpecs != nil {
		localPortForwardSpecsCopy = make([]clissh.LocalPortForward, len(localPortForwardSpecs))
		copy(localPortForwardSpecsCopy, localPortForwardSpecs)
	}
	fake.localPortForwardMutex.Lock()
	ret, specificReturn := fake.localPortForwardReturnsOnCall[len(fake.localPortForwardArgsForCall)]
	fake.localPortForwardArgsForCall = append(fake.localPortForwardArgsForCall, struct {
		localPortForwardSpecs []clissh.LocalPortForward
	}{localPortForwardSpecsCopy})
	fake.recordInvocation("LocalPortForward", []interface{}{localPortForwardSpecsCopy})
	fake.localPortForwardMutex.Unlock()
	if fake.LocalPortForwardStub != nil {
		return fake.LocalPortForwardStub(localPortForwardSpecs)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.localPortForwardReturns.result1
}

func (fake *FakeSecureShellClient) LocalPortForwardCallCount() int {
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	return len(fake.localPortForwardArgsForCall)
}

func (fake *FakeSecureShellClient) LocalPortForwardArgsForCall(i int) []clissh.LocalPortForward {
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	return fake.localPortForwardArgsForCall[i].localPortForwardSpecs
}

func (fake *FakeSecureShellClient) LocalPortForwardReturns(result1 error) {
	fake.LocalPortForwardStub = nil
	fake.localPortForwardReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) LocalPortForwardReturnsOnCall(i int, result1 error) {
	fake.LocalPortForwardStub = nil
	if fake.localPortForwardReturnsOnCall == nil {
		fake.localPortForwardReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.localPortForwardReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.waitReturns.result1
}

func (fake *FakeSecureShellClient) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSecureShellClient) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) WaitReturnsOnCall(i int, result1 error) {
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecureShellClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.interactiveSessionMutex.RLock()
	defer fake.interactiveSessionMutex.RUnlock()
	fake.localPortForwardMutex.RLock()
	defer fake.localPortForwardMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecureShellClient) recordInvocation(key string, args []interface{}) {
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

var _ sharedaction.SecureShellClient = new(FakeSecureShellClient)
