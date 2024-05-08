<script lang="tsx" setup>
import { NCard } from 'naive-ui';
import { ref } from 'vue';
const selectValue = ref('2');
const selectValueList = ref(['1', '2']);
</script>

<template>
  <NCard>
    <SysDict v-model:select-value="selectValue" dict-key="DICT_LABEL" />
    <div>{{ selectValue }}</div>
    <SysDict v-model:select-value="selectValue" dict-key="DICT_LABEL" type="radio-button" />
    <div>{{ selectValue }}</div>
    <SysDict v-model:select-value="selectValueList" dict-key="DICT_LABEL" type="check" />
    <div>{{ selectValueList }}</div>
    <SysDict v-model:select-value="selectValue" dict-key="DICT_LABEL" type="radio" />
    <div>{{ selectValue }}</div>
    <SysDict v-model:select-value="selectValueList" dict-key="DICT_LABEL" type="select-multiple" />
    <div>{{ selectValueList }}</div>
  </NCard>
</template>

<style scoped></style>
