# goutils
!['Status Badge'](https://github.com/incu6us/goutils/workflows/build/badge.svg)


Utils for day-to-day work

## Assert library
* assert.Equals - lib to compare two structs(or its pointers, or generic types) with skipping a list of fields for exclusion from comparing.
(supports only first level fields)
Example:
```go
        actual := &testStruct{
            ID:          1,
			Value:       "test",
			CreatedAt:   time.Now(),
			InnerStruct: innerStruct{Value: "123"},
		}

        expected := &testStruct{
            ID:          1,
			Value:       "test",
			CreatedAt:   time.Now(),
			InnerStruct: innerStruct{Value: "123"},
		}

        if assert.Equals(expected, actual, []{"CreatedAt"}...){
            // do some stuff
        }           
```
* assert.EqualsWithDiffFunc - the same, but with middleware func to pass a custom function to check the equality.
