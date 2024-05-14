import { request } from '../request';
export function fetchCornList(id: string | number | undefined, param: Api.Dict.AddOrUpdateDictDataParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/dict/data/${id}`,
    method: 'put',
    data: param
  });
}
