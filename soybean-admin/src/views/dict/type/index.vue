<script lang="tsx" setup>
import { NButton, NCard, NDataTable, NPopconfirm, NTag } from 'naive-ui';
import dayjs from 'dayjs';

import { useTable, useTableOperate } from '@/hooks/common/table';
import { deleteDictTypeByIds, fetchGetDictTypeList } from '@/service/api';
import TableHeaderOperation from '@/components/advanced/table-header-operation.vue';
import { useAppStore } from '@/store/modules/app';
import { $t } from '@/locales';
import { enableStatusRecord } from '@/constants/business';
import { useRouterPush } from '@/hooks/common/router';
import DictTypeOperateDrawer from './modules/dict-type-drawer.vue';

const appStore = useAppStore();
const { routerPush } = useRouterPush();
const { loading, data, columns, getData, mobilePagination, columnChecks } = useTable({
  apiFn: fetchGetDictTypeList,
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
      key: 'name',
      title: $t('page.dict.type.name'),
      align: 'center'
    },
    {
      key: 'code',
      title: $t('page.dict.type.code'),
      align: 'center',
      render: row => {
        return (
          <div
            onClick={() => {
              routerPush(`/dict/data/${row.code}`);
            }}
          >
            <span class="cursor-pointer border-blue-6 p-1 color-blue hover:border-b-1.5 hover:color-blue-6">
              {row.code}
            </span>
          </div>
        );
      }
    },
    {
      key: 'description',
      title: $t('page.dict.type.desc')
    },
    {
      key: 'status',
      title: $t('page.dict.type.status'),
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
  const { error } = await deleteDictTypeByIds(checkedRowKeys.value);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onBatchDeleted();
  }
}

async function handleDelete(id: number) {
  const { error } = await deleteDictTypeByIds([id]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onDeleted();
  }
}

function edit(id: number) {
  handleEdit(id);
}

// const ThisComponent = () => {
//   return (
//     <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
//       <NCard title={$t('page.dict.tableTitle')} bordered={false} size="small" class="sm:flex-1-hidden card-wrapper">
//         {{
//           'header-extra': () => {
//             return (
//               <TableHeaderOperation
//                 columns={columnChecks.value}
//                 loading={loading.value}
//                 disabledDelete={checkedRowKeys.value.length === 0}
//                 onAdd={handleAdd}
//                 onDelete={handleBatchDelete}
//                 onRefresh={getData}
//               />
//             );
//           },
//           default: () => {
//             return (
//               <div>
//                 <NDataTable
//                   class="sm:full"
//                   checkedRowKeys={checkedRowKeys.value}
//                   columns={columns.value}
//                   data={data.value}
//                   size="small"
//                   flexHeight={!appStore.isMobile}
//                   scrollX={705}
//                   loading={loading.value}
//                   remote
//                   rowKey={row => row.id}
//                   pagination={mobilePagination.value}
//                 />
//                 <DictTypeOperateDrawer
//                   visible={drawerVisible.value}
//                   operateType={operateType.value}
//                   rowData={editingData.value}
//                   onUpdate:visible={value => {
//                     drawerVisible.value = value;
//                   }}
//                   onSubmitted={getData}
//                 />
//               </div>
//             );
//           }
//         }}
//       </NCard>
//     </div>
//   );
// };
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard
      :title="$t('page.dict.type.tableTitle')"
      :bordered="false"
      size="small"
      class="sm:flex-1-hidden card-wrapper"
    >
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
      <DictTypeOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
    </NCard>
  </div>
</template>
