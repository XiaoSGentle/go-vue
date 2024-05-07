/**
 * Namespace Api
 *
 * All backend api type
 */
declare namespace Api {
  namespace Common {
    /** common params of paginating */
    interface PaginatingCommonParams {
      /** current page number */
      current: number;
      /** page size */
      size: number;
      /** total count */
      total: number;
    }

    /** common params of paginating query list data */
    interface PaginatingQueryRecord<T = any> extends PaginatingCommonParams {
      records: T[];
    }

    /**
     * enable status
     *
     * - "1": enabled
     * - "2": disabled
     */
    type EnableStatus = '1' | '2';

    /** common record */
    type CommonRecord<T = any> = {
      /** record id */
      id: number;
      /** record creator */
      createBy: string;
      /** record create time */
      createTime: string;
      /** record updater */
      updateBy: string;
      /** record update time */
      updateTime: string;
      /** record status */
      status: EnableStatus | null;
    } & T;
  }

  /**
   * namespace Auth
   *
   * backend api module: "auth"
   */
  namespace Auth {
    interface LoginToken {
      token: string;
      refreshToken: string;
    }

    interface UserInfo {
      userId: string;
      userName: string;
      roles: string[];
      apis: string[];
    }
  }

  /**
   * namespace Route
   *
   * backend api module: "route"
   */
  namespace Route {
    type ElegantConstRoute = import('@elegant-router/types').ElegantConstRoute;

    interface MenuRoute extends ElegantConstRoute {
      id: string;
    }

    interface UserRoute {
      routes: MenuRoute[];
      home: import('@elegant-router/types').LastLevelRouteKey;
    }
  }

  namespace BaseCurd {
    type SuccessNoDataResponse = string;
  }

  /**
   * namespace SystemManage
   *
   * backend api module: "systemManage"
   */
  namespace SystemManage {
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    type RolePermit = {
      menuIds: string[];
      apiCodes: string[];
    };

    /** role */
    type Role = Common.CommonRecord<{
      /** role name */
      roleName: string;
      /** role code */
      roleCode: string;
      /** role description */
      roleDesc: string;

      roleHome: string;
      menuIds: string[];
      apiCodes: string[];
    }>;

    type AddOrUpdateRoleParams = Pick<
      Api.SystemManage.Role,
      'roleName' | 'roleCode' | 'roleDesc' | 'status' | 'apiCodes' | 'menuIds'
    >;
    /** role search params */
    type RoleSearchParams = CommonType.RecordNullable<
      Pick<Api.SystemManage.Role, 'roleName' | 'roleCode' | 'status'> & CommonSearchParams
    >;

    /** role list */
    type RoleList = Common.PaginatingQueryRecord<Role>;

    /** all role */
    type AllRole = Pick<Role, 'id' | 'roleName' | 'roleCode'>;

    /**
     * user gender
     *
     * - "1": "male"
     * - "2": "female"
     */
    type UserGender = '1' | '2';

    type AddOrUpdateUserParams = Pick<
      Api.SystemManage.User,
      'userName' | 'userGender' | 'nickName' | 'userPhone' | 'userEmail' | 'userRoles' | 'status'
    >;

    /** user */
    type User = Common.CommonRecord<{
      /** user name */
      userName: string;
      /** user gender */
      userGender: UserGender | null;
      /** user nick name */
      nickName: string;
      /* user lastOnLine */
      lastOnLine: string;
      /* user lastCpWd */
      lastCpWd: string;
      /** user phone */
      userPhone: string;
      /** user email */
      userEmail: string;
      /** user role code collection */
      userRoles: string[];
    }>;

    /** user search params */
    type UserSearchParams = CommonType.RecordNullable<
      Pick<Api.SystemManage.User, 'userName' | 'userGender' | 'nickName' | 'userPhone' | 'userEmail' | 'status'> &
        CommonSearchParams
    >;

    /** user list */
    type UserList = Common.PaginatingQueryRecord<User>;

    /**
     * menu type
     *
     * - "1": directory
     * - "2": menu
     */
    type MenuType = '1' | '2';

    type MenuButton = {
      /**
       * button code
       *
       * it can be used to control the button permission
       */
      code: string;
      /** button description */
      desc: string;
    };

    /**
     * icon type
     *
     * - "1": iconify icon
     * - "2": local icon
     */
    type IconType = '1' | '2';
    type MenuPropsOfRoute = Pick<
      import('vue-router').RouteMeta,
      | 'i18nKey'
      | 'keepAlive'
      | 'constant'
      | 'order'
      | 'href'
      | 'hideInMenu'
      | 'activeMenu'
      | 'multiTab'
      | 'fixedIndexInTab'
      | 'query'
    >;
    type AddOrUpdateMenuParams = Pick<
      Api.SystemManage.Menu,
      | 'menuType'
      | 'menuName'
      | 'routeName'
      | 'routePath'
      | 'component'
      | 'order'
      | 'i18nKey'
      | 'icon'
      | 'iconType'
      | 'status'
      | 'parentId'
      | 'keepAlive'
      | 'constant'
      | 'href'
      | 'hideInMenu'
      | 'activeMenu'
      | 'multiTab'
      | 'fixedIndexInTab'
    > & {
      query: NonNullable<Api.SystemManage.Menu['query']>;
      layout: string;
      page: string;
      pathParam: string;
    };

    type Menu = Common.CommonRecord<{
      /** parent menu id */
      parentId: number;
      /** menu type */
      menuType: MenuType;
      /** menu name */
      menuName: string;
      /** route name */
      routeName: string;
      /** route path */
      routePath: string;
      /** component */
      component?: string;
      /** iconify icon name or local icon name */
      icon: string;
      /** icon type */
      iconType: IconType;
      /** buttons */
      buttons?: MenuButton[] | null;
      /** children menu */
      children?: Menu[] | null;
    }> &
      MenuPropsOfRoute;

    /** menu list */
    type MenuList = Common.PaginatingQueryRecord<Menu>;

    type MenuTree = {
      id: number;
      label: string;
      pId: number;
      children?: MenuTree[];
    };

    type BackApi = {
      code: string;
      name: string;
    };
  }

  namespace Dict {
    type LoggerTypes = 'info' | 'error';
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'> & {
      type: LoggerTypes;
    };
    type DictType = Common.CommonRecord<{
      name: string;
      code: string;
      description: string;
    }>;
    type DictTypeList = Common.PaginatingQueryRecord<DictType>;
    type AddOrUpdateDictTypeParams = Pick<DictType, 'code' | 'name' | 'description' | 'status'>;

    type DictTypeSearchParams = CommonType.RecordNullable<
      Pick<Api.Dict.DictType, 'name' | 'description'> & CommonSearchParams
    >;

    type DictData = Common.CommonRecord<{
      label: string;
      i18nKey: string;
      value: string;
      sort: number;
      code: string;
    }>;
    type DictDataSearchParams = CommonType.RecordNullable<Pick<Api.Dict.DictData, 'code'> & CommonSearchParams>;

    type DictDataList = Common.PaginatingQueryRecord<DictData>;
    type AddOrUpdateDictDataParams = Pick<DictData, 'label' | 'i18nKey' | 'value' | 'sort' | 'status' | 'code'>;

    type DictTypeDataSearchParams = CommonType.RecordNullable<
      Pick<Api.Dict.DictData, 'label' | 'i18nKey' | 'sort' | 'value'> & CommonSearchParams
    >;

    type Dict = {
      label: string;
      value: string;
      sort: number;
    };
  }
  namespace Logger {
    type CommonSearchParams = Pick<Common.PaginatingCommonParams, 'current' | 'size'>;

    type DecodeLoggerType = {
      time: string;
      level: string;
      source: { function: string; file: string; line: number };
      msg: string;
      cosTime: number;
      ip: string;
      operateBy: string;
      method: string;
      code: number;
      data: string;
    };
    type LoggerList = Common.PaginatingQueryRecord<string>;

    type LoggerSearchParams = CommonType.RecordNullable<
      Pick<Api.Dict.DictType, 'name' | 'description'> & CommonSearchParams
    >;

    type LoggerFiles = {
      fileName: string;
      fileSize: string;
      createData: string;
    };
  }
}
