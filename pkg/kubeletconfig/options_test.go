package kubeletconfig

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var _ = Describe("KubeletConfigOptions", func() {

	It("Add all flags to command", func() {
		cmd := &cobra.Command{}
		flags := cmd.Flags()
		Expect(flags).NotTo(BeNil())
		Expect(flags.Lookup(PodPidsLimitOption)).To(BeNil())
		Expect(flags.Lookup(NameOption)).To(BeNil())

		options := NewKubeletConfigOptions()
		options.AddAllFlags(cmd)

		flag := flags.Lookup(PodPidsLimitOption)
		assertFlag(flag, PodPidsLimitOption, PodPidsLimitOptionUsage)

		flag = flags.Lookup(NameOption)
		assertFlag(flag, NameOption, NameOptionUsage)
	})

	It("Adds name flag to command", func() {
		cmd := &cobra.Command{}
		flags := cmd.Flags()
		Expect(flags).NotTo(BeNil())
		Expect(flags.Lookup(NameOption)).To(BeNil())

		options := NewKubeletConfigOptions()
		options.AddNameFlag(cmd)

		flag := flags.Lookup(PodPidsLimitOption)
		Expect(flag).To(BeNil())

		flag = flags.Lookup(NameOption)
		assertFlag(flag, NameOption, NameOptionUsage)
	})

	It("Fails HCP validation if no name supplied", func() {
		options := NewKubeletConfigOptions()
		err := options.ValidateForHypershift()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("The --name flag is required for Hosted Control Plane clusters."))
	})

	It("Passes HCP validation if the name is supplied", func() {
		options := NewKubeletConfigOptions()
		options.Name = "foo"
		err := options.ValidateForHypershift()
		Expect(err).NotTo(HaveOccurred())
	})

	It("Binds name from args if name not set by flag", func() {
		options := NewKubeletConfigOptions()
		options.BindFromArgs([]string{"bob"})
		Expect(options.Name).To(Equal("bob"))
	})

	It("Does not bind name from args if set by --name flag", func() {
		options := NewKubeletConfigOptions()
		options.Name = "foo"
		options.BindFromArgs([]string{"bob"})
		Expect(options.Name).To(Equal("foo"))
	})
})

func assertFlag(flag *flag.Flag, name string, usage string) {
	Expect(flag).NotTo(BeNil())
	Expect(flag.Name).To(Equal(name))
	Expect(flag.Usage).To(Equal(usage))
}
