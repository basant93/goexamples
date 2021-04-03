//go version
//go env -json
//go  env GOOS GOARCH
//go list .
//list the import path that we need to add to import the package
//go list github.com/basant93/goexamples/gotest/gocover/
//list of all packages that are available for import
//go list -f {{.GoFiles}} github.com/basant93/goexamples/gotest/gocover/
//get dependent files
//go list -f {{.Deps}} net/http
//go list -json -f {{.Deps}} net/http
//-e flag for returning empty list if package dependency are not present
//go doc text/template.new
//go doc time
//godoc -http :8000

//get package if it is not present
//update the package
//go get -u time
//-d flag: download the package but not install it
//-f flag: install but do not verify from where it came from
//go bug


