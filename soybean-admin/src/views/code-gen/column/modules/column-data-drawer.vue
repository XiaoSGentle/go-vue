<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';
import { updateTableColumnsById } from '@/service/api';

defineOptions({
  name: 'ColumnDataOperateDrawer'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.CodeGen.TableColumnInfoType | null;
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});

const { formRef, validate, restoreValidation } = useNaiveForm();
const { defaultRequiredRule } = useFormRules();

const title = computed(() => {
  const titles: Record<NaiveUI.TableOperateType, string> = {
    add: $t('page.gen.columnType.form.add'),
    edit: $t('page.gen.columnType.form.edit', { columnName: props.rowData?.comment })
  };
  return titles[props.operateType];
});

type Model = Api.CodeGen.AddOrUpdateTableColumnsInfoDataParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    snakeCase: '',
    comment: '',
    dictType: '',
    goType: '',
    htmlType: '',
    isQuery: '1',
    queryType: '',
    isShow: '1',
    required: '1',
    sort: 0,
    tsType: ''
  };
}

type RuleKey = keyof Model & ('snakeCase' | 'value');
const rules: Record<RuleKey, App.Global.FormRule> = {
  snakeCase: defaultRequiredRule
};

function handleUpdateModelWhenEdit() {
  if (props.operateType === 'add') {
    Object.assign(model, createDefaultModel());
    return;
  }

  if (props.operateType === 'edit' && props.rowData) {
    Object.assign(model, props.rowData);
  }
}

function closeDrawer() {
  visible.value = false;
}

async function handleSubmit() {
  await validate();

  if (props.operateType === 'edit') {
    const { error } = await updateTableColumnsById(props.rowData?.id, model);
    if (!error) {
      window.$message?.success($t('common.updateSuccess'));
      emit('submitted');
      closeDrawer();
    }
  }
}

watch(visible, () => {
  if (visible.value) {
    handleUpdateModelWhenEdit();
    restoreValidation();
  }
});
</script>

<template>
  <NDrawer v-model:show="visible" display-directive="show" :width="360">
    <NDrawerContent :title="title" :native-scrollbar="false" closable>
      <NForm ref="formRef" :model="model" :rules="rules">
        <NFormItem :label="$t('page.dict.data.label')" path="snakeCase">
          <NInput v-model:value="model.snakeCase" disabled :placeholder="$t('page.dict.data.form.label')" />
        </NFormItem>

        <NFormItem :label="$t('page.gen.columnType.htmlType')">
          <SysDict v-model:select-value="model.htmlType" dict-key="SYS_GEN_HTML_TYPE" />
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.tsType')">
          <SysDict v-model:select-value="model.tsType" dict-key="SYS_GEN_TS_TYPE" />
        </NFormItem>

        <NFormItem :label="$t('page.gen.columnType.goType')">
          <SysDict v-model:select-value="model.goType" dict-key="SYS_GEN_GO_TYPE" />
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.queryType')">
          <SysDict v-model:select-value="model.queryType" dict-key="SYS_GEN_QUERY_TYPE" />
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.sort')">
          <NInputNumber
            v-model:value="model.sort"
            button-placement="right"
            class="w-full"
            :placeholder="$t('page.gen.columnType.form.sort')"
          />
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.required')">
          <NRadioGroup v-model:value="model.required">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>

        <NFormItem :label="$t('page.gen.columnType.isQuery')">
          <NRadioGroup v-model:value="model.isQuery">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.isShow')">
          <NRadioGroup v-model:value="model.isShow">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.gen.columnType.comment')">
          <NInput v-model:value="model.comment" type="textarea" />
        </NFormItem>
      </NForm>
      <template #footer>
        <NSpace :size="16" align="center">
          <NButton @click="closeDrawer">{{ $t('common.cancel') }}</NButton>
          <NButton type="primary" @click="handleSubmit">{{ $t('common.confirm') }}</NButton>
        </NSpace>
      </template>
    </NDrawerContent>
  </NDrawer>
</template>

<style scoped></style>
