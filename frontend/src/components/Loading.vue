<script lang="ts" setup>
import { ref, watch } from "vue";
import { useLoadingStore } from "../store/loading";
const loadingStore = useLoadingStore();
const showOverlay = ref(false);
watch(
  () => loadingStore.count,
  (count) => {
    if (count > 0) {
      showOverlay.value = true;
    } else {
      showOverlay.value = false;
    }
  }
);
</script>

<template>
  <v-overlay v-model="showOverlay" class="loading">
    <v-progress-circular
      :size="70"
      :width="7"
      color="purple"
      indeterminate
    ></v-progress-circular
    ><span style="padding-left: 1em">Loading...</span></v-overlay
  >
</template>

<style scoped>
.loading {
  background: rgba(85, 85, 85, 0.65);
  display: flex;
  position: fixed;
  inset: 0px;
  z-index: 9999;
  justify-content: center;
  align-items: center;
}
</style>
