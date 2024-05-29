import { defineStore } from 'pinia';

import { ref } from 'vue';
import { SetupStoreId } from '@/enum';
import { fetchDict } from '@/service/api';
export const useDictStore = defineStore(SetupStoreId.Dict, () => {
  const thisStore = ref<Map<string, Api.Dict.Dict[]>>(new Map());
  const promiseCache = ref<Map<string, Promise<Api.Dict.Dict[]>>>(new Map());

  async function getDictData(dictCode: string) {
    // 如果已经缓存了该字典数据，直接返回
    if (thisStore.value.has(dictCode)) {
      return thisStore.value.get(dictCode) || [];
    }

    // 如果已经有相同字典代码的请求正在执行，则返回该请求的 Promise
    if (promiseCache.value.has(dictCode)) {
      return promiseCache.value.get(dictCode) || [];
    }

    // 创建一个新的 Promise，并将其缓存到 promiseCache 中
    const promise = fetchDict(dictCode)
      .then(({ data }) => {
        if (!data) {
          return [];
        }
        // 请求成功，将数据缓存到 thisStore 中
        thisStore.value.set(dictCode, data);
        return data;
      })
      .catch(error => {
        // 请求失败，打印错误信息
        window.$message?.error(`获取字典数据失败: ${error}`);
        return [];
      })
      .finally(() => {
        // 无论成功或失败，都从 promiseCache 中移除该 Promise
        promiseCache.value.delete(dictCode);
      });

    // 将新的 Promise 缓存到 promiseCache 中
    promiseCache.value.set(dictCode, promise);

    // 返回新的 Promise
    return promise;
  }

  return {
    getDictData
  };
});
