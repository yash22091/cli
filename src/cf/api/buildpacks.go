package api

import (
	"bytes"
	"cf/configuration"
	"cf/errors"
	"cf/models"
	"cf/net"
	"encoding/json"
	"fmt"
	"net/url"
)

type BuildpackRepository interface {
	FindByName(name string) (buildpack models.Buildpack, apiResponse errors.Error)
	ListBuildpacks(func(models.Buildpack) bool) errors.Error
	Create(name string, position *int, enabled *bool, locked *bool) (createdBuildpack models.Buildpack, apiResponse errors.Error)
	Delete(buildpackGuid string) (apiResponse errors.Error)
	Update(buildpack models.Buildpack) (updatedBuildpack models.Buildpack, apiResponse errors.Error)
}

type CloudControllerBuildpackRepository struct {
	config  configuration.Reader
	gateway net.Gateway
}

func NewCloudControllerBuildpackRepository(config configuration.Reader, gateway net.Gateway) (repo CloudControllerBuildpackRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerBuildpackRepository) ListBuildpacks(cb func(models.Buildpack) bool) errors.Error {
	return repo.gateway.ListPaginatedResources(
		repo.config.ApiEndpoint(),
		repo.config.AccessToken(),
		buildpacks_path,
		BuildpackResource{},
		func(resource interface{}) bool {
			return cb(resource.(BuildpackResource).ToFields())
		})
}

func (repo CloudControllerBuildpackRepository) FindByName(name string) (buildpack models.Buildpack, apiResponse errors.Error) {
	foundIt := false
	apiResponse = repo.gateway.ListPaginatedResources(
		repo.config.ApiEndpoint(),
		repo.config.AccessToken(),
		fmt.Sprintf("%s?q=%s", buildpacks_path, url.QueryEscape("name:"+name)),
		BuildpackResource{},
		func(resource interface{}) bool {
			buildpack = resource.(BuildpackResource).ToFields()
			foundIt = true
			return false
		})

	if !foundIt {
		apiResponse = errors.NewNotFoundError("%s %s not found", "Buildpack", name)
	}
	return
}

func (repo CloudControllerBuildpackRepository) Create(name string, position *int, enabled *bool, locked *bool) (createdBuildpack models.Buildpack, apiResponse errors.Error) {
	path := repo.config.ApiEndpoint() + buildpacks_path
	entity := BuildpackEntity{Name: name, Position: position, Enabled: enabled, Locked: locked}
	body, err := json.Marshal(entity)
	if err != nil {
		apiResponse = errors.NewErrorWithError("Could not serialize information", err)
		return
	}

	resource := new(BuildpackResource)
	apiResponse = repo.gateway.CreateResourceForResponse(path, repo.config.AccessToken(), bytes.NewReader(body), resource)
	if apiResponse != nil {
		return
	}

	createdBuildpack = resource.ToFields()
	return
}

func (repo CloudControllerBuildpackRepository) Delete(buildpackGuid string) (apiResponse errors.Error) {
	path := fmt.Sprintf("%s%s/%s", repo.config.ApiEndpoint(), buildpacks_path, buildpackGuid)
	apiResponse = repo.gateway.DeleteResource(path, repo.config.AccessToken())
	return
}

func (repo CloudControllerBuildpackRepository) Update(buildpack models.Buildpack) (updatedBuildpack models.Buildpack, apiResponse errors.Error) {
	path := fmt.Sprintf("%s%s/%s", repo.config.ApiEndpoint(), buildpacks_path, buildpack.Guid)

	entity := BuildpackEntity{buildpack.Name, buildpack.Position, buildpack.Enabled, "", "", buildpack.Locked}

	body, err := json.Marshal(entity)
	if err != nil {
		apiResponse = errors.NewErrorWithError("Could not serialize updates.", err)
		return
	}

	resource := new(BuildpackResource)
	apiResponse = repo.gateway.UpdateResourceForResponse(path, repo.config.AccessToken(), bytes.NewReader(body), resource)
	if apiResponse != nil {
		return
	}

	updatedBuildpack = resource.ToFields()
	return
}

const buildpacks_path = "/v2/buildpacks"

func (resource BuildpackResource) ToFields() models.Buildpack {
	return models.Buildpack{
		Guid:     resource.Metadata.Guid,
		Name:     resource.Entity.Name,
		Position: resource.Entity.Position,
		Enabled:  resource.Entity.Enabled,
		Key:      resource.Entity.Key,
		Filename: resource.Entity.Filename,
		Locked:   resource.Entity.Locked,
	}
}

type BuildpackResource struct {
	Metadata Metadata
	Entity   BuildpackEntity
}

type BuildpackEntity struct {
	Name     string `json:"name"`
	Position *int   `json:"position,omitempty"`
	Enabled  *bool  `json:"enabled,omitempty"`
	Key      string `json:"key,omitempty"`
	Filename string `json:"filename,omitempty"`
	Locked   *bool  `json:"locked,omitempty"`
}
