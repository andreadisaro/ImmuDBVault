<template>
  <v-layout class="rounded rounded-md">
    <v-main class="d-flex align-center justify-center" style="min-height: 300px">
      <router-view />
    </v-main>
  </v-layout>
  <Loading />
</template>

<script lang="ts" setup>
import Loading from "@/components/Loading.vue";
</script>
