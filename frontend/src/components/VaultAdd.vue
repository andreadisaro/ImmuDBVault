<template>
  <v-sheet width="300" class="mx-auto">
    <v-form ref="form" v-model="formValidation">
      <v-text-field
        v-model="obj.accountNumber"
        :counter="7"
        :rules="accountNumberRules"
        label="Account Number"
        required
      ></v-text-field>
      <v-text-field
        v-model="obj.accountName"
        :counter="10"
        :rules="accountNameRules"
        label="Account Name"
        required
      ></v-text-field>
      <v-text-field
        v-model="obj.iban"
        :rules="ibanRules"
        label="IBAN"
        required
      ></v-text-field>
      <v-text-field v-model="obj.address" label="Address" required></v-text-field>
      <v-text-field
        v-model.number="obj.amount"
        :rules="amountRules"
        label="Amount"
        required
        type="number"
      ></v-text-field>

      <v-select v-model="obj.type" :items="types" label="Type" :rules="typeRules" required
        ><template v-slot:item="{ props, item }">
          <v-list-item v-bind="props"
            ><template v-slot:prepend>
              <v-icon v-if="item.title === 'sending'">mdi-upload</v-icon>
              <v-icon v-else>mdi-download</v-icon>
            </template></v-list-item
          >
        </template></v-select
      >

      <div class="d-flex flex-column">
        <v-btn type="submit" color="success" class="mt-4" block @click="putToVault">
          Send
        </v-btn>

        <v-btn color="error" class="mt-4" block @click="reset"> Reset Form </v-btn>
      </div>
    </v-form>
  </v-sheet>
  <v-dialog v-model="showSuccessDialog" width="auto">
    <v-card>
      <v-card-text> Record successfully added to vault </v-card-text>
      <v-card-actions>
        <v-btn color="primary" block @click="showSuccessDialog = false"
          >Close Dialog</v-btn
        >
      </v-card-actions>
    </v-card>
  </v-dialog>
  <v-dialog v-model="showErrorDialog" width="auto">
    <v-card>
      <v-card-text> {{ errorMessage }} </v-card-text>
      <v-card-actions>
        <v-btn color="primary" block @click="showErrorDialog = false">Close Dialog</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts" setup>
import { ref, nextTick } from "vue";
import { Account } from "@/types/account";
import { useLoadingStore } from "@/store/loading";
import { PutToVaultResponse } from "@/types/putToVaultResponse";
import { ErrorResponse } from "@/types/errorResponse";
const loadingStore = useLoadingStore();
const obj = ref<Account>({
  accountNumber: "",
  accountName: "",
  iban: "",
  address: "",
  amount: 0,
  type: "",
});
const showSuccessDialog = ref(false);
const showErrorDialog = ref(false);
const errorMessage = ref("");
const form = ref();
const formValidation = ref(false);
const types = ["sending", "receiving"];
const accountNameRules = [
  (v: string) => (v && v.length <= 10) || "Name must be less than 10 characters",
];
const accountNumberRules = [
  (v: string) => !!v || "Account Number is required",
  (v: string) => (v && v.length === 7) || "Account number must be 7 chars",
];
const ibanRules = [
  (v: string) => !!v || "IBAN is required",
  (v: string) => (v && validateIbanChecksum(v)) || "Invalid IBAN format",
];
const amountRules = [
  (v: number) => !!v || "Amount is required",
  (v: number) => (v && !isNaN(v)) || "Amount is not a number",
];
const typeRules = [(v: string) => !!v || "Type is required"];

function validateIbanChecksum(iban: string) {
  const ibanStripped = iban
    .replace(/[^A-Z0-9]+/gi, "") //keep numbers and letters only
    .toUpperCase(); //calculation expects upper-case
  const m = ibanStripped.match(/^([A-Z]{2})([0-9]{2})([A-Z0-9]{9,30})$/);
  if (!m) return false;

  const numbericed = (m[3] + m[1] + m[2]).replace(/[A-Z]/g, (ch) => {
    //replace upper-case characters by numbers 10 to 35
    return (ch.charCodeAt(0) - 55).toString();
  });
  //The resulting number would be too long for javascript to handle without loosing precision.
  //So the trick is to chop the string up in smaller parts.
  let mod97 = numbericed
    .match(/\d{1,7}/g)
    ?.reduce((total, curr) => (Number(total + curr) % 97).toString(), "");

  return mod97 === "1";
}

const reset = () => {
  form.value.reset();
};

const putToVault = (event: Event) => {
  event.preventDefault();
  if (!formValidation.value) {
    return;
  }
  loadingStore.addLoading();
  nextTick(() => {
    fetch("/be/putToVault", {
      method: "PUT",
      body: JSON.stringify(obj.value),
    })
      .then(function (response) {
        if (response.status != 200) {
          console.error(response.status);
          errorMessage.value = "Server error. Please try again later.";
          showErrorDialog.value = true;
        } else {
          response.json().then((res: PutToVaultResponse | ErrorResponse) => {
            console.log(res);
            if ("error" in res) {
              console.error(res.code, res.error);
              errorMessage.value = res.error;
              showErrorDialog.value = true;
            } else {
              showSuccessDialog.value = true;
              reset();
            }
          });
        }
      })
      .finally(() => {
        loadingStore.removeLoading();
      });
  });
};
</script>
