<template>
  <div style="display: flex; flex-direction: column">
    <div style="display: flex; flex-direction: row; font-weight: 600" class="px-2">
      <div style="width: 15%">Account Number</div>
      <div style="width: 15%">Account Name</div>
      <div style="width: 25%">IBAN</div>
      <div style="width: 25%">Address</div>
      <div style="width: 15%">Amount</div>
      <div style="width: 5%">Type</div>
    </div>
    <v-infinite-scroll height="300" mode="manual" @load="getAllFromVault">
      <template v-for="(item, index) in items" :key="item">
        <div :class="['px-2 py-2', index % 2 === 0 ? 'bg-grey-lighten-2' : '']">
          <div style="display: flex; flex-direction: row">
            <div
              style="
                width: 15%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              {{ item.accountNumber }}
            </div>
            <div
              style="
                width: 15%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              {{ item.accountName }}
            </div>
            <div
              style="
                width: 25%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              {{ item.iban }}
            </div>
            <div
              style="
                width: 25%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              {{ item.address }}
            </div>
            <div
              style="
                width: 15%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              {{ item.amount }}
            </div>
            <div
              style="
                width: 5%;
                overflow: hidden;
                white-space: nowrap;
                text-overflow: ellipsis;
              "
            >
              <v-icon v-if="item.type === 'sending'">mdi-upload</v-icon>
              <v-icon v-else>mdi-download</v-icon>
            </div>
          </div>
        </div>
      </template>
      <template v-slot:load-more="{ props }">
        <v-btn color="primary" variant="flat" v-bind="props">
          <v-icon icon="mdi-refresh" size="large" start />

          LOAD MORE
        </v-btn>
      </template>
      <template v-slot:empty>
        <v-alert type="warning" style="margin: 0 40%">No more items!</v-alert>
      </template>
    </v-infinite-scroll>
  </div>
</template>

<script lang="ts" setup>
import { ref, nextTick } from "vue";
import { PageRequest } from "../types/pageRequest";
import { PageResponse, Document } from "@/types/pageResponse";
import { Account } from "@/types/account";
const pager = ref<PageRequest>({
  page: 0,
  perPage: 5,
});
const items = ref<Document<Account>[]>([]);

const getAllFromVault = ({ side, done }: any) => {
  done("loading");
  pager.value = { ...pager.value, page: pager.value.page + 1 };
  nextTick(() => {
    fetch("/be/getAllFromVault", {
      method: "POST",
      body: JSON.stringify(pager.value),
    }).then(function (response) {
      if (response.status != 200) {
        console.error(response.status);
        done("error");
      } else {
        response.json().then((res: PageResponse<Account>) => {
          console.log(res);
          if (res.revisions.length === 0) done("empty");
          else {
            if (side === "end")
              items.value = [...items.value, ...res.revisions.map((r) => r.document)];
            else items.value = [...res.revisions.map((r) => r.document), ...items.value];
            done("ok");
          }
        });
      }
    });
  });
};
getAllFromVault({ side: "end", done: () => {} });
</script>
