<script lang="tsx" setup>
import { NButton, NCard, NDataTable, NPopconfirm, NTag } from 'naive-ui';
import dayjs from 'dayjs';

import { useTable, useTableOperate } from '@/hooks/common/table';
import { deleteDictDataByIds, fetchGetDictDataList } from '@/service/api';
import TableHeaderOperation from '@/components/advanced/table-header-operation.vue';
import { useAppStore } from '@/store/modules/app';
import { $t } from '@/locales';
import { enableStatusRecord } from '@/constants/business';

import DictDataOperateDrawer from './modules/dict-data-drawer.vue';
interface Props {
  id: string;
}

const { id } = defineProps<Props>();

const appStore = useAppStore();

const { loading, data, columns, getData, mobilePagination, columnChecks } = useTable({
  apiFn: fetchGetDictDataList,
  apiParams: {
    current: 1,
    size: 10,
    code: id
  },
  columns: () => [
    {
      type: 'selection',
      align: 'center',
      width: 48
    },
    {
      key: 'label',
      title: $t('page.dict.name'),
      align: 'left'
    },
    {
      key: 'value',
      title: $t('page.dict.code')
    },
    {
      key: 'sort',
      title: $t('page.dict.desc')
    },
    {
      key: 'status',
      title: $t('page.dict.status'),
      align: 'center',
      width: 100,
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
      key: 'createTime',
      title: $t('page.manage.updateTime'),
      align: 'center',
      render: row => {
        return row.updateTime !== null ? <span>{dayjs(row.updateTime).format('YYYY-MM-DD HH:mm:ss')}</span> : null;
      }
    },
    {
      key: 'operate',
      title: $t('common.operate'),
      align: 'center',
      width: 130,
      render: row => (
        <div class="flex-center gap-8px">
          <NButton type="primary" ghost size="small" onClick={() => edit(row.id)}>
            {$t('common.edit')}
          </NButton>
          <NPopconfirm onPositiveClick={() => handleDelete(row.id)}>
            {{
              default: () => $t('common.confirmDelete'),
              trigger: () => (
                <NButton type="error" ghost size="small">
                  {$t('common.delete')}
                </NButton>
              )
            }}
          </NPopconfirm>
        </div>
      )
    }
  ]
});
const {
  drawerVisible,
  operateType,
  editingData,
  handleAdd,
  handleEdit,
  checkedRowKeys,
  onBatchDeleted,
  onDeleted
  // closeDrawer
} = useTableOperate(data, getData);

async function handleBatchDelete() {
  const { error } = await deleteDictDataByIds(checkedRowKeys.value);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onBatchDeleted();
  }
}

async function handleDelete(id: number) {
  const { error } = await deleteDictDataByIds([id]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onDeleted();
  }
}

function edit(id: number) {
  handleEdit(id);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('page.dict.tableTitle')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @add="handleAdd"
          @delete="handleBatchDelete"
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
      ></NDataTable>
      <DictDataOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
    </NCard>
  </div>
</template>
