package v3action_test

import (
	"errors"

	. "code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/actor/v3action/v3actionfakes"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Instance Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v3actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v3actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil, nil)
	})

	Describe("ShareServiceInstanceByOrganizationAndSpaceName", func() {
		var (
			serviceInstanceName string
			orgGUID             string
			spaceName           string

			warnings       Warnings
			executionError error
		)

		BeforeEach(func() {
			serviceInstanceName = "some-service-instance"
			orgGUID = "some-org-guid"
			spaceName = "some-space-name"
		})

		JustBeforeEach(func() {
			warnings, executionError = actor.ShareServiceInstanceByOrganizationAndSpaceName(serviceInstanceName, orgGUID, spaceName)
		})

		Context("when the service instance name is valid", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetServiceInstancesReturns([]ccv3.ServiceInstance{
					{
						Name: "some-service-instance",
						GUID: "some-service-instance-guid",
					},
				}, ccv3.Warnings{}, nil)
			})

			Context("when the space name is valid", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns([]ccv3.Space{
						{
							Name: "some-space",
							GUID: "some-space-guid",
						},
					}, ccv3.Warnings{}, nil)
				})

				Context("when the post request to the shared spaces endpoint succeeds", func() {
					BeforeEach(func() {
						fakeCloudControllerClient.PostServiceInstanceSharedSpacesReturns(ccv3.RelationshipList{}, ccv3.Warnings{}, nil)
					})

					It("calls to create a new service instance share", func() {
						Expect(fakeCloudControllerClient.PostServiceInstanceSharedSpacesCallCount()).To(Equal(1))
					})

					It("calls to create a new service instance share", func() {
						si_guid, space_guids := fakeCloudControllerClient.PostServiceInstanceSharedSpacesArgsForCall(0)
						Expect(si_guid).To(Equal("some-service-instance-guid"))
						Expect(space_guids).To(Equal([]string{"some-space-guid"}))
					})

					It("does not return warnings or errors", func() {
						Expect(executionError).ToNot(HaveOccurred())
						Expect(warnings).To(BeEmpty())
					})
				})

				Context("when the post request to the shared spaces endpoint fails", func() {
					err := errors.New("Share failed")
					BeforeEach(func() {
						fakeCloudControllerClient.PostServiceInstanceSharedSpacesReturns(ccv3.RelationshipList{}, ccv3.Warnings{}, err)
					})
					It("returns error", func() {
						Expect(executionError).To(Equal(err))
					})
				})
			})

			Context("when resolving the space name fails", func() {
				err := errors.New("Space name doesn't exist")

				BeforeEach(func() {
					fakeCloudControllerClient.GetSpacesReturns([]ccv3.Space{}, ccv3.Warnings{}, err)
				})

				It("returns error", func() {
					Expect(executionError).To(Equal(err))
				})
			})
		})

		Context("when resolving the service instance name fails", func() {
			err := errors.New("service name doesn't exist")

			BeforeEach(func() {
				fakeCloudControllerClient.GetServiceInstancesReturns([]ccv3.ServiceInstance{}, ccv3.Warnings{}, err)
			})

			It("returns error", func() {
				Expect(executionError).To(Equal(err))
			})
		})
	})

	Describe("GetServiceInstanceByName", func() {
		//todo
	})
})
