import { defineStore } from 'pinia';

import { ref } from 'vue';
import { SetupStoreId } from '@/enum';
import { fetchDict } from '@/service/api';

export const useDictStore = defineStore(SetupStoreId.Dict, () => {
  const thisStore = ref<Map<string, Api.Dict.Dict[]>>(new Map());
  function getDictData(dictCode: string) {
    if (thisStore.value.has(dictCode)) {
      const data = thisStore.value.get(dictCode);
      if (data !== undefined) return data;
      return [];
    }
    fetchDict(dictCode).then(res => {
      if (res.data) thisStore.value.set(dictCode, res.data);
      if (res.data) return res.data;
      return [];
    });

    window.$message?.error(`字典数据不存在`);
    return [];
  }
  return {
    getDictData
  };
});
