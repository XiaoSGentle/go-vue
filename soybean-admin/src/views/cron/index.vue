<script lang="tsx" setup>
import { NButton,NTag } from 'naive-ui';
import { h } from 'vue';
import { fetchGetSysCronList } from '@/service/api';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import CronOperateDrawer from './modules/cron-operate-drawer.vue';
import {enableStatusRecord, menuTypeRecord} from "@/constants/business";

const appStore = useAppStore();
const { columns, columnChecks, data, loading, getData, mobilePagination, searchParams, resetSearchParams } = useTable({
  apiFn: fetchGetSysCronList,
  apiParams: {
    current: 1,
    size: 10
  },
  columns: () => [
    {
      type: 'selection',
      align: 'center',
      width: 48
    },
    {
      key: 'key',
      title: $t('page.cron.key'),
      align: 'center',
      render: row => <span>{row.key}</span>
    },

    {
      key: 'schedule',
      title: $t('page.cron.schedule'),
      align: 'center',
      render: row => <span>{row.schedule}</span>
    },

    {
      key: 'status',
      title: $t('page.cron.status'),
      align: 'center',
      render: row => {
        if (row.status === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };
        const label = $t(enableStatusRecord[row.status]);
        return <NTag type={tagMap[row.status]}>{label}</NTag>;
      }
    },
    {
      key: 'description',
      title: $t('page.cron.description'),
      align: 'center',
      render: row => <span>{row.description}</span>
    },
    {
      key: 'operate',
      title: $t('common.operate'),
      align: 'center',
      width: 130,
      render: row => {
        const handleEdit = () => edit(row.key);
        return h('div', { class: 'flex-center gap-8px' }, [
          h(
            NButton,
            {
              type: 'primary',
              ghost: true,
              size: 'small',
              onClick: handleEdit
            },
            { default: () => $t('common.edit') }
          )
        ]);
      }
    }
  ]
});

const { drawerVisible, operateType, editingData, handleEdit, checkedRowKeys } = useTableOperate(data, getData);

function edit(id: string) {
  operateType.value = 'edit';
  editingData.value = data.value.find(item => item.key === id) || null;
  drawerVisible.value = true;
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">

    <NCard :bordered="false" :title="$t('page.cron.tableTitle')" class="sm:flex-1-hidden card-wrapper" size="small">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :disabled-delete="checkedRowKeys.length === 0"
          :extra-button="['delete', 'add']"
          :loading="loading"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
        :columns="columns"
        :data="data"
        :flex-height="!appStore.isMobile"
        :loading="loading"
        :pagination="mobilePagination"
        :row-key="row => row.key"
        :scroll-x="702"
        class="sm:h-full"
        remote
        size="small"
      />
      <CronOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
