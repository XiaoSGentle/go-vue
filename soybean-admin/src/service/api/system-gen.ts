import { request } from '../request';
export function fetchCanGenTables() {
  return request<string[]>({
    url: `/system/table/models`,
    method: 'get'
  });
}
export function fetchGenTables(param: Api.CodeGen.CommonSearchParams) {
  return request<Api.CodeGen.TableInfoList>({
    url: `/system/table/list`,
    method: 'get',
    params: param
  });
}
export function deleteTableByIds(param: (string | number)[]) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/table',
    method: 'delete',
    data: { ids: param }
  });
}

export function addTable(param: string[]) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/table/gen',
    method: 'post',
    data: {
      names: param
    }
  });
}

export function updateTableById(id: string | number | undefined, param: Api.CodeGen.AddOrUpdateTableInfoDataParams) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/table/${id}`,
    method: 'put',
    data: param
  });
}
export function fetchTableColumns(param: Api.CodeGen.TableColumnSearchParams) {
  return request<Api.CodeGen.TableColumnInfoList>({
    url: `/system/column/list`,
    method: 'get',
    params: param
  });
}

export function updateTableColumnsById(
  id: string | number | undefined,
  param: Api.CodeGen.AddOrUpdateTableColumnsInfoDataParams
) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/column/${id}`,
    method: 'put',
    data: param
  });
}
