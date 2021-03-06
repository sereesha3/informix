package emit_test

import (
    linuxproc "github.com/c9s/goprocinfo/linux"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/andrew-bodine/informix/analytics/emit"
)

var _ = Describe("emit", func() {
    Context("Memory()", func() {
        It("creates a memory generator", func() {
            mem := emit.Memory()
            Expect(mem).NotTo(BeNil())
        })
    })

    Context("memory", func() {

        // Test the Generator implementation.
        Context("Generator", func() {
            Context("Generate()", func() {
                It("should return the current memory stats", func() {
                    mem := emit.Memory()
                    data := mem.Generate()
                    Expect(data).NotTo(BeNil())
                    info := data.(*linuxproc.MemInfo)
                    Expect(info.MemAvailable).NotTo(Equal(0))
                })

                Context("when there is an error", func() {})
            })
        })
    })
})
