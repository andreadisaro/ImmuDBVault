<template>
  <v-navigation-drawer>
    <v-list-item title="Vault App" subtitle="by Codenotary"></v-list-item>
    <v-divider></v-divider>
    <v-list-item></v-list-item>
    <v-list-item
      link
      title="List items"
      :active="appStore.vaultPage === 'list'"
      @click="appStore.setVaultPage('list')"
    ></v-list-item>
    <v-list-item
      link
      title="Add item"
      :active="appStore.vaultPage === 'add'"
      @click="appStore.setVaultPage('add')"
    ></v-list-item>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>
import { useAppStore } from "@/store/app";
const appStore = useAppStore();
</script>
