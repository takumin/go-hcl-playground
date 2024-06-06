package metadata

var (
	appName    string = "go-hcl-playground"
	appDesc    string = "Golang HCL Playground"
	authorName string = "Takumi Takahashi"
)

func AppName() string {
	return appName
}

func AppDesc() string {
	return appDesc
}

func AuthorName() string {
	return authorName
}
