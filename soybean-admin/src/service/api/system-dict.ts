import { request } from '../request';

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

export function updateDictTypeById(id: string | number | undefined, param: Api.Dict.AddOrUpdateDictTypeParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/dict/${id}`,
    method: 'put',
    data: param
  });
}

export function deleteDictTypeByIds(param: (string | number)[]) {
  // 发送携带菜单ID的DELETE请求至'/system/menu'接口
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
