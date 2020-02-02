# env

Package `env` enhances `os.Getenv` and `os.LookupEnv` and adds `FallbackLookup` and `MustLookup`.

## Use in testing

This package uses `os.Setenv("FOUND", "found")` to prepare its 'found' value, and `os.Unsetenv("MISSING")` to ensure 
environment pollution won't affect the test.

It also minimally tests `os.Getenv`, and `os.LookupEnv` by running them without testing results for coverage numbers.
Their behavior is tested thoroughly by the Go team.