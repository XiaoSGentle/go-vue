import { downloadFile, request } from '../request';

export function fetchLoggerList(params?: Api.Dict.CommonSearchParams) {
  return request<Api.Logger.LoggerList>({
    url: `/system/log/${params?.type}`,
    method: 'get',
    params
  });
}

export function fetchLoggerFileList() {
  return request<Api.Logger.LoggerFiles[]>({
    url: '/system/log/list',
    method: 'get'
  });
}
export function downloadLoggerFile(fileName: string) {
  downloadFile(`/system/log/download/${fileName}`, fileName);
}
