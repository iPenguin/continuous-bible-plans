package resources

// List of static resource paths
const (
	_cssPath   = "/css"
	_jsPath    = "/js"
	_imagePath = "/images"
)

// Return the URL for this page for use in routing and other external functions
func PageNames() []string {
	return []string{_cssPath, _jsPath, _imagePath}
}
