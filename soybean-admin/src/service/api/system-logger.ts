import { request } from '../request';

export function fetchLoggerList(params?: Api.Dict.CommonSearchParams) {
  return request<Api.Logger.LoggerList>({
    url: `/system/log/${params?.type}`,
    method: 'get',
    params
  });
}

export function downLoggerGzFile(param: Api.SystemManage.AddOrUpdateRoleParams) {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/role',
    method: 'post',
    data: param
  });
}
