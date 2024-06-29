<script lang="tsx" setup>


import {LAYOUT_SCROLL_EL_ID, SimpleScrollbar} from "~/packages/materials/src/index.js";
import {useAppStore} from "@/store/modules/app/index.js";
import {useThemeStore} from "@/store/modules/theme/index.js";
import {useBoolean} from "~/packages/hooks/src/index.js";
import AdminLayout from "~/packages/materials/src/libs/admin-layout/index.vue";
import GlobalContent from "@/layouts/modules/global-content/index.vue";
import type {MenuOption} from "naive-ui";
import {computed} from "vue";
import {useRouteStore} from "@/store/modules/route";
import {getGlobalMenusByAuthRoutes} from "@/store/modules/route/shared";
import {useRouterPush} from "@/hooks/common/router";

defineOptions({
  name: 'ProjectTwoLayout'
});
const route = useRouterPush()
const appStore = useAppStore();
const themeStore = useThemeStore();
const {bool: menuLoading, setBool: setMenuLoading} = useBoolean()
const routeStore = useRouteStore();


// Menus
const menuOptions = computed(() => {
  setMenuLoading(true)
  const _projectMenus = routeStore.nativeRoutes.filter(route => route.name === 'projects');
  if (_projectMenus.length === 0 || _projectMenus[0].children === undefined) return []
  const menu = getGlobalMenusByAuthRoutes(_projectMenus[0].children) as unknown as MenuOption[];
  setMenuLoading(false)
  return menu
})

const handleMenuValueUpdate = (_: string, b: MenuOption) => {
  route.routerPush(b.routePath as string)
}

const {bool: isProjectCollapse, setBool: setProjectCollapse} = useBoolean(true);


</script>

<template>
  <AdminLayout
      v-model:sider-collapse="isProjectCollapse"
      :content-class="appStore.contentXScrollable ? 'overflow-x-hidden'  :  ''"
      :is-mobile="appStore.isMobile"
      :right-footer="themeStore.footer.right"
      :scroll-el-id="LAYOUT_SCROLL_EL_ID"
      :scroll-mode="themeStore.layout.scrollMode"
      :sider-collapsed-width="200"
      :sider-width="50"
  >
    <template #sider>
      <DarkModeContainer class="size-full flex-col-stretch shadow-sider">
        <SimpleScrollbar>
          <NSkeleton v-if="menuLoading" :repeat="10" :sharp="false" animated size="large" text/>
          <NMenu v-else :collapsed="!isProjectCollapse" :collapsed-icon-size="20"
                 :on-update:value="handleMenuValueUpdate"
                 :options="menuOptions"/>
        </SimpleScrollbar>
      </DarkModeContainer>
      <NButton
          v-if="!menuLoading"
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
