package handle

import (
	"basic-gin/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllData(t *testing.T) {

	m := NewMember()
	result := m.AllData()

	var expect []model.Member

	assert.Equal(t, expect, result)

}

func TestAddData(t *testing.T) {

	m := NewMember()

	data := model.Member{
		Name:  "charan",
		Phone: "0123456789",
		Email: "charan@ru.ac.th",
	}

	_ = m.AddData(data)

	result := m.AllData()

	var expect []model.Member
	expect = append(expect, data)

	assert.Equal(t, expect, result)

}
