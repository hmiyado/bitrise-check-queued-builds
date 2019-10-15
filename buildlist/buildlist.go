package buildlist

import (
	"fmt"
	"os"

	apiclient "github.com/bitrise-io/bitrise-api-client/client"
	buildrequest "github.com/bitrise-io/bitrise-api-client/client/build_request"
	"github.com/bitrise-io/go-utils/log"
	httptransport "github.com/go-openapi/runtime/client"
)

// ListShow ... debug purpose
func ListShow(appSlug string) {
	response, err := list(appSlug)
	log.Printf("List Show")
	if err != nil {
		log.Printf("%s", err)
		return
	}
	log.Printf("%#v", response.Payload.Data)
}

// ListNum ...
func ListNum(appSlug string) int {
	response, err := list(appSlug)
	if err != nil {
		return 0
	}
	return len(response.Payload.Data)
}

func list(appSlug string) (*buildrequest.BuildRequestListOK, error) {
	apiToken := os.Getenv("BITRISE_TOKEN")
	apiKeyTokenAuth := httptransport.APIKeyAuth("Authorization", "header", fmt.Sprintf("token %s", apiToken))
	apiClient := apiclient.Default

	params := buildrequest.NewBuildRequestListParams().WithAppSlug(appSlug)
	return apiClient.BuildRequest.BuildRequestList(params, apiKeyTokenAuth)
}
