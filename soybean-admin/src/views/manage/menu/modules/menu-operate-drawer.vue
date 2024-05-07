<script setup lang="tsx">
import { computed, reactive, watch } from 'vue';
import type { SelectOption } from 'naive-ui';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions, menuIconTypeOptions, menuTypeOptions } from '@/constants/business';
import SvgIcon from '@/components/custom/svg-icon.vue';
import { getLocalIcons } from '@/utils/icon';
import { addMenu, updateMenuById } from '@/service/api';

import locales from '@/locales/locale';
import { views } from '@/router/elegant/imports';
import { getLayoutAndPage, getRoutePathByRouteName, transformLayoutAndPageToComponent } from './shared';

defineOptions({
  name: 'MenuOperateDrawer'
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

type Model = Api.SystemManage.AddOrUpdateMenuParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    menuType: '1',
    menuName: '',
    routeName: '',
    routePath: '',
    pathParam: '',
    component: '',
    page: '',
    layout: 'base',
    i18nKey: undefined,
    icon: '',
    iconType: '1',
    parentId: 0,
    status: '1',
    keepAlive: false,
    constant: false,
    order: 0,
    href: undefined,
    hideInMenu: false,
    activeMenu: undefined,
    multiTab: false,
    fixedIndexInTab: null,
    query: []
  };
}
type RuleKey = Extract<keyof Model, 'menuName' | 'status' | 'routeName' | 'routePath' | 'i18nKey'>;

const i18nKeyMap = Object.entries(locales['zh-CN'].route).map(([key, _]) => {
  return { label: key, value: `route.${key}` };
});

const rules: Record<RuleKey, App.Global.FormRule> = {
  menuName: defaultRequiredRule,
  status: defaultRequiredRule,
  routeName: defaultRequiredRule,
  routePath: defaultRequiredRule,
  i18nKey: defaultRequiredRule
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
  }
];

function handleUpdateModel() {
  if (props.operateType === 'add') {
    Object.assign(model, createDefaultModel());

    return;
  }

  if (props.operateType === 'addChild' && props.rowData) {
    const { id } = props.rowData;

    Object.assign(model, createDefaultModel(), { parentId: id });
  }

  if (props.operateType === 'edit' && props.rowData) {
    const { component, ...rest } = props.rowData;

    const { layout, page } = getLayoutAndPage(component);

    Object.assign(model, rest, { layout, page });
  }
}

function closeDrawer() {
  visible.value = false;
}

const viewsName = Object.keys(views).map(view => {
  return {
    label: view,
    value: view
  };
});

async function handleSubmit() {
  await validate();

  model.component = transformLayoutAndPageToComponent(model.layout, model.page);

  // request
  if (props.operateType === 'add' || props.operateType === 'addChild') {
    const { error } = await addMenu(model);
    if (!error) {
      window.$message?.success($t('common.addSuccess'));
      closeDrawer();
      emit('submitted');
    }
  }

  if (props.operateType === 'edit') {
    const { error } = await updateMenuById(props.rowData?.id, model);
    if (!error) {
      window.$message?.success($t('common.updateSuccess'));
      closeDrawer();
      emit('submitted');
    }
  }
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
watch(visible, () => {
  if (visible.value) {
    handleUpdateModel();
    restoreValidation();
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
  <NDrawer v-model:show="visible" display-directive="show" :width="500">
    <NDrawerContent :title="title" :native-scrollbar="false" closable>
      <NForm ref="formRef" :model="model" :rules="rules" label-placement="top" :label-width="80">
        <NFormItem :label="$t('page.manage.menu.menuType')" path="menuType">
          <NRadioGroup v-model:value="model.menuType" :disabled="disabledMenuType">
            <NRadio v-for="item in menuTypeOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NDivider>{{ $t('page.manage.menu.form.mainConfig') }}</NDivider>
        <NFormItem :label="$t('page.manage.menu.menuName')" path="menuName">
          <NInput v-model:value="model.menuName" :placeholder="$t('page.manage.menu.form.menuName')" />
        </NFormItem>

        <NFormItem :label="$t('page.manage.menu.routeName')" path="routeName">
          <NSelect
            v-model:value="model.routeName"
            :options="viewsName"
            :placeholder="$t('page.manage.menu.form.routeName')"
          />
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.routePath')" path="routePath">
          <NInput v-model:value="model.routePath" :placeholder="$t('page.manage.menu.form.routePath')" />
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.i18nKey')" path="i18nKeyMap">
          <NSelect
            v-model:value="model.i18nKey"
            :options="i18nKeyMap"
            :placeholder="$t('page.manage.menu.form.i18nKey')"
          />
        </NFormItem>

        <NFormItem v-if="showPage" :label="$t('page.manage.menu.page')" path="page">
          <NSelect v-model:value="model.page" :options="pageOptions" :placeholder="$t('page.manage.menu.form.page')" />
        </NFormItem>
        <NFormItem v-if="showPage" :label="$t('page.manage.menu.hideInMenu')" path="hideInMenu">
          <NRadioGroup v-model:value="model.hideInMenu">
            <!-- eslint-disable-next-line vue/prefer-true-attribute-shorthand -->
            <NRadio :value="true" :label="$t('common.yesOrNo.yes')" />
            <NRadio :value="false" :label="$t('common.yesOrNo.no')" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem v-if="showLayout" :label="$t('page.manage.menu.layout')" path="layout">
          <NSelect
            v-model:value="model.layout"
            :options="layoutOptions"
            :placeholder="$t('page.manage.menu.form.layout')"
          />
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.order')" path="order">
          <NInputNumber v-model:value="model.order" class="w-full" :placeholder="$t('page.manage.menu.form.order')" />
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.menuStatus')" path="status">
          <NRadioGroup v-model:value="model.status">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NDivider>{{ $t('page.manage.menu.icon') }}</NDivider>
        <NFormItem :label="$t('page.manage.menu.iconTypeTitle')" path="iconType">
          <NRadioGroup v-model:value="model.iconType">
            <NRadio v-for="item in menuIconTypeOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>

        <NFormItem :label="$t('page.manage.menu.icon')" path="icon">
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
              :placeholder="$t('page.manage.menu.form.localIcon')"
              :options="localIconOptions"
            />
          </template>
        </NFormItem>
        <NDivider>{{ $t('page.manage.menu.form.otherConfig') }}</NDivider>
        <NFormItem :label="$t('page.manage.menu.keepAlive')" path="keepAlive">
          <NRadioGroup v-model:value="model.keepAlive">
            <NRadio value :label="$t('common.yesOrNo.yes')" />
            <NRadio :value="false" :label="$t('common.yesOrNo.no')" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.constant')" path="constant">
          <NRadioGroup v-model:value="model.constant">
            <!--eslint-disable-next-line vue/prefer-true-attribute-shorthand-->
            <NRadio :value="true" :label="$t('common.yesOrNo.yes')" />
            <NRadio :value="false" :label="$t('common.yesOrNo.no')" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.multiTab')" path="multiTab">
          <NRadioGroup v-model:value="model.multiTab">
            <NRadio value :label="$t('common.yesOrNo.yes')" />
            <NRadio :value="false" :label="$t('common.yesOrNo.no')" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.fixedIndexInTab')" path="fixedIndexInTab">
          <NInputNumber
            v-model:value="model.fixedIndexInTab"
            class="w-full"
            clearable
            :placeholder="$t('page.manage.menu.form.fixedIndexInTab')"
          />
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.query')">
          <NDynamicInput
            v-model:value="model.query"
            preset="pair"
            :key-placeholder="$t('page.manage.menu.form.queryKey')"
            :value-placeholder="$t('page.manage.menu.form.queryValue')"
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
        </NFormItem>
        <NFormItem :label="$t('page.manage.menu.href')" path="href">
          <NInput v-model:value="model.href" :placeholder="$t('page.manage.menu.form.href')" />
        </NFormItem>
      </NForm>
      <template #footer>
        <NSpace :size="16">
          <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
          <NButton type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
        </NSpace>
      </template>
    </NDrawerContent>
  </NDrawer>
</template>

<style scoped></style>
