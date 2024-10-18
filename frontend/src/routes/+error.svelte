<script context="module" lang="ts">
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  const getCookie = (name: string): string | null => {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop()?.split(";").shift() || null;
    return null;
  };

  onMount(() => {
    const token = getCookie("token");

    if (!token) {
      setTimeout(() => {
        goto("/auth/login");
      }, 100);
    }
  });
</script>

<div class="min-h-screen flex items-center justify-center">
  <p class="text-2xl font-bold">Page Not Found</p>
</div>

<style>
  .min-h-screen {
    height: 100vh;
  }
</style>
