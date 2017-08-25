package querypit_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestQuerypit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Querypit Suite")
}
