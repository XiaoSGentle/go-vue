<script setup lang="ts">
import { ref, watch } from 'vue';
import hljs from 'highlight.js/lib/core';
import { useClipboard } from '@vueuse/core';
import { useBoolean } from '~/packages/hooks/src';

const { bool: modelVis } = useBoolean();
type Props = {
  show: boolean;
  tableUuid: string;
};
defineOptions({ name: 'GenCodeShow' });

const codeView = ref<Api.CodeGen.GenCodeItem[]>();

type Emits = {
  (e: 'modelClose'): () => void;
};
const emit = defineEmits<Emits>();

const handlerModelClose = () => {
  emit('modelClose');
};

const selectTableUuid = ref('');
const props = withDefaults(defineProps<Props>(), {
  show: false,
  tableUuid: ''
});

// 复制
const selectedTab = ref<string>('');
const handlerTabSelect = (selectTabName: string) => {
  selectedTab.value = selectTabName;
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
  () => props.show,
  async val => {
    if (val) {
      selectTableUuid.value = props.tableUuid;
      const { data } = await getTableGenView(selectTableUuid.value);
      if (data) {
        modelVis.value = val;
        codeView.value = data;
        selectedTab.value = data[0].fileName;
      } else {
        window.$message?.error('预览失败');
        emit('modelClose');
      }
    }
  }
);
</script>

<template>
  <div>
    <NModal
      v-model:show="modelVis"
      preset="card"
      :close-on-esc="false"
      :mask-closable="false"
      title="代码预览"
      class="w-900px"
      @update:show="handlerModelClose"
    >
      <NTabs type="card" :value="selectedTab" :on-update:value="handlerTabSelect" animated :bar-width="500">
        <NTabPane v-for="(item, index) in codeView" :key="index" :name="item.fileName" :tab="item.fileName">
          <NCode show-line-numbers :code="item.fileContent" :hljs="hljs" :language="item.lang" />
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
