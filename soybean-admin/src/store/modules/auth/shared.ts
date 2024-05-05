import { localStg } from '@/utils/storage';

/** Get token */
export function getToken() {
  return localStg.get('token') || '';
}

/** Get user info */
export function getUserInfo() {
  const emptyInfo: Api.Auth.UserInfo = {
    userId: '',
    userName: '',
    roles: [],
    apis: []
  };
  const userInfo = localStg.get('userInfo') || emptyInfo;

  // fix new property: buttons, this will be removed in the next version `1.1.0`
  if (!userInfo.apis) {
    userInfo.apis = [];
  }

  return userInfo;
}

/** Clear auth storage */
export function clearAuthStorage() {
  localStg.remove('token');
  localStg.remove('refreshToken');
  localStg.remove('userInfo');
}
