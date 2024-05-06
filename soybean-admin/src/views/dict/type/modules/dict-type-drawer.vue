<script setup lang="ts">
import { computed, reactive, watch } from 'vue';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';
import { addDictType, updateDictTypeById } from '@/service/api';

defineOptions({
  name: 'DictTypeOperateDrawer'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.Dict.DictType | null;
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
    add: $t('page.dict.type.form.add'),
    edit: $t('page.dict.type.form.edit')
  };
  return titles[props.operateType];
});

type Model = Api.Dict.AddOrUpdateDictTypeParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    name: '',
    code: '',
    description: '',
    status: '2'
  };
}

type RuleKey = keyof Model & ('name' | 'code');
const rules: Record<RuleKey, App.Global.FormRule> = {
  name: defaultRequiredRule,
  code: defaultRequiredRule
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
  if (props.operateType === 'add') {
    const { error } = await addDictType(model);
    if (!error) {
      window.$message?.success($t('common.addSuccess'));
      closeDrawer();
      emit('submitted');
    }
  }
  if (props.operateType === 'edit') {
    const { error } = await updateDictTypeById(props.rowData?.id, model);
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
        <NFormItem :label="$t('page.dict.type.name')" path="name">
          <NInput v-model:value="model.name" :placeholder="$t('page.dict.type.form.name')" />
        </NFormItem>
        <NFormItem :label="$t('page.dict.type.code')" path="code">
          <NInput v-model:value="model.code" :placeholder="$t('page.dict.type.form.code')" />
        </NFormItem>
        <NFormItem :label="$t('page.dict.type.desc')">
          <NInput v-model:value="model.description" type="textarea" :placeholder="$t('page.dict.type.form.desc')" />
        </NFormItem>
        <NFormItem :label="$t('page.dict.type.status')">
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
