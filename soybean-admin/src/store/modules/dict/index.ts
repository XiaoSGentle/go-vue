import { defineStore } from 'pinia';

import { ref } from 'vue';
import { SetupStoreId } from '@/enum';
import { fetchDict } from '@/service/api';

export const useDictStore = defineStore(SetupStoreId.Dict, () => {
  const thisStore = ref<Map<string, Api.Dict.Dict[]>>(new Map());
  async function getDictData(dictCode: string) {
    if (thisStore.value.has(dictCode)) {
      const data = thisStore.value.get(dictCode);
      if (data !== undefined) return data;
      return [];
    }
    const { error, data } = await fetchDict(dictCode);
    if (error) {
      window.$message?.error(`字典数据不存在`);
    }
    if (!data) {
      return [];
    }
    thisStore.value.set(dictCode, data);
    return data;
  }
  return {
    getDictData
  };
});
