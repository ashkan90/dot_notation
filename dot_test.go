package dot_notation

import (
	"reflect"
	"testing"
)

type a struct {
	Customer struct {
		Name    string
		Surname string
	}
}

type g struct {
	Customer struct {
		Details struct {
			Profile struct {
				Name     string
				FullName string
			}
		}
	}
}

func TestNotate(t *testing.T) {
	n := &a{
		Customer: struct {
			Name    string
			Surname string
		}{Name: "Emirhan", Surname: "Ataman"},
	}

	ng := g{
		Customer: struct {
			Details struct {
				Profile struct {
					Name     string
					FullName string
				}
			}
		}{
			Details: struct {
				Profile struct {
					Name     string
					FullName string
				}
			}{
				Profile: struct {
					Name     string
					FullName string
				}{Name: "Emirhan", FullName: "Emirhan Ataman"}}},
	}

	mapExample := map[string]interface{}{
		"Name": "Emirhan",
		"Details": map[string]interface{}{
			"Name":    "My Company",
			"Address": "Company address",
			"MuchDetails": map[string]interface{}{
				"Detail": "test",
			},
		},
	}
	mapExpected := mapExample

	nExpected := "Ataman"
	ngExpected := "Emirhan Ataman"

	if notate := Notate(n, "Customer.Surname"); notate != nExpected {
		t.Errorf("Notate(...) = %q, expected = %q", notate, nExpected)
	}

	if notate := Notate(ng, "Customer.Details.Profile.FullName"); notate != ngExpected {
		t.Errorf("Notate(...) = %q, expected = %q", notate, ngExpected)
	}

	if mapNotate := Notate(mapExpected, "Details.MuchDetails.Detail"); !mapEqual(mapNotate, mapExpected) {
		t.Errorf("Notate(...) = %q, expected = %q", mapNotate, mapExpected)

	}
}

func mapEqual(comp1 interface{}, comp2 interface{}) bool {
	return reflect.DeepEqual(comp1, comp2)
}
