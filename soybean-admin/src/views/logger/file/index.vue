<!-- eslint-disable vue/multi-word-component-names -->
<script lang="tsx" setup>
import { NCard } from 'naive-ui';

import { ref } from 'vue';
import dayjs from 'dayjs';
import { downLoggerGzFile } from '@/service/api';
import { useLoading } from '~/packages/hooks/src';
const { loading: dataLoading, startLoading, endLoading } = useLoading(false);
const fileList = ref<Api.Logger.LoggerFiles[] | null>([]);
const fetchLogFilesList = () => {
  startLoading();
  downLoggerGzFile().then(({ data }) => {
    data
      ?.sort((i, j) => {
        return i.createData > j.createData ? 1 : -1;
      })
      .forEach(item => fileList.value?.push(item));

    endLoading();
  });
};
fetchLogFilesList();
</script>

<template>
  <div class="flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard title="日志文件" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <NGrid x-gap="10" y-gap="30" :cols="3" responsive="screen">
        <template #header>
          <NSpace justify="space-between">
            <ReloadButton class="m3 mya" :loading="dataLoading" @click="fetchLogFilesList" />
          </NSpace>
        </template>
        <NGi
          v-for="(item, index) in fileList"
          :key="index"
          :class="!item.fileName.includes('error') ? 'text-gray-400' : 'text-red'"
        >
          <NSpace align="start">
            <div
              class="h-35 w-35 flex items-center justify-center border-1 rounded-lg"
              :class="!item.fileName.includes('error') ? 'border-gray-400' : 'border-red'"
            >
              <SvgIcon local-icon="folder-zip" class="h-15 w-15" />
            </div>
            <NDivider vertical class="h-40" />
            <div>
              <div class="my1">Name:{{ item.fileName }}</div>
              <div class="my1">Size:{{ item.fileSize }}</div>
              <div>Data:{{ dayjs(item.createData).format('YYYY-MM-DD HH:mm:ss') }}</div>
              <NButton :color="item.fileName.includes('error') ? '#f87171' : ''" class="my">下载</NButton>
            </div>
          </NSpace>
        </NGi>
      </NGrid>
    </NCard>
  </div>
</template>

<style scoped></style>
