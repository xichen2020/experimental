The microbenchmarks validated a few claims in https://planetscale.com/blog/generics-can-make-your-go-code-slower when running against go 1.20 on my laptop, namely:
* A generic function with numeric type parameters has similar performance to implementing an equivalent non-generic function once per numeric type;
* A generic function with byte sequence type parameter has similar performance to implementing an equivalent non-generic function separately for string and []byte;
* A generic function with an interface type parameter and called with an argument with an interface type is shown to have worse performance than implementing an equivalent non-generic function where the said argument has the raw, non-interface type (e.g., a pointer type).

In practice, whether to implement something using generics not only depends on its performance impact, but also code readability and maintainability. A more readable, slightly slower implementation using generics that is not on the critical path can still be a preferable choice. Measure before premature optimization.
