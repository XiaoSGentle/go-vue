<script lang="ts" setup>
import {computed} from 'vue';
import {AdminLayout, LAYOUT_SCROLL_EL_ID} from '@sa/materials';
import type {LayoutMode} from '@sa/materials';
import {useAppStore} from '@/store/modules/app';
import {useThemeStore} from '@/store/modules/theme';
import GlobalHeader from '../modules/global-header/index.vue';
import GlobalSider from '../modules/global-sider/index.vue';
import GlobalTab from '../modules/global-tab/index.vue';
import GlobalContent from '../modules/global-content/index.vue';
import GlobalFooter from '../modules/global-footer/index.vue';
import ThemeDrawer from '../modules/theme-drawer/index.vue';
import ProjectLeftMenu from "@/layouts/project-layout/global-menu/base-menu.vue";
import {setupMixMenuContext} from '../context';
import ProjectHeader from "@/layouts/project-layout/global-header/index.vue";
import {useBoolean} from "~/packages/hooks";
import ProjectTwoLayout from "@/layouts/project-layout/global-project-layout/two-layout.vue";

defineOptions({
  name: 'ProjectLayout'
});

const appStore = useAppStore();
const themeStore = useThemeStore();
const {menus} = setupMixMenuContext();

const layoutMode = computed(() => {
  const vertical: LayoutMode = 'vertical';
  const horizontal: LayoutMode = 'horizontal';
  return themeStore.layout.mode.includes(vertical) ? vertical : horizontal;
});

const headerPropsConfig: Record<UnionKey.ThemeLayoutMode, App.Global.HeaderProps> = {
  vertical: {
    showLogo: false,
    showMenu: false,
    showMenuToggler: true
  },
  'vertical-mix': {
    showLogo: false,
    showMenu: false,
    showMenuToggler: false
  },
  horizontal: {
    showLogo: true,
    showMenu: true,
    showMenuToggler: false
  },
  'horizontal-mix': {
    showLogo: true,
    showMenu: true,
    showMenuToggler: false
  }
};

const headerProps = computed(() => headerPropsConfig[themeStore.layout.mode]);

const siderVisible = computed(() => themeStore.layout.mode !== 'horizontal');

const isVerticalMix = computed(() => themeStore.layout.mode === 'vertical-mix');

const isHorizontalMix = computed(() => themeStore.layout.mode === 'horizontal-mix');

const siderWidth = computed(() => getSiderWidth());

const siderCollapsedWidth = computed(() => getSiderCollapsedWidth());

function getSiderWidth() {
  const {width, mixWidth, mixChildMenuWidth} = themeStore.sider;

  let w = isVerticalMix.value || isHorizontalMix.value ? mixWidth : width;

  if (isVerticalMix.value && appStore.mixSiderFixed && menus.value.length) {
    w += mixChildMenuWidth;
  }

  return w;
}

function getSiderCollapsedWidth() {
  const {collapsedWidth, mixCollapsedWidth, mixChildMenuWidth} = themeStore.sider;

  let w = isVerticalMix.value || isHorizontalMix.value ? mixCollapsedWidth : collapsedWidth;

  if (isVerticalMix.value && appStore.mixSiderFixed && menus.value.length) {
    w += mixChildMenuWidth;
  }

  return w;
}

const {bool: isProjectCollapse, setBool: setProjectCollapse} = useBoolean(true)
const projectLeftMenuWidth = computed(() => isProjectCollapse.value ? 240 : 80);
</script>

<template>
  <AdminLayout
      v-model:sider-collapse="appStore.siderCollapse"
      :content-class="appStore.contentXScrollable ? 'overflow-x-hidden' : ''"
      :fixed-footer="themeStore.footer.fixed"
      :fixed-top="themeStore.fixedHeaderAndTab"
      :footer-height="themeStore.footer.height"
      :footer-visible="themeStore.footer.visible"
      :full-content="appStore.fullContent"
      :header-height="themeStore.header.height"
      :is-mobile="appStore.isMobile"
      :mode="layoutMode"
      :right-footer="themeStore.footer.right"
      :scroll-el-id="LAYOUT_SCROLL_EL_ID"
      :scroll-mode="themeStore.layout.scrollMode"
      :sider-collapsed-width="siderCollapsedWidth"
      :sider-visible="siderVisible"
      :sider-width="siderWidth"
      :tab-height="themeStore.tab.height"
      :tab-visible="themeStore.tab.visible"
  >
    <template #header>
      <GlobalHeader v-bind="headerProps"/>
    </template>
    <template #tab>
      <GlobalTab/>
    </template>
    <template #sider>
      <GlobalSider/>
    </template>
    <ProjectTwoLayout class="z-1"/>
    <ThemeDrawer/>
    <template #footer>
      <GlobalFooter/>
    </template>
  </AdminLayout>
</template>

<style lang="scss">
#__SCROLL_EL_ID__ {
  @include scrollbar();
}

</style>
