package imgtobase64

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ImgTestSuite struct {
	suite.Suite
}

func TestImgTestSuite(t *testing.T) {
	suite.Run(t, new(ImgTestSuite))
}

func (suite *ImgTestSuite) TestImgtobase64() {
	suite.T().Run("should succesfully generate base64 from source file", func(t *testing.T) {
		expect, err := ioutil.ReadFile("test/fixtures/base64_sample.txt")

		require.NoError(t, err)

		img, err := Imgtobase64("test/fixtures/test.png")

		require.NoError(t, err)

		if img != string(expect) {
			t.Error("Imgtobase64 mismatch from expected output")
		}
	})

	suite.T().Run("should return error if image doesn't exist", func(t *testing.T) {
		img, err := Imgtobase64("doesntexist.png")

		if img != "" {
			t.Error("Imgtobase64 is't return empty line in error case")
		}

		if err == nil {
			t.Error("Imgtobase64 returns empty error")
		}
	})
}
