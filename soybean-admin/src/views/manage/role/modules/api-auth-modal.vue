<script setup lang="tsx">
import { computed, ref, shallowRef } from 'vue';
import { NButton, NCheckbox, NCheckboxGroup, NDivider, NGi, NGrid, NModal, NScrollbar, NSpace, NTag } from 'naive-ui';
import { $t } from '@/locales';
import { fetchGetAllApis, getRolePermitByCode, updateRoleApiCodesPermitByCode } from '@/service/api';

defineOptions({
  name: 'ApiAuthModal'
});

interface Props {
  roleCode: string | undefined;
}

const props = defineProps<Props>();

type ApisGroup = {
  title: string;
  apis: Api.SystemManage.BackApi[];
};

const visible = defineModel<boolean>('visible', {
  default: false
});

function closeModal() {
  visible.value = false;
}

const title = computed(() => $t('common.edit') + $t('page.manage.role.apiAuth'));

const tree = shallowRef<Api.SystemManage.BackApi[]>([]);

async function getAllApis() {
  const { data } = await fetchGetAllApis();
  if (data) tree.value = data;
}

const checks = ref<(string | number)[]>([]);

async function getChecks() {
  const { data } = await getRolePermitByCode(props.roleCode);
  if (data) checks.value = data?.apiCodes;
}

async function handleSubmit() {
  const { error } = await updateRoleApiCodesPermitByCode({ roleCode: props.roleCode, apiCodes: checks.value });
  if (!error) {
    window.$message?.success?.($t('common.modifySuccess'));
    closeModal();
  }
}

const apisGroups = computed(() => {
  const regex = /\/[^/]+\/([^/]+)/;
  const groups: ApisGroup[] = [];
  tree.value.forEach(item => {
    const match = item.code.match(regex);
    if (match) {
      const findGroup = groups.find(_group => _group.title === match[1]);
      if (findGroup) {
        findGroup.apis.push(item);
      } else {
        groups.push({
          title: match[1],
          apis: [item]
        });
      }
    }
  });
  return groups;
});

function checkIsSelectAll(param: Api.SystemManage.BackApi[]): boolean {
  for (const item of param) {
    if (checks.value.includes(item.code)) {
      return false;
    }
  }
  return true;
}
function init() {
  getAllApis();
  getChecks();
}

// init
init();
const tagMap: Record<string, NaiveUI.ThemeColor> = {
  GET: 'success',
  POST: 'warning',
  PUT: 'info',
  DELETE: 'error'
};
const Render = () => (
  <NModal show={visible.value} title={title.value} preset="card" class="w-700px" onClose={closeModal}>
    {{
      default: () => (
        <NScrollbar class="h-450px">
          <NCheckboxGroup value={checks.value} onUpdateValue={value => (checks.value = value)}>
            {apisGroups.value.map(_group => {
              return (
                <>
                  <div class="my-a mt h-50px flex">
                    <NDivider titlePlacement="left">
                      <div class="flex">
                        <NButton
                          size="small"
                          type="primary"
                          class="mr"
                          onClick={() => {
                            const selectedCodes = _group.apis.map(item => item.code);
                            checks.value = checkIsSelectAll(_group.apis)
                              ? [...new Set([...checks.value, ...selectedCodes])]
                              : checks.value.filter(code => !selectedCodes.includes(code.toString()));
                          }}
                        >
                          {$t(
                            checkIsSelectAll(_group.apis)
                              ? 'page.manage.role.apiEdit.selectAll'
                              : 'page.manage.role.apiEdit.selectNone'
                          )}
                        </NButton>
                        <div class="text-16px font-bold"> {_group.title}</div>
                      </div>
                    </NDivider>
                  </div>
                  <NSpace vertical>
                    <NGrid yGap={8} cols={2}>
                      {_group.apis.map(api => {
                        return (
                          <NGi>
                            <NCheckbox key={api.code} value={api.name}>
                              <div class="flex">
                                <div class="w-70px flex justify-center">
                                  <NTag type={tagMap[api.name.split('::')[0]]}>{api.name.split('::')[0]}</NTag>
                                </div>
                                <span>{api.name.split('::')[1]}</span>
                              </div>
                            </NCheckbox>
                          </NGi>
                        );
                      })}
                    </NGrid>
                  </NSpace>
                </>
              );
            })}
          </NCheckboxGroup>
        </NScrollbar>
      ),
      footer: () => (
        <NSpace justify="center" align="center">
          <NButton size="small" class="mt-16px" onClick={closeModal}>
            {$t('common.cancel')}
          </NButton>
          <NButton type="primary" size="small" onClick={handleSubmit} class="mt-16px">
            {$t('common.confirm')}
          </NButton>
        </NSpace>
      )
    }}
  </NModal>
);
</script>

<template>
  <Render />
</template>

<style scoped></style>
