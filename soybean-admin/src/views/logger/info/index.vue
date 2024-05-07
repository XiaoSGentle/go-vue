<script lang="tsx" setup>
import { NCard } from 'naive-ui';

import { onMounted, reactive, ref, watch } from 'vue';
import hljs from 'highlight.js/lib/core';
import json from 'highlight.js/lib/languages/json';
import { fetchLoggerList } from '@/service/api';
import { useLoading } from '~/packages/hooks/src';
hljs.registerLanguage('json', json);
const { loading: dataLoading, startLoading, endLoading } = useLoading(false);
const loggerInfos = ref<string[]>([]);
const fetchParams = reactive<Api.Dict.CommonSearchParams>({ size: 30, current: 1, type: 'info' });
const fetchLoggerContent = () => {
  startLoading();
  fetchLoggerList(fetchParams).then(({ data }) => {
    data?.records?.forEach(item => {
      loggerInfos.value.push(item);
    });
    endLoading();
  });
};

watch(
  () => fetchParams.current,
  () => {
    fetchLoggerContent();
  }
);
watch(
  () => fetchParams.type,
  () => {
    loggerInfos.value.length = 0;
    fetchParams.current = 1;
    fetchLoggerContent();
  }
);

const refreshLog = () => {
  loggerInfos.value.length = 0;
  fetchParams.current = 1;
  fetchLoggerContent();
};

onMounted(() => {
  fetchLoggerContent();
});
</script>

<template>
  <div class="flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header>
        <NSpace justify="space-between">
          <NRadioGroup v-model:value="fetchParams.type">
            <NRadioButton key="Info" value="info" label="Info" />
            <NRadioButton key="Error" value="error" label="Error" />
          </NRadioGroup>
          <ReloadButton class="m3 mya" :loading="dataLoading" @click="refreshLog" />
        </NSpace>
      </template>
      <!-- <NSkeleton v-if="dataLoading" height="18" size="medium" round :repeat="33" text /> -->
      <NLog
        :hljs="hljs"
        :rows="37"
        :lines="loggerInfos"
        trim
        :loading="dataLoading"
        language="json"
        :on-reach-bottom="() => fetchParams.current++"
      />
    </NCard>
  </div>
</template>
