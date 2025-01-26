import { createRouter, createWebHistory } from "vue-router";
import UserLogin from "@/views/UserLogin.vue";
import UserRegister from "@/views/UserRegister.vue";
import JournalEntry from "@/views/JournalEntry.vue";
import UserHome from "@/views/UserHome.vue";

const routes = [
  { path: "/", component: UserHome },
  { path: "/login", component: UserLogin },
  { path: "/register", component: UserRegister },
  { path: "/journal", component: JournalEntry },
  { path: "/:pathMatch(.*)*", redirect: "/" }, // New catch-all syntax for Vue Router 4
];

const router = createRouter({
  history: createWebHistory(), // Use the history mode
  routes,
});

export default router;
