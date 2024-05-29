import { request } from '../request';
export function fetchGetSysCronList(param: Api.Cron.CommonSearchParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.Cron.CronTypeList>({
    url: `/system/cron/list`,
    method: 'get',
    params: param
  });
}
export function updateCronDataById(param: Api.Cron.AddOrUpdateCronTypeParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/cron`,
    method: 'put',
    data: param
  });
}
