## Dynamic Params
This package allows you to manage your parameters of any type, 
while allowing you to use different sources of data (now, only in-memory
 and cli arguments),
scan them based on a pattern or iterate through them and apply a callback,

Current Version: **1.0**

#### Tables of Contents
1. [Import](#import)
2. [Usage](#usage)
    1. [Quick Way](#quick-way)
    2. [Simple](#simple)
    3. [TypeConversion](#typeconversion)
    4. [Reading from Args](#reading-from-args)
    5. [Compound Types](#compound-types)
3. [List of Methods](#list-of-methods)
4. [Concurrency](#concurrency)
5. [Change Log](#change-log)
6. [Development](#development)
    
#### Import
```shell script
go get github.com/mostafatalebi/dynamic-params
```

#### Usage
Create an instance, and then add your params:

##### Quick Way
To quickly accessing values without worrying
or handling the errors, you can use Q methods.
Each method found in this doc has a Q method
corresponding to it. Q methods are good for 
ignoring errors because they only return
the value (and in case of error, a zero-value of the type).
If the method you want to use is
`GetAsString()`, then it becomes `QGetString()`

To create an instance with Q method quickly, use:
```go
dp := dynamic_params.QNewDynamicParams()
dp.Set("sample-param", 25)
v := dp.QGetInt("sample-param")
```
Note: `dynamic_params.QNewDynamicParams()` is like `NewDyanmicParams`
but it also can be invoked without any params. Still, if you want to
create dynamic params from argument list or pass mutex to the instance, you 
need to pass the parameters.

##### Simple
```go
p := dyanmic_params.NewDyanmicParams(SrcNameInternal)
p.Set("sample-int", 55)
val := p.Get("sample-int")
```

##### TypeConversion
If you have saved an int value, and you want to get it
as an int and not interface, you simple can use GetAs* group
of functions, and for your use case, it is `GetAsInt()` func.
Be careful to check for its errors.
```go
p := dyanmic_params.DyanmicParams(SrcNameInternal)
p.Set("sample-int", 55)
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
p.Set("key", &CustomType{Name: "Robert"})
v := p.Get("key")
assert.NotNil(t, v)
r, ok := v.(*CustomType)
assert.True(t, ok)
if r != nil {
    assert.Equal(t, "Robert", r.Name)
}
```


##### List of Methods
**Set**
Sets a key and a value. 

**Has**
Checks to see if a key exists or not.

**Get**
Returns the raw value passed when using `Set()`

**Count**
Returns the number of params

**Scan**
Scans param based on a given regex pattern, and returns the matched collection
or nil if none

**Iterate**
Iterates over params and applies the given callback

**GetAsString** or `QGetString()`
Tries to convert the value to `string` before returning, error if conversion fails.

**GetAsQuotedString** or `GetQuotedString()`
Tries to convert the value to `string`, and removes
 any surrounding single and double quotations 
 before returning, error if conversion fails.

**GetAsInt** or `GetInt()`
Tries to convert the value to `int` before returning, error if conversion fails.

**GetAsTimeDuration** or `QGetTimeDuration()`
Tries to convert the value to `time.Duration` before returning, error if conversion fails.

**GetStringAsInt** or `QGetStringAsInt()`
Tries to convert a non-zero starting, numeric string value
 to `int` before returning, error if conversion fails.
 
**GetStringAsTimeDuration** or `QGetStringAsTimeDuration()`
Tries to convert a duration string (1ms or 2h1m) to
a time.Duration type using time.ParseDuration() function

**GetStringAsBool** or `QGetStringAsBool()`
Tries to convert a non-zero starting, numeric string value
 to `bool` before returning, error if conversion fails.

**GetAsInt32** or `QGetInt32()`
Refer to `GetAsInt()`

**GetAsInt64** or `QGetInt64()`
Refer to `GetAsInt()`

**GetAsInt8** or `QGetInt8()`
Refer to `GetAsInt()`

**GetAsInt16** or `QGetInt16()`
Refer to `GetAsInt()`

**GetAsBool** or `QGetBool()`
Tries to convert the value to `bool` before returning, error if conversion fails.



##### Concurrency
By default, the instance you create is not concurrent safe, and
concurrent access to it leads to unexpected results, or in many cases,
fatal panic (due to concurrent map read and write). To make it concurrent
safe, pass a mutex to the instance in its creation time:
```go
lock := &sync.Mutex{}
dp := dynamic_params.NewDynamicParams(SrcNameInternal, lock)
dp.Set("sample-param", 25)
v := dp.QGetInt("sample-param")
```

##### Change Log

**1.0** 
- adding QMethods for quick use of methods

**<= 0.7**
- adding mutexes and lock for concurrency safety
- adding more methods to supports different go data types
- adding basic methods and supporting CLI argument lists
- adding iteration and scanning of params


##### Development

Upcoming features:
- Plan to support redis as a data source
- Plan to support mongo as a data source
- Plan to support JSON string as a data source
- Plan to add JSON marshalling and unmarshalling

