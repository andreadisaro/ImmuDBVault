<template>
  <v-container class="fill-height">
    <v-responsive class="align-center text-center fill-height">
      <template v-if="appStore.vaultPage === 'list'">
        <VaultList />
      </template>
      <template v-else-if="appStore.vaultPage === 'add'"><VaultAdd /> </template>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>
import { useAppStore } from "@/store/app";
import VaultList from "@/components/VaultList.vue";
import VaultAdd from "@/components/VaultAdd.vue";
const appStore = useAppStore();
</script>
