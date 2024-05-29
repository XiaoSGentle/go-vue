<script setup lang="ts">
import { ref, watch } from 'vue';

import { useClipboard } from '@vueuse/core';
import hljs from 'highlight.js/lib/core';

import typescript from 'highlight.js/lib/languages/typescript';
import golang from 'highlight.js/lib/languages/go';
import html from 'highlight.js/lib/languages/haml';
import { fetchTableGenPreview } from '@/service/api';
import { useLoading } from '~/packages/hooks/src';
defineOptions({ name: 'GenCodeShow' });
hljs.registerLanguage('typescript', typescript);
hljs.registerLanguage('golang', golang);
hljs.registerLanguage('html', html);
const codeView = ref<Api.CodeGen.GenCodeItem[]>();

type Emits = {
  (e: 'modelClose'): () => void;
};
const emit = defineEmits<Emits>();
const tableName = defineModel<string>('tableName', {
  default: ''
});
const { loading: dataLoading, startLoading, endLoading } = useLoading(false);
// 复制
const selectedTab = ref<string>('');
const handlerTabSelect = (selectTabName: string) => {
  selectedTab.value = selectTabName;
};
async function fetchCodePreview() {
  startLoading();
  const { data } = await fetchTableGenPreview(tableName.value);
  if (data) {
    codeView.value = data;
    selectedTab.value = data[0].fileName;
  }
  endLoading();
}

const visible = defineModel<boolean>('visible', {
  default: false
});

const handlerModelClose = () => {
  emit('modelClose');
};

const { copy, isSupported } = useClipboard();
const handlerCopyButtonClick = () => {
  if (!isSupported) {
    window.$message?.warning('您的浏览器不支持,请更换最新版谷歌或者火狐浏览器再次尝试');
  } else {
    const table = codeView.value?.find(item => {
      return item.fileName === selectedTab.value;
    });
    if (table?.fileContent) copy(table?.fileContent);
    window.$message?.success('复制成功');
  }
};

watch(
  () => visible.value,
  val => val && fetchCodePreview()
);
</script>

<template>
  <div>
    <NModal
      v-model:show="visible"
      :loading="dataLoading"
      preset="card"
      :close-on-esc="false"
      :mask-closable="false"
      title="代码预览"
      class="w-900px"
      @update:show="handlerModelClose"
    >
      <NTabs type="card" :value="selectedTab" :on-update:value="handlerTabSelect" animated :bar-width="500">
        <NTabPane v-for="(item, index) in codeView" :key="index" :name="item.fileName" :tab="item.fileName">
          <NScrollbar>
            <NScrollbar class="h-60vh">
              <NCode word-wrap :code="item.fileContent" :hljs="hljs" :language="item.lang" />
            </NScrollbar>
          </NScrollbar>
        </NTabPane>
      </NTabs>
      <template #footer>
        <NSpace justify="center">
          <NButton @click="handlerCopyButtonClick">复制</NButton>
        </NSpace>
      </template>
    </NModal>
  </div>
</template>

<style lang="scss" scoped>
.n-code {
  height: 60vh;
  overflow: auto;
}
</style>
