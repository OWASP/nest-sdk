# Errors


## Supported Types

### 

```go
errors := components.CreateErrorsArrayOfMapOfAny([]map[string]any{/* values here */})
```

### 

```go
errors := components.CreateErrorsMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch errors.Type {
	case components.ErrorsTypeArrayOfMapOfAny:
		// errors.ArrayOfMapOfAny is populated
	case components.ErrorsTypeMapOfAny:
		// errors.MapOfAny is populated
}
```
