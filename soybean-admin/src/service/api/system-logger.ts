import { request } from '../request';

export function fetchLoggerList(params?: Api.Dict.CommonSearchParams) {
  return request<Api.Logger.LoggerList>({
    url: `/system/log/${params?.type}`,
    method: 'get',
    params
  });
}

export function downLoggerGzFile() {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.Logger.LoggerFiles[]>({
    url: '/system/log/list',
    method: 'get'
  });
}
