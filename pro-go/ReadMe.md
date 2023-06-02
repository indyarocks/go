## Go packages
1. https://pkg.go.dev
2. https://github.com/golang/go/wiki/Projects

### Struct initialisation
var product = new(Product)
OR
var product = Product{}

### Compositions
field promotion - If the enclosing type has a field same as a field of
embedded type's field, then the direct access to the attribute will
render enclosing type's field value.
Fields and Methods can be promoted to the enclosing type from embedded type
```go
type Boat struct {
        *Product
        Capacity  int
        Motorized bool
}

type Product struct {
        Name, Category string
        price          float64
}

func (p *Product)Price(tax) float64{
	return p.price + p.price*tax
}
```
`Name` can be accessed by either `boat.Product.Name` and `boat.Name`
`boat.Price(0.2)` or `boat.Product.Price(0.2)` - Both are equivalent.

