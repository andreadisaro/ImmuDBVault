import { ref } from "vue";
import { defineStore } from "pinia";

export const useLoadingStore = defineStore("loading", () => {
  const count = ref(0);
  function addLoading() {
    count.value++;
  }
  function removeLoading() {
    count.value = count.value > 0 ? count.value - 1 : 0;
  }

  return { count, addLoading, removeLoading };
});
