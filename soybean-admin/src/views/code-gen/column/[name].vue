<!-- eslint-disable vue/multi-word-component-names -->
<script lang="tsx" setup>
import { NButton, NCard, NDataTable, NTag } from 'naive-ui';
import dayjs from 'dayjs';

import { computed } from 'vue';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { fetchTableColumns } from '@/service/api';
import TableHeaderOperation from '@/components/advanced/table-header-operation.vue';
import { useAppStore } from '@/store/modules/app';
import { $t } from '@/locales';
import { enableStatusRecord } from '@/constants/business';
import ColumnDataOperateDrawer from './modules/column-data-drawer.vue';
interface Props {
  name: string;
}

const { name: tableName } = defineProps<Props>();
const thisTableName = computed(() => tableName);
const appStore = useAppStore();

const { loading, data, columns, getData, mobilePagination, columnChecks } = useTable({
  apiFn: fetchTableColumns,
  apiParams: {
    current: 1,
    size: 10,
    tableName: thisTableName.value
  },
  columns: () => [
    {
      key: 'snakeCase',
      title: $t('page.gen.columnType.snakeCase')
    },

    {
      key: 'goType',
      title: $t('page.gen.columnType.goType'),
      align: 'center'
    },
    {
      key: 'tsType',
      title: $t('page.gen.columnType.tsType'),
      align: 'center'
    },
    {
      key: 'htmlType',
      title: $t('page.gen.columnType.htmlType'),
      align: 'center'
    },
    {
      key: 'required',
      width: 70,
      title: $t('page.gen.columnType.required'),
      align: 'center',
      render: row => {
        if (row.required === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.required]);

        return <NTag type={tagMap[row.required]}>{label}</NTag>;
      }
    },
    {
      key: 'isQuery',
      width: 70,
      title: $t('page.gen.columnType.isQuery'),
      align: 'center',
      render: row => {
        if (row.isQuery === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.isQuery]);

        return <NTag type={tagMap[row.isQuery]}>{label}</NTag>;
      }
    },
    {
      key: 'sort',
      title: $t('page.gen.columnType.sort'),
      align: 'center'
    },
    {
      key: 'length',
      title: $t('page.gen.columnType.length')
    },
    {
      key: 'comment',
      title: $t('page.gen.columnType.comment')
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
        </div>
      )
    }
  ]
});
const {
  drawerVisible,
  operateType,
  editingData,

  handleEdit,
  checkedRowKeys

  // closeDrawer
} = useTableOperate(data, getData);

function edit(id: number) {
  handleEdit(id);
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard
      :title="$t('page.dict.data.tableTitle')"
      :bordered="false"
      size="small"
      class="sm:flex-1-hidden card-wrapper"
    >
      <template #header-extra>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :extra-button="['add', 'delete']"
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
      ></NDataTable>
      <ColumnDataOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
    </NCard>
  </div>
</template>
