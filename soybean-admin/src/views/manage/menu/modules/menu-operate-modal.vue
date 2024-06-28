<script lang="tsx" setup>
import { computed, reactive, ref, watch } from 'vue';
import type { SelectOption } from 'naive-ui';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions, menuIconTypeOptions, menuTypeOptions } from '@/constants/business';
import SvgIcon from '@/components/custom/svg-icon.vue';
import { getLocalIcons } from '@/utils/icon';
import { addMenu, fetchGetAllRoles, updateMenuById } from '@/service/api';
import {
  getLayoutAndPage,
  getPathParamFromRoutePath,
  getRoutePathByRouteName,
  getRoutePathWithParam,
  transformLayoutAndPageToComponent
} from './shared';

defineOptions({
  name: 'MenuOperateModal'
});

export type OperateType = NaiveUI.TableOperateType | 'addChild';

interface Props {
  /** the type of operation */
  operateType: OperateType;
  /** the edit menu data or the parent menu data when adding a child menu */
  rowData?: Api.SystemManage.Menu | null;
  /** all pages */
  allPages: string[];
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});

const { formRef, validate, restoreValidation } = useNaiveForm();
const { defaultRequiredRule } = useFormRules();

const title = computed(() => {
  const titles: Record<OperateType, string> = {
    add: $t('page.manage.menu.addMenu'),
    addChild: $t('page.manage.menu.addChildMenu'),
    edit: $t('page.manage.menu.editMenu')
  };
  return titles[props.operateType];
});

type Model = Pick<
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
  buttons: NonNullable<Api.SystemManage.Menu['buttons']>;
  layout: string;
  page: string;
  pathParam: string;
};

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    menuType: '1',
    menuName: '',
    routeName: '',
    routePath: '',
    pathParam: '',
    component: '',
    layout: '',
    page: '',
    i18nKey: null,
    icon: '',
    iconType: '1',
    parentId: 0,
    status: '1',
    keepAlive: false,
    constant: false,
    order: 0,
    href: null,
    hideInMenu: false,
    activeMenu: null,
    multiTab: false,
    fixedIndexInTab: null,
    query: [],
    buttons: []
  };
}

type RuleKey = Extract<keyof Model, 'menuName' | 'status' | 'routeName' | 'routePath'>;

const rules: Record<RuleKey, App.Global.FormRule> = {
  menuName: defaultRequiredRule,
  status: defaultRequiredRule,
  routeName: defaultRequiredRule,
  routePath: defaultRequiredRule
};

const disabledMenuType = computed(() => props.operateType === 'edit');

const localIcons = getLocalIcons();
const localIconOptions = localIcons.map<SelectOption>(item => ({
  label: () => (
    <div class="flex-y-center gap-16px">
      <SvgIcon localIcon={item} class="text-icon" />
      <span>{item}</span>
    </div>
  ),
  value: item
}));

const showLayout = computed(() => model.parentId === 0);

const showPage = computed(() => model.menuType === '2');

const pageOptions = computed(() => {
  const allPages = [...props.allPages];

  if (model.routeName && !allPages.includes(model.routeName)) {
    allPages.unshift(model.routeName);
  }

  const opts: CommonType.Option[] = allPages.map(page => ({
    label: page,
    value: page
  }));

  return opts;
});

const layoutOptions: CommonType.Option[] = [
  {
    label: 'base',
    value: 'base'
  },
  {
    label: 'blank',
    value: 'blank'
  },
  {
    label: 'project',
    value: 'project'
  }
];

/** the enabled role options */
const roleOptions = ref<CommonType.Option<string>[]>([]);

async function getRoleOptions() {
  const { error, data } = await fetchGetAllRoles();

  if (!error) {
    const options = data.map(item => ({
      label: item.roleName,
      value: item.roleCode
    }));

    roleOptions.value = [...options];
  }
}

function handleInitModel() {
  Object.assign(model, createDefaultModel());

  if (!props.rowData) return;

  if (props.operateType === 'addChild') {
    const { id } = props.rowData;

    Object.assign(model, { parentId: id });
  }

  if (props.operateType === 'edit') {
    const { component, ...rest } = props.rowData;

    const { layout, page } = getLayoutAndPage(component);
    const { path, param } = getPathParamFromRoutePath(rest.routePath);

    Object.assign(model, rest, { layout, page, routePath: path, pathParam: param });
  }

  if (!model.query) {
    model.query = [];
  }
  if (!model.buttons) {
    model.buttons = [];
  }
}

function closeDrawer() {
  visible.value = false;
}

function handleUpdateRoutePathByRouteName() {
  if (model.routeName) {
    model.routePath = getRoutePathByRouteName(model.routeName);
  } else {
    model.routePath = '';
  }
}

function handleUpdateI18nKeyByRouteName() {
  if (model.routeName) {
    model.i18nKey = `route.${model.routeName}` as App.I18n.I18nKey;
  } else {
    model.i18nKey = null;
  }
}

function handleCreateButton() {
  const buttonItem: Api.SystemManage.MenuButton = {
    code: '',
    desc: ''
  };

  return buttonItem;
}

function getSubmitParams() {
  const { layout, page, pathParam, ...params } = model;

  const component = transformLayoutAndPageToComponent(layout, page);
  const routePath = getRoutePathWithParam(model.routePath, pathParam);

  params.component = component;
  params.routePath = routePath;

  return params;
}

async function handleSubmit() {
  await validate();

  const params = getSubmitParams();

  // request
  if (props.operateType === 'add' || props.operateType === 'addChild') {
    const { error } = await addMenu(params);
    if (!error) {
      window.$message?.success($t('common.addSuccess'));
    }
  }

  if (props.operateType === 'edit') {
    params.query = [];
    const { error } = await updateMenuById(props.rowData?.id, params);
    if (!error) {
      window.$message?.success($t('common.updateSuccess'));
    }
  }

  closeDrawer();
  emit('submitted');
}

watch(visible, () => {
  if (visible.value) {
    handleInitModel();
    restoreValidation();
    getRoleOptions();
  }
});

watch(
  () => model.routeName,
  () => {
    handleUpdateRoutePathByRouteName();
    handleUpdateI18nKeyByRouteName();
  }
);
</script>

<template>
  <NModal v-model:show="visible" :title="title" class="w-800px" preset="card">
    <NScrollbar class="h-480px pr-20px">
      <NForm ref="formRef" :label-width="100" :model="model" :rules="rules" label-placement="left">
        <NGrid item-responsive responsive="screen">
          <NFormItemGi :label="$t('page.manage.menu.menuType')" path="menuType" span="24 m:12">
            <NRadioGroup v-model:value="model.menuType" :disabled="disabledMenuType">
              <NRadio v-for="item in menuTypeOptions" :key="item.value" :label="$t(item.label)" :value="item.value" />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.menuName')" path="menuName" span="24 m:12">
            <NInput v-model:value="model.menuName" :placeholder="$t('page.manage.menu.form.menuName')" />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.routeName')" path="routeName" span="24 m:12">
            <NInput v-model:value="model.routeName" :placeholder="$t('page.manage.menu.form.routeName')" />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.routePath')" path="routePath" span="24 m:12">
            <NInput v-model:value="model.routePath" :placeholder="$t('page.manage.menu.form.routePath')" disabled />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.pathParam')" path="pathParam" span="24 m:12">
            <NInput v-model:value="model.pathParam" :placeholder="$t('page.manage.menu.form.pathParam')" />
          </NFormItemGi>
          <NFormItemGi v-if="showLayout" :label="$t('page.manage.menu.layout')" path="layout" span="24 m:12">
            <NSelect
              v-model:value="model.layout"
              :options="layoutOptions"
              :placeholder="$t('page.manage.menu.form.layout')"
            />
          </NFormItemGi>
          <NFormItemGi v-if="showPage" :label="$t('page.manage.menu.page')" path="page" span="24 m:12">
            <NSelect
              v-model:value="model.page"
              :options="pageOptions"
              :placeholder="$t('page.manage.menu.form.page')"
            />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.i18nKey')" path="i18nKey" span="24 m:12">
            <NInput v-model:value="model.i18nKey" :placeholder="$t('page.manage.menu.form.i18nKey')" />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.order')" path="order" span="24 m:12">
            <NInputNumber v-model:value="model.order" :placeholder="$t('page.manage.menu.form.order')" class="w-full" />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.iconTypeTitle')" path="iconType" span="24 m:12">
            <NRadioGroup v-model:value="model.iconType">
              <NRadio
                v-for="item in menuIconTypeOptions"
                :key="item.value"
                :label="$t(item.label)"
                :value="item.value"
              />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.icon')" path="icon" span="24 m:12">
            <template v-if="model.iconType === '1'">
              <NInput v-model:value="model.icon" :placeholder="$t('page.manage.menu.form.icon')" class="flex-1">
                <template #suffix>
                  <SvgIcon v-if="model.icon" :icon="model.icon" class="text-icon" />
                </template>
              </NInput>
            </template>
            <template v-if="model.iconType === '2'">
              <NSelect
                v-model:value="model.icon"
                :options="localIconOptions"
                :placeholder="$t('page.manage.menu.form.localIcon')"
              />
            </template>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.menuStatus')" path="status" span="24 m:12">
            <NRadioGroup v-model:value="model.status">
              <NRadio
                v-for="item in enableStatusOptions"
                :key="item.value"
                :label="$t(item.label)"
                :value="item.value"
              />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.keepAlive')" path="keepAlive" span="24 m:12">
            <NRadioGroup v-model:value="model.keepAlive">
              <NRadio :label="$t('common.yesOrNo.yes')" :value="true" />
              <NRadio :label="$t('common.yesOrNo.no')" :value="false" />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.constant')" path="constant" span="24 m:12">
            <NRadioGroup v-model:value="model.constant">
              <NRadio :label="$t('common.yesOrNo.yes')" :value="true" />
              <NRadio :label="$t('common.yesOrNo.no')" :value="false" />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.href')" path="href" span="24 m:12">
            <NInput v-model:value="model.href" :placeholder="$t('page.manage.menu.form.href')" />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.hideInMenu')" path="hideInMenu" span="24 m:12">
            <NRadioGroup v-model:value="model.hideInMenu">
              <NRadio :label="$t('common.yesOrNo.yes')" :value="true" />
              <NRadio :label="$t('common.yesOrNo.no')" :value="false" />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi
            v-if="model.hideInMenu"
            :label="$t('page.manage.menu.activeMenu')"
            path="activeMenu"
            span="24 m:12"
          >
            <NSelect
              v-model:value="model.activeMenu"
              :options="pageOptions"
              :placeholder="$t('page.manage.menu.form.activeMenu')"
              clearable
            />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.multiTab')" path="multiTab" span="24 m:12">
            <NRadioGroup v-model:value="model.multiTab">
              <NRadio :label="$t('common.yesOrNo.yes')" :value="true" />
              <NRadio :label="$t('common.yesOrNo.no')" :value="false" />
            </NRadioGroup>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.fixedIndexInTab')" path="fixedIndexInTab" span="24 m:12">
            <NInputNumber
              v-model:value="model.fixedIndexInTab"
              :placeholder="$t('page.manage.menu.form.fixedIndexInTab')"
              class="w-full"
              clearable
            />
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.query')" span="24">
            <NDynamicInput
              v-model:value="model.query"
              :key-placeholder="$t('page.manage.menu.form.queryKey')"
              :value-placeholder="$t('page.manage.menu.form.queryValue')"
              preset="pair"
            >
              <template #action="{ index, create, remove }">
                <NSpace class="ml-12px">
                  <NButton size="medium" @click="() => create(index)">
                    <icon-ic:round-plus class="text-icon" />
                  </NButton>
                  <NButton size="medium" @click="() => remove(index)">
                    <icon-ic-round-remove class="text-icon" />
                  </NButton>
                </NSpace>
              </template>
            </NDynamicInput>
          </NFormItemGi>
          <NFormItemGi :label="$t('page.manage.menu.button')" span="24">
            <NDynamicInput v-model:value="model.buttons" :on-create="handleCreateButton">
              <template #default="{ value }">
                <div class="ml-8px flex-y-center flex-1 gap-12px">
                  <NInput
                    v-model:value="value.code"
                    :placeholder="$t('page.manage.menu.form.buttonCode')"
                    class="flex-1"
                  />
                  <NInput
                    v-model:value="value.desc"
                    :placeholder="$t('page.manage.menu.form.buttonDesc')"
                    class="flex-1"
                  />
                </div>
              </template>
              <template #action="{ index, create, remove }">
                <NSpace class="ml-12px">
                  <NButton size="medium" @click="() => create(index)">
                    <icon-ic:round-plus class="text-icon" />
                  </NButton>
                  <NButton size="medium" @click="() => remove(index)">
                    <icon-ic-round-remove class="text-icon" />
                  </NButton>
                </NSpace>
              </template>
            </NDynamicInput>
          </NFormItemGi>
        </NGrid>
      </NForm>
    </NScrollbar>
    <template #footer>
      <NSpace :size="16" justify="end">
        <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
        <NButton type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<style scoped></style>
