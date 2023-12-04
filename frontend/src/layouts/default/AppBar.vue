<template>
  <v-app-bar rounded :height="100" style="padding: 0 0 0 80px">
    <template v-slot:prepend>
      <CodenotaryLogo />
    </template>
    <template v-slot:append>
      <v-btn text="Home" to="/home"></v-btn>

      <v-btn text="Vault" to="/vault"></v-btn>
    </template>
  </v-app-bar>
</template>

<script lang="ts" setup>
import CodenotaryLogo from "@/assets/CodenotaryLogo.vue";
</script>
