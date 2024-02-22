# Go-genenum
This is a tool to generate enums and common optimized helpers for them.

## Installation
```bash
go install github.com/saturn4er/go-genenum/cmd/go-genenum  
```

## Basic usage
Create a file with enums declaration in yaml format. For example, `enums.yaml`:
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
```
then run `go-genenum enums.yaml -o enums.go -p my_package` and it will generate `enums.go`

```go
package my_package

type A byte

const (
	AValue1 A = iota + 1
	AValue2
	AValue3
)
```

## Helpers
Usually, you need some additional methods for enums, like `String()`, `IsValid()`, `IsValue1()`, 
or create some group of enum values, so you can easily check if this value belongs to some group. 
But you don't want to support it, as it's really easy to forget to add some handler for new value.

`go-genenum` can generate these methods for you. Just add `helpers` section to your yaml file:
### String
String helper will generate optimized `String()` method for your enum.
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
  helpers:
    string: true
```
generates:
```go
func (a A) String() string {
	if a < 1 || a > 3 {
		return "A(" + strconv.FormatInt(int64(a), 10) + ")"
	}
	const names = "Value1Value2Value3"

	var indexes = [...]int32{0, 6, 12, 18}

	return names[indexes[a-1]:indexes[a]]
}
```

### IsValid
IsValid helper will generate `IsValid()` method, which will check if value is valid for this enum.
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
  helpers:
    is_valid: true
```
generates:
```go
func (a A) IsValid() bool {
	return a > 0 && a < 4
}
```

### Validate
Validate helper will generate `Validate()` method, which will check if value is valid for this enum and return error if not.
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
  helpers:
    validate: true
```
generates:
```go
type InvalidAValueError byte

func (e InvalidAValueError) Error() string {
	return fmt.Sprintf("invalid A(%d)", e)
}
func (a A) Validate() error {
	if a < 1 || a > 4 {
		return InvalidAValueError(a)
	}
	return nil
}
```

### Is
Is helper will generate methods for each value, which will check if value is equal to this value.
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
  helpers:
    is: true
```
generates:
```go
func (a A) IsValue1() bool {
	return a == AValue1
}
func (a A) IsValue2() bool {
	return a == AValue2
}
func (a A) IsValue3() bool {
	return a == AValue3
}
```

### Categories
Categories helper will generate methods for each category, which will check if value belongs to this category.
```yaml
- name: A
  values: [ "Value1", "Value2", "Value3" ]
  helpers:
    categories:
      - name: "First"
        values: [ "Value1", "Value2" ]
      - name: "Second"
        values: [ "Value3" ]
```
generates:
```go
var FirstAS = []A{
	AValue1,
	AValue2,
}

func (a A) IsFirst() bool {
	if a < 1 || a > 4 {
		return false
	}
	return []bool{false, true, true, false}[a]
}

var SecondAS = []A{
	AValue3,
}

func (a A) IsSecond() bool {
	if a < 1 || a > 4 {
		return false
	}
	return []bool{false, false, false, true}[a]
}
```
