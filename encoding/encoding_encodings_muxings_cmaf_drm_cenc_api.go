package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmCencApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafDrmCencCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafDrmCencApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmCencApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmCencApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafDrmCencCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmCencApi) Create(encodingId string, muxingId string, cencDrm model.CencDrm) (*model.CencDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.CencDrm(cencDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/cenc", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsCmafDrmCencApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/cenc/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmCencApi) Get(encodingId string, muxingId string, drmId string) (*model.CencDrm, error) {
    var resp *model.CencDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/cenc/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmCencApi) List(encodingId string, muxingId string, queryParams ...func(*query.CencDrmListQueryParams)) (*pagination.CencDrmsListPagination, error) {
    queryParameters := &query.CencDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.CencDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/cenc", &resp, reqParams)
    return resp, err
}