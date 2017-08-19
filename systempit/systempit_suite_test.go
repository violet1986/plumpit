package systempit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSystempit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Systempit Suite")
}
