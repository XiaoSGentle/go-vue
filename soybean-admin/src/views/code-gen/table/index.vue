<script lang="tsx" setup>
import { NButton, NCard, NDataTable, NPopconfirm, NSpace, NTag } from 'naive-ui';
import dayjs from 'dayjs';

import { ref } from 'vue';
import { useTable, useTableOperate } from '@/hooks/common/table';
import { addTable, deleteTableByIds, fetchCanGenTables, fetchGenTables } from '@/service/api';
import TableHeaderOperation from '@/components/advanced/table-header-operation.vue';
import { useAppStore } from '@/store/modules/app';
import { $t } from '@/locales';
import { enableStatusRecord } from '@/constants/business';
import { useBoolean } from '~/packages/hooks/src';
import { useRouterPush } from '@/hooks/common/router';
import TableOperateDrawer from './modules/table-data-drawer.vue';
const appStore = useAppStore();
const { routerPush } = useRouterPush();
const { loading, data, columns, getData, mobilePagination, columnChecks } = useTable({
  apiFn: fetchGenTables,
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
      key: 'tableName',
      title: $t('page.gen.type.tableName'),

      render: row => {
        return (
          <NSpace>
            <div
              onClick={() => {
                routerPush(`/gen/column/${row.tableName}`);
              }}
              class="mt1"
            >
              <span class="cursor-pointer border-blue-6 p-1 color-blue hover:border-b-1.5 hover:color-blue-6">
                {row.tableName}
              </span>
            </div>
          </NSpace>
        );
      }
    },
    {
      key: 'relativePath',
      title: $t('page.gen.type.relativePath')
    },
    {
      key: 'checkToken',
      title: $t('page.gen.type.checkToken'),
      align: 'center',
      render: row => {
        if (row.checkToken === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.checkToken]);

        return <NTag type={tagMap[row.checkToken]}>{label}</NTag>;
      }
    },
    {
      key: 'checkAuth',
      title: $t('page.gen.type.checkAuth'),
      align: 'center',
      render: row => {
        if (row.checkAuth === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.checkAuth]);

        return <NTag type={tagMap[row.checkAuth]}>{label}</NTag>;
      }
    },
    {
      key: 'addLog',
      title: $t('page.gen.type.addLog'),
      align: 'center',
      render: row => {
        if (row.addLog === null) {
          return null;
        }
        const tagMap: Record<Api.Common.EnableStatus, NaiveUI.ThemeColor> = {
          1: 'success',
          2: 'warning'
        };

        const label = $t(enableStatusRecord[row.addLog]);

        return <NTag type={tagMap[row.addLog]}>{label}</NTag>;
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
  handleEdit,
  checkedRowKeys,
  onBatchDeleted,
  onDeleted
  // closeDrawer
} = useTableOperate(data, getData);

async function handleBatchDelete() {
  const { error } = await deleteTableByIds(checkedRowKeys.value);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onBatchDeleted();
  }
}

async function handleDelete(id: number) {
  const { error } = await deleteTableByIds([id]);
  if (!error) {
    window.$message?.success($t('common.deleteSuccess'));
    onDeleted();
  }
}

function edit(id: number) {
  handleEdit(id);
}

// 添加表信息

const canAddGenTables = ref<string[]>([]);
const selectGenTables = ref<string[]>([]);
const { bool: getLoading, setBool: setGetLoading } = useBoolean(false);
async function fetchCanGenTablesList() {
  setGetLoading(true);
  const { data: _data } = await fetchCanGenTables();
  if (_data) canAddGenTables.value = _data;
  setGetLoading(false);
}

const { bool: addTableModelVis, setBool: setAddTableModelVis } = useBoolean(false);

async function handleAddBtnClick() {
  setGetLoading(true);
  await fetchCanGenTablesList();
  setAddTableModelVis(true);
  setGetLoading(false);
}

const { bool: addLoading, setBool: setAddLoading } = useBoolean(false);
async function handleAddModelBtnClick() {
  setAddLoading(true);
  await addTable(selectGenTables.value);
  setAddLoading(false);
  setAddTableModelVis(false);
  getData();
}
</script>

<template>
  <div class="min-h-500px flex-col-stretch gap-16px overflow-hidden lt-sm:overflow-auto">
    <NCard :title="$t('page.gen.type.tableTitle')" :bordered="false" size="small" class="sm:flex-1-hidden card-wrapper">
      <template #header-extra>
        <NButton :loading="getLoading" size="small" ghost type="primary" class="mr" @click="handleAddBtnClick">
          <template #icon>
            <icon-ic-round-plus class="text-icon" />
          </template>
          {{ $t('common.add') }}
        </NButton>
        <TableHeaderOperation
          v-model:columns="columnChecks"
          :extra-button="['add']"
          :disabled-delete="checkedRowKeys.length === 0"
          :loading="loading"
          @delete="handleBatchDelete"
          @refresh="getData"
        ></TableHeaderOperation>
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
      <TableOperateDrawer
        v-model:visible="drawerVisible"
        :operate-type="operateType"
        :row-data="editingData"
        @submitted="getData"
      />
      <NDrawer v-model:show="addTableModelVis" display-directive="show" :width="300">
        <NDrawerContent :title="$t('page.gen.type.form.add')" closable>
          <NCheckboxGroup v-model:value="selectGenTables">
            <NSpace vertical>
              <NCheckbox
                v-for="(item, index) in canAddGenTables"
                :key="index"
                size="large"
                :label="item"
                :value="item"
              />
            </NSpace>
          </NCheckboxGroup>
          <template #footer>
            <NSpace :size="16" align="center">
              <NButton @click="setAddTableModelVis(false)">{{ $t('common.cancel') }}</NButton>
              <NButton :loading="addLoading" type="primary" @click="handleAddModelBtnClick">
                {{ $t('common.confirm') }}
              </NButton>
            </NSpace>
          </template>
        </NDrawerContent>
      </NDrawer>
    </NCard>
  </div>
</template>
