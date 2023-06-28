package service

import (
	"errors"
	"github.com/basebytes/elastic-go/client"
	"github.com/basebytes/elastic-go/client/api"
	"github.com/basebytes/elastic-go/client/entity"
	"github.com/basebytes/elastic-go/service/constructor"
	"net/http"
	"sort"
	"strings"
)

type Service struct {
	Constructor *constructor.Constructor
	client      *client.Client
}

func NewDefaultService() (service *Service, err error) {
	var esClient *client.Client
	if esClient, err = client.NewDefaultClient(); err == nil {
		service = &Service{
			Constructor: constructor.New(),
			client:      esClient,
		}
	}
	return service, err
}

func NewService(config *client.Config) (*Service, error) {
	service := &Service{
		Constructor: constructor.New(),
	}
	esClient, err := client.NewClient(config)
	if err != nil {
		return nil, err
	}
	service.client = esClient
	return service, nil
}

func (service *Service) Bulk(index string, docs []api.Doc, failedInfoOnly bool) (*entity.DocumentBulkResult, error) {
	var (
		rsp    *api.Response
		err    error
		result *entity.DocumentBulkResult
		bulk   = service.client.Document.Bulk
		opt    []func(*api.DocumentBulkRequest)
	)
	opt = append(opt, bulk.WithIndex(index))
	if failedInfoOnly {
		opt = append(opt, bulk.WithFilterPath("items.*.error"))
	}
	if rsp, err = bulk(docs, opt...); err == nil {
		result = new(entity.DocumentBulkResult)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) IndexExists(index string) (bool, error) {
	var (
		rsp         *api.Response
		err         error
		code        = http.StatusInternalServerError
		indexExists = service.client.Indexer.Exists
	)
	if rsp, err = indexExists(indexExists.WithIndex(index)); err == nil {
		code = rsp.StatusCode
	}
	return code == http.StatusOK, err
}

func (service *Service) GetIndexAlias(index, alias []string) (*entity.IndexAlias, error) {
	var (
		rsp      *api.Response
		err      error
		result   *entity.IndexAlias
		aliasGet = service.client.Indexer.AliasGet
	)
	if rsp, err = aliasGet(aliasGet.WithIndex(index...), aliasGet.WithAlias(alias...)); err == nil {
		result = new(entity.IndexAlias)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) GetIndexMapping(index string) (*entity.IndexMapping, error) {
	var (
		rsp        *api.Response
		err        error
		result     *entity.IndexMapping
		getMapping = service.client.Indexer.GetMapping
	)
	if rsp, err = getMapping(getMapping.WithIndex(index)); err == nil {
		result = new(entity.IndexMapping)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) GetIndexSetting(index string, settings ...string) (*entity.IndexSetting, error) {
	var (
		rsp        *api.Response
		err        error
		result     *entity.IndexSetting
		getSetting = service.client.Indexer.SettingsGet
	)
	if rsp, err = getSetting(getSetting.WithIndex(index), getSetting.WithSettings(settings...)); err == nil {
		result = new(entity.IndexSetting)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) CreateIndex(index string, body *[]byte) (*entity.OperateResult, error) {
	var (
		rsp    *api.Response
		err    error
		result *entity.OperateResult
		create = service.client.Indexer.Create
	)
	if rsp, err = create(index, create.WithBody(body)); err == nil {
		result = new(entity.OperateResult)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) DeleteIndex(index ...string) (*entity.OperateResult, error) {
	var (
		rsp         *api.Response
		err         error
		result      *entity.OperateResult
		indexDelete = service.client.Indexer.Delete
	)
	if rsp, err = indexDelete(index); err == nil {
		result = new(entity.OperateResult)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) GetIndexNames(names ...string) ([]string, error) {
	var (
		rsp        *api.Response
		err        error
		indices    []string
		catIndices = service.client.Cat.Indices
	)
	if rsp, err = catIndices(catIndices.WithIndex(names...), catIndices.WithColumns("index")); err == nil {
		indices = strings.Split(strings.TrimSpace(rsp.ResultString()), "\n")
		sort.Strings(indices)
	}
	return indices, err
}

func (service *Service) UpdateIndexAlias(actions ...*api.AliasAction) (*entity.OperateResult, error) {
	var (
		rsp     *api.Response
		err     error
		result  *entity.OperateResult
		aliases = service.client.Indexer.Aliases
	)
	if rsp, err = aliases(actions); err == nil {
		result = new(entity.OperateResult)
		err = rsp.Result(result)
	}
	return result, err
}

func (service *Service) Ping() bool {
	_, err := service.client.Ping()
	return err == nil
}

func (service *Service) Search(indexName string, body *[]byte) (*entity.EsQueryResult, error) {
	var result *entity.EsQueryResult
	searcher := service.client.Searcher.Search
	resp, err := searcher(
		searcher.WithIndex(indexName),
		searcher.WithBody(body),
		searcher.WithTrackTotalHits(true),
	)
	if err == nil {
		err = resp.Result(&result)
	}

	return result, err
}

func (service *Service) CatIndices(indices string) ([]*entity.IndexInfo, error) {
	var (
		result []*entity.IndexInfo
		esErr  entity.EsError
	)
	cats := service.client.Cat.Indices
	resp, err := cats(
		cats.WithIndex(indices),
		cats.WithFormat("json"),
	)
	if err == nil {
		if resp.IsError() {
			if err = resp.Result(&esErr); err == nil {
				err = errors.New(esErr.Error)
			}
		} else {
			err = resp.ResultJson(&result)
		}
	}

	return result, err
}
