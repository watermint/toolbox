package report_column

import (
	"go.uber.org/zap"
	"testing"
)

func TestColumnMarshaller_ColumnTypes(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}

	type Member1 struct {
		Id    int    `column:"id"`
		Name  string `column:"name"`
		Email string `column:"email"`
	}
	m1 := Member1{
		Id:    123,
		Name:  "abc",
		Email: "abc@example.com",
	}
	cm1 := ColumnMarshaller{
		Logger: log,
	}

	ct1 := cm1.ColumnTypes(m1)
	if len(ct1) != 3 ||
		ct1[0].Tag != "id" ||
		ct1[1].Tag != "name" ||
		ct1[2].Tag != "email" {
		t.Error("Invalid column types")
	}

	// pointer
	ct1 = cm1.ColumnTypes(&m1)
	if len(ct1) != 3 ||
		ct1[0].Tag != "id" ||
		ct1[1].Tag != "name" ||
		ct1[2].Tag != "email" {
		t.Error("Invalid column types")
	}

	type Member2 struct {
		Id    int    `column:"id"`
		Name  string `column:"name"`
		Data  []byte `column:"-"`
		Email string `column:"email"`
	}
	m2 := Member2{
		Id:    123,
		Name:  "abc",
		Data:  []byte("ABC"),
		Email: "abc@example.com",
	}
	cm2 := ColumnMarshaller{
		Logger: log,
	}

	ct2 := cm2.ColumnTypes(m2)
	if len(ct2) != 3 ||
		ct2[0].Tag != "id" ||
		ct2[1].Tag != "name" ||
		ct2[2].Tag != "email" {
		t.Error("Invalid column types")
	}

	// pointer
	ct2 = cm1.ColumnTypes(&m2)
	if len(ct2) != 3 ||
		ct2[0].Tag != "id" ||
		ct2[1].Tag != "name" ||
		ct2[2].Tag != "email" {
		t.Error("Invalid column types")
	}
}

func TestColumnMarshaller_Row(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}

	type Member1 struct {
		Id    int    `column:"id"`
		Name  string `column:"name"`
		Email string `column:"email"`
	}
	m1 := Member1{
		Id:    123,
		Name:  "abc",
		Email: "abc@example.com",
	}
	cm1 := ColumnMarshaller{
		Logger: log,
	}

	if ct1, err := cm1.Row(m1); err != nil {
		t.Error(err)
	} else {
		if len(ct1) != 3 ||
			ct1[0].ColumnName != "id" || ct1[0].Value != "123" ||
			ct1[1].ColumnName != "name" || ct1[1].Value != "abc" ||
			ct1[2].ColumnName != "email" || ct1[2].Value != "abc@example.com" {
			t.Error("Invalid column types")
		}
	}

	//pointer
	if ct1, err := cm1.Row(&m1); err != nil {
		t.Error(err)
	} else {
		if len(ct1) != 3 ||
			ct1[0].ColumnName != "id" || ct1[0].Value != "123" ||
			ct1[1].ColumnName != "name" || ct1[1].Value != "abc" ||
			ct1[2].ColumnName != "email" || ct1[2].Value != "abc@example.com" {
			t.Error("Invalid column types")
		}
	}

	type Member2 struct {
		Id    int    `column:"id"`
		Name  string `column:"name"`
		Data  []byte `column:"-"`
		Email string `column:"email"`
	}
	m2 := Member2{
		Id:    123,
		Name:  "abc",
		Data:  []byte("ABC"),
		Email: "abc@example.com",
	}
	cm2 := ColumnMarshaller{
		Logger: log,
	}

	if ct2, err := cm2.Row(m2); err != nil {
		t.Error(err)
	} else {
		if len(ct2) != 3 ||
			ct2[0].ColumnName != "id" || ct2[0].Value != "123" ||
			ct2[1].ColumnName != "name" || ct2[1].Value != "abc" ||
			ct2[2].ColumnName != "email" || ct2[2].Value != "abc@example.com" {
			t.Error("Invalid column types")
		}
	}

	// pointer
	if ct2, err := cm2.Row(&m2); err != nil {
		t.Error(err)
	} else {
		if len(ct2) != 3 ||
			ct2[0].ColumnName != "id" || ct2[0].Value != "123" ||
			ct2[1].ColumnName != "name" || ct2[1].Value != "abc" ||
			ct2[2].ColumnName != "email" || ct2[2].Value != "abc@example.com" {
			t.Error("Invalid column types")
		}
	}

}
