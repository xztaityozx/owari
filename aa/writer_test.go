package aa

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type dummyWriter struct{}

func (d dummyWriter) Write(_ []byte) (n int, err error) {
	panic("implement me")
}

type SetterTestSuite struct {
	suite.Suite
	w Writer
}

func (suite *SetterTestSuite) SetupTest() {
	suite.w = NewWriter(dummyWriter{})
}

func Test_NewWriter(t *testing.T) {
	aaw := NewWriter(dummyWriter{})

	assert.NotNil(t, aaw)
	assert.False(t, aaw.overwrite, "overwriteはデフォルトでFalseであるべき")
	assert.False(t, aaw.colorful, "colorfulはデフォルトでFalseであるべき")
	assert.False(t, aaw.colorfulAlways, "colorfulAlwaysはデフォルトでFalseであるべき")

	assert.Equal(t, 0, aaw.offset, "offsetはデフォルトで0であるべき")
	assert.Equal(t, 1, aaw.times, "timesはデフォルトで1であるべき")
	assert.Equal(t, time.Duration(0), aaw.duration, "durationはデフォルトで0であるべき")
}

// Test_SetterTestSuite はSet**系のテストスイートを開始する
func Test_SetterTestSuite(t *testing.T) {
	suite.Run(t, new(SetterTestSuite))
}

func (suite *SetterTestSuite) Test_Writer_SetColorful_To_False() {
	suite.w.SetColorful(false)
	suite.False(suite.w.colorful, "colorfulがtureであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetColorfulAlways_To_False() {
	suite.w.SetColorfulAlways(false)
	suite.False(suite.w.colorful, "colorfulがtureであるべき")
	suite.False(suite.w.colorfulAlways, "colorfulAlwaysがtureであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetOverwrite_To_False() {
	suite.w.SetOverwrite(false)
	suite.False(suite.w.overwrite, "overwriteがfalseであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetColorful_To_True() {
	suite.w.SetColorful(true)
	suite.True(suite.w.colorful, "colorfulがtureであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetColorfulAlways_To_True() {
	suite.w.SetColorfulAlways(true)
	suite.True(suite.w.colorful, "colorfulがtureであるべき")
	suite.True(suite.w.colorfulAlways, "colorfulAlwaysがtureであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetOverwrite_To_True() {
	suite.w.SetOverwrite(true)
	suite.True(suite.w.overwrite, "overwriteがtrueであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetBeginEmpty_To_True() {
	suite.w.SetInsertEmpty(true)
	suite.True(suite.w.beginEmpty, "beginEmptyがtrueであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetBeginEmpty_To_False() {
	suite.w.SetInsertEmpty(false)
	suite.False(suite.w.beginEmpty, "beginEmptyがfalseであるべき")
}

func (suite *SetterTestSuite) Test_Writer_SetOffset() {
	offset := rand.Int()
	suite.w.SetOffset(offset)

	suite.Equal(offset, suite.w.offset, "offsetの値が設定されているべき")
}

func (suite *SetterTestSuite) Test_Writer_SetTimes() {
	times := rand.Int()
	suite.w.SetTimes(times)

	suite.Equal(times, suite.w.times, "timesの値が設定されているべき")
}

func (suite *SetterTestSuite) Test_Writer_SetDuration() {
	times := time.Duration(10)
	suite.w.SetDuration(times)

	suite.Equal(times, suite.w.duration, "durationの値が設定されているべき")
}

func TestWriter_Write(t *testing.T) {
	var b []byte
	buf := bytes.NewBuffer(b)

	w := NewWriter(buf)

	aa := []string{
		"line 1", "line 2",
	}

	err := w.Write(aa)

	assert.Nil(t, err)
	assert.Equal(t, append(aa, ""), strings.Split(buf.String(), "\n"))
}
