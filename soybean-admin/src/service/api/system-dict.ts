import { request } from '../request';

/**
 * 函数 fetchGetDictTypeList 发送一个 GET 请求，从 '/system/dict/list' 端点检索字典类型列表。
 *
 * @param [params] - `fetchGetDictTypeList` 函数用于从服务器获取字典类型列表。`params` 参数是一个可选对象，可以传递该参数来指定用于筛选字典类型的任何搜索参数。这些参数将作为 GET
 *   请求中的查询参数发送到
 * @returns 函数“fetchGetDictTypeList”正在返回从 API 获取字典类型列表的请求。
 */
export function fetchGetDictTypeList(params?: Api.Dict.DictTypeDataSearchParams) {
  return request<Api.Dict.DictTypeList>({
    url: '/system/dict/list',
    method: 'get',
    params
  });
}

export function addDictType(param: Api.Dict.AddOrUpdateDictTypeParams) {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/dict',
    method: 'post',
    data: param
  });
}

/**
 * 根据ID更新字典类型信息
 *
 * @param id 字典类型的ID，可以是字符串、数字或未定义，用于指定要更新的字典类型
 * @param param 包含要更新的字典类型信息的对象，参见 Api.Dict.AddOrUpdateDictTypeParams
 * @returns 返回一个Promise，其解析为Api.BaseCurd.SuccessNoDataResponse类型的对象，表示操作是否成功，不返回具体的数据
 */
export function updateDictTypeById(id: string | number | undefined, param: Api.Dict.AddOrUpdateDictTypeParams) {
  // 发起一个PUT请求，更新指定ID的字典类型信息
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/dict/${id}`,
    method: 'put',
    data: param
  });
}

export function deleteDictTypeByIds(param: (string | number)[]) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/dict',
    method: 'delete',
    data: { ids: param }
  });
}

export function fetchGetDictDataList(params?: Api.Dict.DictDataSearchParams) {
  return request<Api.Dict.DictDataList>({
    url: '/system/dict/data/list',
    method: 'get',
    params
  });
}

export function addDictData(param: Api.Dict.AddOrUpdateDictDataParams) {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/dict/data',
    method: 'post',
    data: param
  });
}

export function updateDictDataById(id: string | number | undefined, param: Api.Dict.AddOrUpdateDictDataParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/dict/data/${id}`,
    method: 'put',
    data: param
  });
}

export function deleteDictDataByIds(param: (string | number)[]) {
  // 发送携带菜单ID的DELETE请求至'/system/menu'接口
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/dict/data',
    method: 'delete',
    data: { ids: param }
  });
}

export function fetchDict(param: string) {
  return request<Api.Dict.Dict[]>({
    url: `/dict/${param}`,
    method: 'get'
  });
}
