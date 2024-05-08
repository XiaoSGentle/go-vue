<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { useDictStore } from '@/store/modules/dict';
const { getDictData } = useDictStore();
defineOptions({
  name: 'SysDict'
});

interface Props {
  /*
   * 字典key */
  dictKey: string;
  type?: 'check' | 'radio' | 'radio-button' | 'select' | 'select-multiple';
}
const props = withDefaults(defineProps<Props>(), {
  type: 'select'
});

// 要双向绑定的值
const selectValue = defineModel<string[] | string>('selectValue', { required: true });
// store中获取的Label列表
const options = ref<Array<Api.Dict.Dict>>([]);

// 自定义组件中的值
const multipleValue = computed<string[]>({
  get: () => {
    if (!(typeof selectValue.value === 'string')) return selectValue.value;
    return [];
  },
  set: value => {
    return (selectValue.value = value);
  }
});
const singleValue = computed<string>({
  get: () => {
    if (typeof selectValue.value === 'string') return selectValue.value;
    return '';
  },
  set: value => {
    return (selectValue.value = value);
  }
});

// 定义获取字典数据的方法
const fetchDictData = async () => {
  const data = await getDictData(props.dictKey);
  options.value = data;
};

onMounted(() => {
  fetchDictData();
});
</script>

<template>
  <div>
    <NCheckboxGroup v-if="type === 'check'" v-model:value="multipleValue">
      <NCheckbox v-for="item in options" :key="item.value" :value="item.value" :label="item.label" />
    </NCheckboxGroup>
    <NRadioGroup v-if="type === 'radio' || type === 'radio-button'" v-model:value="singleValue">
      <div v-if="type === 'radio'">
        <NRadio v-for="item in options" :key="item.value" :value="item.value" :label="item.label" />
      </div>
      <div v-if="type === 'radio-button'">
        <NRadioButton v-for="item in options" :key="item.value" :value="item.value" :label="item.label" />
      </div>
    </NRadioGroup>
    <NSelect v-if="type === 'select'" v-model:value="singleValue" :options="options" />
    <NSelect
      v-if="type === 'select-multiple'"
      v-model:value="multipleValue"
      :multiple="type === 'select-multiple'"
      :options="options"
    />
  </div>
</template>

<style scoped></style>
