package analytics_test

import (
    "net/http"
    "time"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"

    "github.com/andrew-bodine/informix/analytics"
    "github.com/andrew-bodine/informix/analytics/emit"
)

var _ = Describe("analytics", func() {
    Context("builtin", func() {
        var b analytics.Builtin

        BeforeEach(func() {
            b = analytics.NewBuiltin()
        })

        Context("Cache()", func() {
            Context("when not running", func() {
                It("returns an empty value for any key", func() {
                    data := b.Cache("")
                    Expect(data).NotTo(BeNil())
                    Expect(len(data)).To(Equal(0))
                })
            })

            Context("when running", func() {
                BeforeEach(func () {
                    b.Run(time.Microsecond)
                })

                AfterEach(func() {
                    b.Stop()
                })

                It("returns empty values for an invalid key", func() {
                    data := b.Cache("")
                    Expect(data).NotTo(BeNil())
                    Expect(len(data)).To(Equal(0))
                })

                It("returns the current value for a valid key", func() {
                    timer := time.NewTimer(time.Millisecond * 10)
                    <- timer.C

                    data := b.Cache(emit.MEMORY)
                    Expect(data).NotTo(BeNil())
                    Expect(len(data)).NotTo(Equal(0))

                    // Assert something about the data.
                    d := data[0].(map[string]int)
                    Expect(d["MemAvailable"]).NotTo(Equal(0))
                })
            })
        })

        Context("CacheHandler()", func() {
            Context("when not running", func() {
                It("returns empty values for all keys", func() {
                    r, err := http.NewRequest("GET", "/analytics/builtin", nil)
                    Expect(err).To(BeNil())

                    w := &MockResponseWriter{}
                    b.CacheHandler(w, r)

                    Expect(w.Buf).NotTo(Equal(""))
                })
            })

            Context("when running", func() {
                BeforeEach(func () {
                    b.Run(time.Microsecond)
                })

                AfterEach(func() {
                    b.Stop()
                })

                It("returns values for all keys", func() {
                    timer := time.NewTimer(time.Millisecond * 10)
                    <- timer.C

                    r, err := http.NewRequest("GET", "/analytics/builtin", nil)
                    Expect(err).To(BeNil())

                    w := &MockResponseWriter{}
                    b.CacheHandler(w, r)

                    Expect(w.Buf).NotTo(Equal(""))
                    Expect(w.Buf).To(ContainSubstring("MemAvailable"))
                })
            })
        })

        // Test the Builtin implementation.
        Context("Builtin", func() {
            Context("Run()", func() {
                AfterEach(func() {
                    b.Stop()
                })

                It("starts collecting builtin telemetry", func() {
                    b.Run(time.Microsecond)

                    timer := time.NewTimer(time.Millisecond * 10)
                    <- timer.C

                    data := b.Cache(emit.MEMORY)
                    Expect(data).NotTo(BeNil())
                    Expect(len(data)).NotTo(Equal(0))
                })

                Context("when already running", func() {
                    BeforeEach(func() {
                        b.Run(time.Microsecond)
                    })

                    It("doesn't do anything", func() {
                        timer := time.NewTimer(time.Millisecond)
                        <- timer.C

                        data := b.Cache(emit.MEMORY)
                        b.Run(time.Microsecond)
                        data2 := b.Cache(emit.MEMORY)
                        Expect(data).To(Equal(data2))
                    })
                })
            })

            Context("Stop()", func() {
                BeforeEach(func() {
                    b.Run(time.Microsecond)
                })

                It("stops collecting telemetry, and drains cache", func() {
                    timer := time.NewTimer(time.Millisecond)
                    <- timer.C

                    b.Stop()
                    data := b.Cache(emit.MEMORY)
                    Expect(data).NotTo(BeNil())
                    Expect(len(data)).To(Equal(0))
                })

                Context("when already stopped", func() {
                    // It("doesn't do anything", func() {})
                })
            })
        })
    })
})
