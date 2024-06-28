<script lang="tsx" setup>


import {LAYOUT_SCROLL_EL_ID, SimpleScrollbar} from "~/packages/materials/src/index.js";
import {useAppStore} from "@/store/modules/app/index.js";
import {useThemeStore} from "@/store/modules/theme/index.js";
import {useBoolean} from "~/packages/hooks/src/index.js";
import AdminLayout from "~/packages/materials/src/libs/admin-layout/index.vue";
import GlobalContent from "@/layouts/modules/global-content/index.vue";
import {useSvgIcon} from "@/hooks/common/icon.js";
import type {MenuOption} from "naive-ui";
import {computed} from "vue";

defineOptions({
  name: 'ProjectTwoLayout'
});

const appStore = useAppStore();
const themeStore = useThemeStore();


const {SvgIconVNode} = useSvgIcon();


const menuOptions: MenuOption[] = [
  {
    label: '项目概览',
    key: 'project-a',
    icon: SvgIconVNode({icon: 'iconamoon:apps-fill'}),

  }, {
    label: '项目协同',
    key: 'project-c',
    icon: SvgIconVNode({icon: 'carbon:application'}),

  },
  {
    label: '代码仓库',
    key: 'bear-paw',
    icon: SvgIconVNode({icon: 'carbon:logo-gitlab'}),
  },
  {
    label: '文档管理',
    key: 'both',
    icon: SvgIconVNode({icon: 'hugeicons:folder-file-storage'}),
    children: [
      {
        label: '知识管理',
        key: 'can-not-1'
      }, {
        label: 'Wiki',
        key: 'can-not-2'
      },
      {
        label: '文件网盘',
        key: 'can-not-3'
      }, {
        label: 'API文档',
        key: 'can-not-4'
      },
    ]
  }, {
    label: '测试管理',
    key: 'both-1',
    icon: SvgIconVNode({icon: 'hugeicons:notebook-02'}),
    children: [
      {
        label: '测试概览',
        key: 'can-not-5'
      }, {
        label: '测试用例',
        key: 'can-not-6'
      }, {
        label: '测试计划',
        key: 'can-not-7'
      }, {
        label: '测试报告',
        key: 'can-not-8'
      },
    ]
  }
]
const {bool: isProjectCollapse, setBool: setProjectCollapse} = useBoolean(true);

</script>

<template>
  <AdminLayout
      v-model:sider-collapse="isProjectCollapse"
      :content-class="appStore.contentXScrollable ? 'overflow-x-hidden' : ''"
      :is-mobile="appStore.isMobile"
      :right-footer="themeStore.footer.right"
      :scroll-el-id="LAYOUT_SCROLL_EL_ID"
      :scroll-mode="themeStore.layout.scrollMode"
      :sider-collapsed-width="180"
      :sider-width="50"
  >
    <template #sider>
      <DarkModeContainer class="size-full flex-col-stretch shadow-sider">
        <SimpleScrollbar>
          <NMenu :collapsed="!isProjectCollapse" :collapsed-icon-size="20" :options="menuOptions"/>
        </SimpleScrollbar>
      </DarkModeContainer>
      <NButton
          circle
          class="red right--2.8 top-80% absolute z-100"
          secondary
          size="tiny"
          strong
          @click="setProjectCollapse(!isProjectCollapse)">
        <template #icon>
          <NIcon>
            <SvgIcon :icon="!isProjectCollapse?'mdi:chevron-double-right':'mdi:chevron-double-left'"/>
          </NIcon>
        </template>
      </NButton>
    </template>
    <GlobalContent/>
  </AdminLayout>
</template>
