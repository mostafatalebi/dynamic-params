## Dyanmic Params
You can use this go package for managing
your collection of parameters, without dealing with maps and conditional checks etc.

#### Import
```shell script
go get github.com/mostafatalebi/dyanmicparams
```

#### Usage
Create an instance, and then add your params:
```go
p := dyanmic_params.NewCustomParams()
p.Add("sample-int", 55)
val, err := p.GetAsInt("sample-int")
```

It supports `int`, `string` and `bool` type and more types will be supported soon.

**Args**

Or you can add items from argument list, which must be in --name=value format,
so if you run you program with `--url=some-url` then you can get it:
```go
p := dyanmic_params.NewCustomParams()
p.AddFromArgsAsInt("sample-int", os.Args)
val, err := p.GetAsInt("sample-int")
```

