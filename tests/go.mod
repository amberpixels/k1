// k1's test suite, deliberately its own module. This is permanent, not a workaround:
// it is what lets k1's own go.mod stay at zero requirements, so importing k1 adds
// nothing to anyone's module graph. It also decouples the two Go floors - `be`
// currently needs 1.26 while k1 promises 1.25 (see `just floor`). Even once those
// floors converge, the split stays; only the `floor` recipe would become redundant.
module github.com/amberpixels/k1/tests

go 1.26

require (
	github.com/amberpixels/k1 v0.1.6
	github.com/expectto/be v1.0.0-rc.8
)

require (
	github.com/IGLOU-EU/go-wildcard v1.0.3 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/onsi/gomega v1.42.1 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/text v0.38.0 // indirect
)

replace github.com/amberpixels/k1 => ../
