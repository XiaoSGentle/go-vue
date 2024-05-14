<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';
import { updateTableById } from '@/service/api';

defineOptions({
  name: 'TableOperateDrawer'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.CodeGen.TableInfoType | null;
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
    add: $t('page.gen.type.form.add'),
    edit: $t('page.gen.type.form.edit')
  };
  return titles[props.operateType];
});

type Model = Api.CodeGen.AddOrUpdateTableInfoDataParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    tableName: '',
    authorName: '',
    checkToken: '1',
    checkAuth: '1',
    addLog: '1',
    remarks: '',
    relativePath: ''
  };
}

type RuleKey = keyof Model & 'relativePath';
const rules: Record<RuleKey, App.Global.FormRule> = {
  relativePath: defaultRequiredRule
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
    const { error } = await updateTableById(props.rowData?.id, model);
    if (!error) {
      window.$message?.success($t('common.updateSuccess'));
      closeDrawer();
      emit('submitted');
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
        <NFormItem :label="$t('page.gen.type.tableName')">
          <NInput v-model:value="model.tableName" disabled />
        </NFormItem>
        <NFormItem :label="$t('page.gen.type.relativePath')" path="relativePath">
          <NInput v-model:value="model.relativePath" :placeholder="$t('page.gen.type.form.relativePath')" />
        </NFormItem>
        <NFormItem :label="$t('page.gen.type.authorName')">
          <NInput v-model:value="model.authorName" :placeholder="$t('page.gen.type.form.authorName')" />
        </NFormItem>

        <NFormItem :label="$t('page.gen.type.remarks')">
          <NInput v-model:value="model.remarks" :placeholder="$t('page.gen.type.form.remarks')" />
        </NFormItem>

        <NFormItem :label="$t('page.gen.type.checkAuth')">
          <NRadioGroup v-model:value="model.checkAuth">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.gen.type.checkToken')">
          <NRadioGroup v-model:value="model.checkToken">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
        </NFormItem>
        <NFormItem :label="$t('page.gen.type.addLog')">
          <NRadioGroup v-model:value="model.addLog">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :value="item.value" :label="$t(item.label)" />
          </NRadioGroup>
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
