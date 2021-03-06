// Code generated by counterfeiter. DO NOT EDIT.
package cfmysqlfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/plugin"
	sdkModels "code.cloudfoundry.org/cli/plugin/models"
	"github.com/andreasf/cf-mysql-plugin/cfmysql"
	pluginModels "github.com/andreasf/cf-mysql-plugin/cfmysql/models"
)

type FakeApiClient struct {
	GetServiceBindingsStub        func(cliConnection plugin.CliConnection) ([]pluginModels.ServiceBinding, error)
	getServiceBindingsMutex       sync.RWMutex
	getServiceBindingsArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getServiceBindingsReturns struct {
		result1 []pluginModels.ServiceBinding
		result2 error
	}
	getServiceBindingsReturnsOnCall map[int]struct {
		result1 []pluginModels.ServiceBinding
		result2 error
	}
	GetServiceInstancesStub        func(cliConnection plugin.CliConnection) ([]pluginModels.ServiceInstance, error)
	getServiceInstancesMutex       sync.RWMutex
	getServiceInstancesArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getServiceInstancesReturns struct {
		result1 []pluginModels.ServiceInstance
		result2 error
	}
	getServiceInstancesReturnsOnCall map[int]struct {
		result1 []pluginModels.ServiceInstance
		result2 error
	}
	GetStartedAppsStub        func(cliConnection plugin.CliConnection) ([]sdkModels.GetAppsModel, error)
	getStartedAppsMutex       sync.RWMutex
	getStartedAppsArgsForCall []struct {
		cliConnection plugin.CliConnection
	}
	getStartedAppsReturns struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}
	getStartedAppsReturnsOnCall map[int]struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApiClient) GetServiceBindings(cliConnection plugin.CliConnection) ([]pluginModels.ServiceBinding, error) {
	fake.getServiceBindingsMutex.Lock()
	ret, specificReturn := fake.getServiceBindingsReturnsOnCall[len(fake.getServiceBindingsArgsForCall)]
	fake.getServiceBindingsArgsForCall = append(fake.getServiceBindingsArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetServiceBindings", []interface{}{cliConnection})
	fake.getServiceBindingsMutex.Unlock()
	if fake.GetServiceBindingsStub != nil {
		return fake.GetServiceBindingsStub(cliConnection)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceBindingsReturns.result1, fake.getServiceBindingsReturns.result2
}

func (fake *FakeApiClient) GetServiceBindingsCallCount() int {
	fake.getServiceBindingsMutex.RLock()
	defer fake.getServiceBindingsMutex.RUnlock()
	return len(fake.getServiceBindingsArgsForCall)
}

func (fake *FakeApiClient) GetServiceBindingsArgsForCall(i int) plugin.CliConnection {
	fake.getServiceBindingsMutex.RLock()
	defer fake.getServiceBindingsMutex.RUnlock()
	return fake.getServiceBindingsArgsForCall[i].cliConnection
}

func (fake *FakeApiClient) GetServiceBindingsReturns(result1 []pluginModels.ServiceBinding, result2 error) {
	fake.GetServiceBindingsStub = nil
	fake.getServiceBindingsReturns = struct {
		result1 []pluginModels.ServiceBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetServiceBindingsReturnsOnCall(i int, result1 []pluginModels.ServiceBinding, result2 error) {
	fake.GetServiceBindingsStub = nil
	if fake.getServiceBindingsReturnsOnCall == nil {
		fake.getServiceBindingsReturnsOnCall = make(map[int]struct {
			result1 []pluginModels.ServiceBinding
			result2 error
		})
	}
	fake.getServiceBindingsReturnsOnCall[i] = struct {
		result1 []pluginModels.ServiceBinding
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetServiceInstances(cliConnection plugin.CliConnection) ([]pluginModels.ServiceInstance, error) {
	fake.getServiceInstancesMutex.Lock()
	ret, specificReturn := fake.getServiceInstancesReturnsOnCall[len(fake.getServiceInstancesArgsForCall)]
	fake.getServiceInstancesArgsForCall = append(fake.getServiceInstancesArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetServiceInstances", []interface{}{cliConnection})
	fake.getServiceInstancesMutex.Unlock()
	if fake.GetServiceInstancesStub != nil {
		return fake.GetServiceInstancesStub(cliConnection)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceInstancesReturns.result1, fake.getServiceInstancesReturns.result2
}

func (fake *FakeApiClient) GetServiceInstancesCallCount() int {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	return len(fake.getServiceInstancesArgsForCall)
}

func (fake *FakeApiClient) GetServiceInstancesArgsForCall(i int) plugin.CliConnection {
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	return fake.getServiceInstancesArgsForCall[i].cliConnection
}

func (fake *FakeApiClient) GetServiceInstancesReturns(result1 []pluginModels.ServiceInstance, result2 error) {
	fake.GetServiceInstancesStub = nil
	fake.getServiceInstancesReturns = struct {
		result1 []pluginModels.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetServiceInstancesReturnsOnCall(i int, result1 []pluginModels.ServiceInstance, result2 error) {
	fake.GetServiceInstancesStub = nil
	if fake.getServiceInstancesReturnsOnCall == nil {
		fake.getServiceInstancesReturnsOnCall = make(map[int]struct {
			result1 []pluginModels.ServiceInstance
			result2 error
		})
	}
	fake.getServiceInstancesReturnsOnCall[i] = struct {
		result1 []pluginModels.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetStartedApps(cliConnection plugin.CliConnection) ([]sdkModels.GetAppsModel, error) {
	fake.getStartedAppsMutex.Lock()
	ret, specificReturn := fake.getStartedAppsReturnsOnCall[len(fake.getStartedAppsArgsForCall)]
	fake.getStartedAppsArgsForCall = append(fake.getStartedAppsArgsForCall, struct {
		cliConnection plugin.CliConnection
	}{cliConnection})
	fake.recordInvocation("GetStartedApps", []interface{}{cliConnection})
	fake.getStartedAppsMutex.Unlock()
	if fake.GetStartedAppsStub != nil {
		return fake.GetStartedAppsStub(cliConnection)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStartedAppsReturns.result1, fake.getStartedAppsReturns.result2
}

func (fake *FakeApiClient) GetStartedAppsCallCount() int {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return len(fake.getStartedAppsArgsForCall)
}

func (fake *FakeApiClient) GetStartedAppsArgsForCall(i int) plugin.CliConnection {
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	return fake.getStartedAppsArgsForCall[i].cliConnection
}

func (fake *FakeApiClient) GetStartedAppsReturns(result1 []sdkModels.GetAppsModel, result2 error) {
	fake.GetStartedAppsStub = nil
	fake.getStartedAppsReturns = struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) GetStartedAppsReturnsOnCall(i int, result1 []sdkModels.GetAppsModel, result2 error) {
	fake.GetStartedAppsStub = nil
	if fake.getStartedAppsReturnsOnCall == nil {
		fake.getStartedAppsReturnsOnCall = make(map[int]struct {
			result1 []sdkModels.GetAppsModel
			result2 error
		})
	}
	fake.getStartedAppsReturnsOnCall[i] = struct {
		result1 []sdkModels.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeApiClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getServiceBindingsMutex.RLock()
	defer fake.getServiceBindingsMutex.RUnlock()
	fake.getServiceInstancesMutex.RLock()
	defer fake.getServiceInstancesMutex.RUnlock()
	fake.getStartedAppsMutex.RLock()
	defer fake.getStartedAppsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApiClient) recordInvocation(key string, args []interface{}) {
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

var _ cfmysql.ApiClient = new(FakeApiClient)
