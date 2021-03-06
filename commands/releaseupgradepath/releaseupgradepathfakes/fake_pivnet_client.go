// This file was generated by counterfeiter
package releaseupgradepathfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-cli/commands/releaseupgradepath"
)

type FakePivnetClient struct {
	ReleasesForProductSlugStub        func(productSlug string) ([]go_pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		productSlug string
	}
	releasesForProductSlugReturns struct {
		result1 []go_pivnet.Release
		result2 error
	}
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	ReleaseUpgradePathsStub        func(productSlug string, releaseID int) ([]go_pivnet.ReleaseUpgradePath, error)
	releaseUpgradePathsMutex       sync.RWMutex
	releaseUpgradePathsArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseUpgradePathsReturns struct {
		result1 []go_pivnet.ReleaseUpgradePath
		result2 error
	}
	AddReleaseUpgradePathStub        func(productSlug string, releaseID int, previousReleaseID int) error
	addReleaseUpgradePathMutex       sync.RWMutex
	addReleaseUpgradePathArgsForCall []struct {
		productSlug       string
		releaseID         int
		previousReleaseID int
	}
	addReleaseUpgradePathReturns struct {
		result1 error
	}
	RemoveReleaseUpgradePathStub        func(productSlug string, releaseID int, previousReleaseID int) error
	removeReleaseUpgradePathMutex       sync.RWMutex
	removeReleaseUpgradePathArgsForCall []struct {
		productSlug       string
		releaseID         int
		previousReleaseID int
	}
	removeReleaseUpgradePathReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleasesForProductSlug(productSlug string) ([]go_pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{productSlug})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(productSlug)
	} else {
		return fake.releasesForProductSlugReturns.result1, fake.releasesForProductSlugReturns.result2
	}
}

func (fake *FakePivnetClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *FakePivnetClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.releasesForProductSlugArgsForCall[i].productSlug
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturns(result1 []go_pivnet.Release, result2 error) {
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 go_pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseUpgradePaths(productSlug string, releaseID int) ([]go_pivnet.ReleaseUpgradePath, error) {
	fake.releaseUpgradePathsMutex.Lock()
	fake.releaseUpgradePathsArgsForCall = append(fake.releaseUpgradePathsArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseUpgradePaths", []interface{}{productSlug, releaseID})
	fake.releaseUpgradePathsMutex.Unlock()
	if fake.ReleaseUpgradePathsStub != nil {
		return fake.ReleaseUpgradePathsStub(productSlug, releaseID)
	} else {
		return fake.releaseUpgradePathsReturns.result1, fake.releaseUpgradePathsReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseUpgradePathsCallCount() int {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return len(fake.releaseUpgradePathsArgsForCall)
}

func (fake *FakePivnetClient) ReleaseUpgradePathsArgsForCall(i int) (string, int) {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return fake.releaseUpgradePathsArgsForCall[i].productSlug, fake.releaseUpgradePathsArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseUpgradePathsReturns(result1 []go_pivnet.ReleaseUpgradePath, result2 error) {
	fake.ReleaseUpgradePathsStub = nil
	fake.releaseUpgradePathsReturns = struct {
		result1 []go_pivnet.ReleaseUpgradePath
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddReleaseUpgradePath(productSlug string, releaseID int, previousReleaseID int) error {
	fake.addReleaseUpgradePathMutex.Lock()
	fake.addReleaseUpgradePathArgsForCall = append(fake.addReleaseUpgradePathArgsForCall, struct {
		productSlug       string
		releaseID         int
		previousReleaseID int
	}{productSlug, releaseID, previousReleaseID})
	fake.recordInvocation("AddReleaseUpgradePath", []interface{}{productSlug, releaseID, previousReleaseID})
	fake.addReleaseUpgradePathMutex.Unlock()
	if fake.AddReleaseUpgradePathStub != nil {
		return fake.AddReleaseUpgradePathStub(productSlug, releaseID, previousReleaseID)
	} else {
		return fake.addReleaseUpgradePathReturns.result1
	}
}

func (fake *FakePivnetClient) AddReleaseUpgradePathCallCount() int {
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	return len(fake.addReleaseUpgradePathArgsForCall)
}

func (fake *FakePivnetClient) AddReleaseUpgradePathArgsForCall(i int) (string, int, int) {
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	return fake.addReleaseUpgradePathArgsForCall[i].productSlug, fake.addReleaseUpgradePathArgsForCall[i].releaseID, fake.addReleaseUpgradePathArgsForCall[i].previousReleaseID
}

func (fake *FakePivnetClient) AddReleaseUpgradePathReturns(result1 error) {
	fake.AddReleaseUpgradePathStub = nil
	fake.addReleaseUpgradePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveReleaseUpgradePath(productSlug string, releaseID int, previousReleaseID int) error {
	fake.removeReleaseUpgradePathMutex.Lock()
	fake.removeReleaseUpgradePathArgsForCall = append(fake.removeReleaseUpgradePathArgsForCall, struct {
		productSlug       string
		releaseID         int
		previousReleaseID int
	}{productSlug, releaseID, previousReleaseID})
	fake.recordInvocation("RemoveReleaseUpgradePath", []interface{}{productSlug, releaseID, previousReleaseID})
	fake.removeReleaseUpgradePathMutex.Unlock()
	if fake.RemoveReleaseUpgradePathStub != nil {
		return fake.RemoveReleaseUpgradePathStub(productSlug, releaseID, previousReleaseID)
	} else {
		return fake.removeReleaseUpgradePathReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveReleaseUpgradePathCallCount() int {
	fake.removeReleaseUpgradePathMutex.RLock()
	defer fake.removeReleaseUpgradePathMutex.RUnlock()
	return len(fake.removeReleaseUpgradePathArgsForCall)
}

func (fake *FakePivnetClient) RemoveReleaseUpgradePathArgsForCall(i int) (string, int, int) {
	fake.removeReleaseUpgradePathMutex.RLock()
	defer fake.removeReleaseUpgradePathMutex.RUnlock()
	return fake.removeReleaseUpgradePathArgsForCall[i].productSlug, fake.removeReleaseUpgradePathArgsForCall[i].releaseID, fake.removeReleaseUpgradePathArgsForCall[i].previousReleaseID
}

func (fake *FakePivnetClient) RemoveReleaseUpgradePathReturns(result1 error) {
	fake.RemoveReleaseUpgradePathStub = nil
	fake.removeReleaseUpgradePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	fake.addReleaseUpgradePathMutex.RLock()
	defer fake.addReleaseUpgradePathMutex.RUnlock()
	fake.removeReleaseUpgradePathMutex.RLock()
	defer fake.removeReleaseUpgradePathMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ releaseupgradepath.PivnetClient = new(FakePivnetClient)
