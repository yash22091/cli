package v3action_test

import (
	. "code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/actor/v3action/v3actionfakes"

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

		Context("something", func() {
			// before each
			It("works", func() {
				Expect(executionError).ToNot(HaveOccurred())
				Expect(warnings).To(BeEmpty())
			})
		})

		Context("!something", func() {
			It("returns error and warnings", func() {
				Expect(executionError).ToNot(HaveOccurred())
				Expect(warnings).To(BeEmpty())
			})
		})
	})
})
