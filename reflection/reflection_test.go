package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Addr Address
}

type Address struct {
	City    string
	Country string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with only one field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two fields",
			Input: struct {
				Name string
				City string
			}{"RunAt", "Delhi"},
			ExpectedCalls: []string{"RunAt", "Delhi"},
		},
		{
			Name: "struct with non-string value",
			Input: struct {
				Name string
				Age  int
			}{"RunAt", 22},
			ExpectedCalls: []string{"RunAt"},
		},
		{
			Name: "struct with a struct inside (nested structs)",
			Input: Person{
				Name: "RunAt",
				Age:  22,
				Addr: Address{
					City:    "Delhi",
					Country: "India",
				},
			},
			ExpectedCalls: []string{"RunAt", "Delhi", "India"},
		},
		{
			Name: "Pointer to struct",
			Input: &Person{
				Name: "Clidle",
				Age:  2,
				Addr: Address{
					City:    "New York",
					Country: "USA",
				},
			},
			ExpectedCalls: []string{"Clidle", "New York", "USA"},
		},
		{
			Name: "slices",
			Input: []Address{
				{"New York", "USA"},
				{"Delhi", "India"},
				{"Madrid", "Spain"},
			},
			ExpectedCalls: []string{"New York", "USA", "Delhi", "India", "Madrid", "Spain"},
		},
		{
			Name: "arrays",
			Input: [2]Address{
				{"Delhi", "India"},
				{"Madrid", "Spain"},
			},
			ExpectedCalls: []string{"Delhi", "India", "Madrid", "Spain"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v but wanted %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"RunAt":     "Me",
			"RunAtPapa": "Papa",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Me")
		assertContains(t, got, "Papa")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected to find %q in haystack %v", needle, haystack)
	}
}
