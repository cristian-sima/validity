package validity

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {

	// value

	Convey("Given the validation rule \"value:0,100\"", t, func() {

		rule := Rules{"Bar": Field{
			Name:  "Baz",
			Type:  "Int",
			Rules: []string{"value:0,100"},
		}}

		Convey("Given the value -1", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": -1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 0", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 0}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 1", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 100}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 101}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// value_strict

	Convey("Given the validation rule \"value_strict:0,100\"", t, func() {

		rule := Rules{"Bar": Field{
			Name:  "Baz",
			Type:  "Int",
			Rules: []string{"value_strict:0,100"},
		},
		}
		Convey("Given the value -1", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": -1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 0", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 0}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 1", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 100}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 101", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 101}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits

	Convey("Given the validation rule \"digits:3\"", t, func() {

		rule := Rules{"Bar": Field{
			Type:  "Int",
			Name:  "Foo",
			Rules: []string{"digits:3"},
		},
		}

		Convey("Given the value 10", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 10}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 100", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 100}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 1000", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 1000}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits_between

	Convey("Given the validation rule \"digits_between:2,5\"", t, func() {

		rule := Rules{"Bar": Field{
			Type:  "Int",
			Name:  "Foo",
			Rules: []string{"digits_between:2,5"},
		},
		}

		Convey("Given the value 1", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 12", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 12}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 123", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 123}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 12345", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 12345}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 123456", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 123456}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

	// digits_between_strict

	Convey("Given the validation rule \"digits_between_strict:2,5\"", t, func() {

		rule := Rules{"Bar": Field{
			Type:  "Int",
			Name:  "Foo",
			Rules: []string{"digits_between_strict:2,5"},
		},
		}

		Convey("Given the value 1", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 1}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 12", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 12}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 123", func() {

			Convey("The result should be valid", func() {
				data := map[string]interface{}{"Bar": 123}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeTrue)
			})

		})

		Convey("Given the value 12345", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 12345}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

		Convey("Given the value 123456", func() {

			Convey("The result should not be valid", func() {
				data := map[string]interface{}{"Bar": 123456}
				result := Validate(data, rule)
				So(result.IsValid, ShouldBeFalse)
			})

		})

	})

}
