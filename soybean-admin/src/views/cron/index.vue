<script setup lang="tsx">
import { NButton } from 'naive-ui';
import { h } from 'vue';
import { fetchGetSysCronList } from '@/service/api';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import OperateDrawer from './modules/cron-operate-drawer.vue';

const appStore = useAppStore();
const { columns, columnChecks, data, loading, getData, mobilePagination, searchParams, resetSearchParams } = useTable({
  apiFn: fetchGetSysCronList,
  apiParams: {
    current: 1,
    size: 1
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
      render: row => <span>{row.status}</span>
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
        const handleEdit = () => edit(row.id);
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

function edit(id: number) {
  handleEdit(id);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <RoleSearch v-model:model="searchParams" @reset="resetSearchParams" @search="getData" />
    <NCard :title="$t('page.cron.tableTitle')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :extra-button="['delete', 'add']"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @refresh="getData"
        />
      </template>
      <NDataTable
        v-model:checked-row-keys="checkedRowKeys"
        :columns="columns"
        :data="data"
        size="small"
        :flex-height="!appStore.isMobile"
        :scroll-x="702"
        :loading="loading"
        remote
        :row-key="row => row.id"
        :pagination="mobilePagination"
        class="sm:h-full"
      />
      <OperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
    </NCard>
  </div>
</template>

<style scoped></style>
