<script lang="tsx" setup>
import { NCard, NCode } from 'naive-ui';

import { ref } from 'vue';
import { fetchLoggerList } from '@/service/api';

const loggerInfos = ref<string[]>([]);
const fetchParams = ref<Api.Dict.CommonSearchParams>({ size: 100, current: 1, type: 'info' });
const fetchLoggerContent = async () => {
  const { data } = await fetchLoggerList(fetchParams.value);
  data?.records.forEach(item => {
    loggerInfos.value.push(item);
  });
};

fetchLoggerContent();

const ThisComponent = () => {
  return (
    <NCard>
      <NCode code={loggerInfos.value.toString()} language={'json'} wordWrap={true} inline={true}></NCode>
    </NCard>
  );
};
</script>

<template>
  <ThisComponent />
</template>
