package route

import "fmt"

const (
    API_V1_PREFIX = "/api/v1"
    API_V1_UPLOAD_SUBPATH = "/upload"
)

var (
    ApiV1UploadRoute = fmt.Sprintf("%s%s", API_V1_PREFIX, API_V1_UPLOAD_SUBPATH)
)
