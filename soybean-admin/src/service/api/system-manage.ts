import { request } from '../request';

// ------------------------------------------- 权限相关
/** get role list */
export function fetchGetRoleList(params?: Api.SystemManage.RoleSearchParams) {
  return request<Api.SystemManage.RoleList>({
    url: '/system/role/list',
    method: 'get',
    params
  });
}

/**
 * get all roles
 *
 * these roles are all enabled
 */
export function fetchGetAllRoles() {
  return request<Api.SystemManage.AllRole[]>({
    url: '/system/route/roles',
    method: 'get'
  });
}

export function addRole(param: Api.SystemManage.AddOrUpdateRoleParams) {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/role',
    method: 'post',
    data: param
  });
}

export function updateRoleById(id: string | number | undefined, param: Api.SystemManage.AddOrUpdateRoleParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/role/${id}`,
    method: 'put',
    data: param
  });
}

export function getRolePermitByCode(code: string | number | undefined) {
  return request<Api.SystemManage.RolePermit>({
    url: `/system/role/${code}`,
    method: 'get'
  });
}
export function updateRoleApiCodesPermitByCode(data: { roleCode: string | undefined; apiCodes: (string | number)[] }) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/role/apis`,
    method: 'put',
    data
  });
}
export function updateRoleMenuIdsPermitByCode(data: { roleCode: string | undefined; menuIds: string[] }) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/role/menus`,
    method: 'put',
    data
  });
}
export function updateRoleHomeByCode(data: { roleCode: string | undefined; home: string }) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/role/home`,
    method: 'put',
    data
  });
}

export function deleteRoleByIds(param: (string | number)[]) {
  // 发送携带菜单ID的DELETE请求至'/system/menu'接口
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/role',
    method: 'delete',
    data: { ids: param }
  });
}

// ------------------------------------------------------------------------------目录相关

/** get menu list */
export function fetchGetMenuList() {
  return request<Api.SystemManage.MenuList>({
    url: '/system/menu/list',
    method: 'get'
  });
}

/**
 * 添加菜单
 *
 * @param data 添加菜单所需参数，类型为 Api.SystemManage.AddOrUpdateMenuParams
 * @returns 返回一个 Promise，其解析值为 Api.BaseCurd.SuccessNodataResponse 类型
 */
export function addMenu(param: Api.SystemManage.AddOrUpdateMenuParams) {
  // 发送 POST 请求到 '/system/menu' 接口，携带参数 data
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/menu',
    method: 'post',
    data: param
  });
}

/**
 * 根据菜单ID更新菜单信息
 *
 * @param id 菜单的唯一标识符，可以是字符串或数字
 * @param param 更新菜单的参数对象，包含菜单的各种属性
 * @returns 返回一个Promise对象，其解析值为操作成功后的响应数据
 */
export function updateMenuById(id: string | number | undefined, param: Api.SystemManage.AddOrUpdateMenuParams) {
  // 使用PUT方法向'/system/menu/${id}'发送请求，将param作为请求体数据
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/menu/${id}`,
    method: 'put',
    data: param
  });
}

/**
 * 根据菜单ID批量删除菜单
 *
 * @param ids 要删除的菜单ID数组，可以是字符串或数字类型的数组
 * @returns 返回一个Promise对象，其解析值为操作成功后的响应数据
 */
export function deleteMenuByIds(param: (string | number)[]) {
  // 发送携带菜单ID的DELETE请求至'/system/menu'接口
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/menu',
    method: 'delete',
    data: { ids: param }
  });
}

/** get all pages */
export function fetchGetAllPages() {
  return request<string[]>({
    url: '/system/route/pages',
    method: 'get'
  });
}

/** get menu tree */
export function fetchGetMenuTree() {
  return request<Api.SystemManage.MenuTree[]>({
    url: '/system/route/tree',
    method: 'get'
  });
}
/** get button list */
export function fetchGetAllApis() {
  return request<Api.SystemManage.BackApi[]>({
    url: '/system/route/apis',
    method: 'get'
  });
}

// -----------------------------------------------------------------------------用户管理相关
/** get user list */
export function fetchGetUserList(params?: Api.SystemManage.UserSearchParams) {
  return request<Api.SystemManage.UserList>({
    url: '/system/user/list',
    method: 'get',
    params
  });
}

export function deleteUserByIds(param: (string | number)[]) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/api/system/user',
    method: 'delete',
    data: { ids: param }
  });
}

export function addUser(param: Api.SystemManage.AddOrUpdateUserParams) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: '/system/user',
    method: 'post',
    data: param
  });
}

export function updateUserById(id: string | number | undefined, param: Api.SystemManage.AddOrUpdateUserParams) {
  return request<Api.BaseCurd.SuccessNoDataResponse>({
    url: `/system/user/${id}`,
    method: 'put',
    data: param
  });
}
