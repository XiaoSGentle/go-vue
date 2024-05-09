<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';
import { addDictData, updateDictDataById } from '@/service/api';

defineOptions({
  name: 'DictTypeOperateDrawer'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.Dict.DictData | null;

  /** the dict type code */
  typeCode: string;
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
    add: $t('page.dict.data.form.add'),
    edit: $t('page.dict.data.form.edit')
  };
  return titles[props.operateType];
});

type Model = Api.Dict.AddOrUpdateDictDataParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    label: '',
    value: '',
    sort: 0,
    code: '',
    enLabel: '',
    status: '2'
  };
}

type RuleKey = keyof Model & ('label' | 'value');
const rules: Record<RuleKey, App.Global.FormRule> = {
  label: defaultRequiredRule,
  value: defaultRequiredRule
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

  model.code = props.typeCode;
  if (props.operateType === 'add') {
    const { error } = await addDictData(model);
    if (!error) {
      window.$message?.success($t('common.addSuccess'));
      closeDrawer();
      emit('submitted');
    }
  }
  if (props.operateType === 'edit') {
    const { error } = await updateDictDataById(props.rowData?.id, model);
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
        <NFormItem :label="$t('page.dict.data.label')" path="label">
          <NInput v-model:value="model.label" :placeholder="$t('page.dict.data.form.label')" />
        </NFormItem>
        <NFormItem :label="$t('page.dict.data.value')" path="value">
          <NInput v-model:value="model.value" :placeholder="$t('page.dict.data.form.value')" />
        </NFormItem>

        <NFormItem :label="$t('page.dict.data.enLabel')">
          <NInput v-model:value="model.enLabel" :placeholder="$t('page.dict.data.form.enLabel')" />
        </NFormItem>

        <NFormItem :label="$t('page.dict.data.sort')">
          <NInputNumber
            v-model:value="model.sort"
            button-placement="right"
            class="w-full"
            :placeholder="$t('page.dict.data.form.sort')"
          />
        </NFormItem>
        <NFormItem :label="$t('page.dict.data.status')">
          <NRadioGroup v-model:value="model.status">
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
