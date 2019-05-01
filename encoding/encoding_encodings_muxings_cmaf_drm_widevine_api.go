package encoding
import (
    "github.com/bitmovin/bitmovin-api-sdk-go/common"
    "github.com/bitmovin/bitmovin-api-sdk-go/query"
    "github.com/bitmovin/bitmovin-api-sdk-go/model"
    "github.com/bitmovin/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsCmafDrmWidevineApi struct {
    apiClient *common.ApiClient
    Customdata *EncodingEncodingsMuxingsCmafDrmWidevineCustomdataApi
}

func NewEncodingEncodingsMuxingsCmafDrmWidevineApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsCmafDrmWidevineApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsCmafDrmWidevineApi{apiClient: apiClient}

    customdataApi, err := NewEncodingEncodingsMuxingsCmafDrmWidevineCustomdataApi(configs...)
    api.Customdata = customdataApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsCmafDrmWidevineApi) Get(encodingId string, muxingId string, drmId string) (*model.WidevineDrm, error) {
    var resp *model.WidevineDrm
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/widevine/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmWidevineApi) Delete(encodingId string, muxingId string, drmId string) (*model.BitmovinResponse, error) {
    var resp *model.BitmovinResponse
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
	}
    err := api.apiClient.Delete("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/widevine/{drm_id}", &resp, reqParams)
    return resp, err
}
func (api *EncodingEncodingsMuxingsCmafDrmWidevineApi) Create(encodingId string, muxingId string, widevineDrm model.WidevineDrm) (*model.WidevineDrm, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }
    payload := model.WidevineDrm(widevineDrm)
    
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/widevine", &payload, reqParams)
    return &payload, err
}
func (api *EncodingEncodingsMuxingsCmafDrmWidevineApi) List(encodingId string, muxingId string, queryParams ...func(*query.WidevineDrmListQueryParams)) (*pagination.WidevineDrmsListPagination, error) {
    queryParameters := &query.WidevineDrmListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }
    var resp *pagination.WidevineDrmsListPagination
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.QueryParams = queryParameters
	}
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/cmaf/{muxing_id}/drm/widevine", &resp, reqParams)
    return resp, err
}