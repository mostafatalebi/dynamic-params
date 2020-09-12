## Dynamic Params
This package allows you to manage your parameters of any type, 
while allowing you to use different sources of data (now, only internal and args),
scan them based on a pattern or iterate through them ans apply a callback, 

#### Tables of Contents
1. [Import](#import)
2. [Usage](#usage)
    1. [Simple](#simple)
    2. [TypeConversion](#typeconversion)
    3. [Reading from Args](#reading-from-args)
    4. [Compound Types](#compound-types)
3. [List of Methods](#list-of-methods)
4. [Development](#development)
    
#### Import
```shell script
go get github.com/mostafatalebi/dynamic-params
```

#### Usage
Create an instance, and then add your params:

##### Simple
```go
p := dyanmic_params.DyanmicParams(SrcNameInternal)
p.Add("sample-int", 55)
val := p.Get("sample-int")
```

##### TypeConversion
If you have saved an int value, and you want to get it
as an int and not interface, you simple can use GetAs* group
of functions, and for your use case, it is `GetAsInt()` func.
Be careful to check for its errors.
```go
p := dyanmic_params.DyanmicParams(SrcNameInternal)
p.Add("sample-int", 55)
val, err := p.GetAsInt("sample-int")
```

##### Reading from Args
If you want to deal with values from argument list, 
you must know that our SrcNameArgs currently support this format:
```shell script
--name1=value --name2=value --other=otherValue ....
```
Upon creating an instance of `DynamicParams`, set source to
`SrcNameArgs` and pass your array of arguments (`os.Args`).
```go
p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=123456"})
v, err := p.GetStringAsInt("key")
assert.NoError(t, err)
assert.Equal(t, 123456, v)
```
To deal with args, it only supports returning values
in string, int or bool types.  If your arguments has a numeric
value not starting with zero, you can get it as int, or if it has 
a value of "false", "true", "1" or "0", you can get it as bool. 
`GetStringAs*()` group of functions are used to convert special values
of string to either int or bool.
```go
p := dp.NewDynamicParams(dp.SrcNameArgs, []string{"--key=true"})
v, err := p.GetStringAsBool("key")
assert.NoError(t, err)
assert.Equal(t, true, v)
```

You can save a param of any value, and upon getting the value, you either
can get the raw value for compound types (array, struct, map etc.) or if the
value was scalar, you can use helper methods to get a converted value.

##### Compound Types
To deal with other values such as struct or map, you should simply get
them as interface{} with `Get()` method and do the type conversion.

Examples:

**struct**

```go
type CustomType struct {
    Name string
}
p := dp.NewDynamicParams(dp.SrcNameInternal)
p.Add("key", &CustomType{Name: "Robert"})
v := p.Get("key")
assert.NotNil(t, v)
r, ok := v.(*CustomType)
assert.True(t, ok)
if r != nil {
    assert.Equal(t, "Robert", r.Name)
}
```


##### List of Methods
**Add**
Adds a key and a value. 

**Has**
Checks to see if a key exists or not.

**Get**
Returns the raw value passed when using `Add()`

**Count**
Returns the number of params

**Scan**
Scans param based on a given regex pattern, and returns the matched collection
or nil if none

**Iterate**
Iterates over params and applies the given callback

**GetAsString**
Tries to convert the value to `string` before returning, error if conversion fails.

**GetAsQuotedString**
Tries to convert the value to `string`, and removes
 any surrounding single and double quotations 
 before returning, error if conversion fails.

**GetAsInt**
Tries to convert the value to `int` before returning, error if conversion fails.

**GetStringAsInt**
Tries to convert a non-zero starting, numeric string value
 to `int` before returning, error if conversion fails.

**GetStringAsBool**
Tries to convert a non-zero starting, numeric string value
 to `bool` before returning, error if conversion fails.

**GetAsInt32**
Refer to `GetAsInt()`

**GetAsInt64**
Refer to `GetAsInt()`

**GetAsInt8**
Refer to `GetAsInt()`

**GetAsInt16**
Refer to `GetAsInt()`

**GetAsBool**
Tries to convert the value to `bool` before returning, error if conversion fails.




##### Development

Upcoming features:
- New Source -> Redis
- New Source -> MongoDB