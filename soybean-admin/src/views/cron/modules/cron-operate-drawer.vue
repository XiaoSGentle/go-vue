<script lang="ts" setup>
import { computed, reactive, watch } from 'vue';

import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { $t } from '@/locales';
import { enableStatusOptions } from '@/constants/business';
import { updateCronDataById } from '@/service/api';
import {REG_CRON} from "@/constants/reg";


defineOptions({
  name: 'CronOperateDrawer'
});

interface Props {
  /** the type of operation */
  operateType: NaiveUI.TableOperateType;
  /** the edit row data */
  rowData?: Api.Cron.CronType | null;
}

const props = defineProps<Props>();

interface Emits {
  (e: 'submitted'): void;
}

const emit = defineEmits<Emits>();

const visible = defineModel<boolean>('visible', {
  default: false
});
const { defaultRequiredRule } = useFormRules();
const { formRef, validate, restoreValidation } = useNaiveForm();

const title = computed(() => {
  const titles: Record<NaiveUI.TableOperateType, string> = {
    add: $t('page.manage.role.addRole'),
    edit: $t('page.manage.role.editRole')
  };
  return titles[props.operateType];
});

type Model = Api.Cron.AddOrUpdateCronTypeParams;

const model: Model = reactive(createDefaultModel());

function createDefaultModel(): Model {
  return {
    key:'',
    arguments:[],
    description:"",
    schedule:"",
    status:'1'
  };
}

type RuleKey =Pick<Model, 'description'|'schedule'|'arguments'|'status'>;

const rules: Record<keyof RuleKey, App.Global.FormRule> = {
  description: defaultRequiredRule,
  schedule: {
    pattern:REG_CRON,
    required: true,
    trigger:"change",
    message:$t('page.cron.form.scheduleWarning'),
  },
  arguments: defaultRequiredRule,
  status:defaultRequiredRule
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
    const { error } = await updateCronDataById(model);
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
  <NDrawer v-model:show="visible" :width="360" display-directive="show">
    <NDrawerContent :native-scrollbar="false" :title="title" closable>
      <NForm ref="formRef" :model="model" :rules="rules">
        <NFormItem :label="$t('page.cron.key')" >
          <NInput v-model:value="model.key" :placeholder="$t('page.cron.form.key')" disabled />
        </NFormItem>
        <NFormItem :label="$t('page.cron.schedule')" path="schedule">
          <NInput v-model:value="model.schedule" :placeholder="$t('page.cron.form.schedule')" />
        </NFormItem>
        <NFormItem :label="$t('page.cron.description')" path="description">
          <NInput v-model:value="model.description" :placeholder="$t('page.cron.form.description')" />
        </NFormItem>
        <NFormItem :label="$t('page.cron.arguments')" path="roleDesc">
          <NInput v-for="(item,index) in model.arguments" v-model:value="model.arguments[index]" :placeholder="$t('page.cron.form.arguments')" />
        </NFormItem>
        <NFormItem :label="$t('page.cron.status')" path="status">
          <NRadioGroup v-model:value="model.status">
            <NRadio v-for="item in enableStatusOptions" :key="item.value" :label="$t(item.label)" :value="item.value" />
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
