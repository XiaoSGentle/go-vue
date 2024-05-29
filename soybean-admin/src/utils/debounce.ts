// 自定义防抖函数，支持传入等待时间(wait)。
export const debounce = (fn: Function, wait: number) => {
  let timer: NodeJS.Timeout | string | number | undefined;
  return (...args: any) => {
    clearTimeout(timer);
    timer = setTimeout(async () => {
      await fn(...args);
    }, wait);
  };
};
