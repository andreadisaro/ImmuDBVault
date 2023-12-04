// Utilities
import { defineStore } from "pinia";
import { ref } from "vue";

export const useAppStore = defineStore("app", () => {
  const vaultPage = ref("list");
  function setVaultPage(page: string) {
    vaultPage.value = page;
  }

  return {
    vaultPage,
    setVaultPage,
  };
});
