// Composables
import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import Vault from "@/views/Vault.vue";

const routes = [
  {
    path: "/",
    component: () => import("@/layouts/default/Default.vue"),
    children: [
      {
        path: "",
        name: "",
        component: Home,
      },
      {
        path: "home",
        name: "Home",
        component: Home,
      },
      {
        path: "vault",
        name: "Vault",
        component: Vault,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
