<script lang="ts">
  import { Button } from "$lib/components/ui/button";
  import * as Card from "$lib/components/ui/card";
  import { Input } from "$lib/components/ui/input";
  import { Label } from "$lib/components/ui/label";
  import { writable } from "svelte/store";
  import { Eye, EyeClosed, LoaderCircle } from "lucide-svelte";
  import { goto } from "$app/navigation";

  type ResponseLogin = {
    ok: boolean;
    token?: string;
    error?: string;
  };

  type Body = {
    username: string;
    password: string;
  };

  let showPassword = writable(false);
  const togglePasswordVisibility = () => {
    showPassword.update(value => !value);
  };

  let username = writable("");
  let password = writable("");
  let errorMessage = writable("");
  let isLoading = writable(false);

  const setCookie = (name: string, value: string, days: number) => {
    const date = new Date();
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000);
    const expires = `; expires=${date.toUTCString()}`;
    document.cookie = `${name}=${value}${expires}; path=/`;
  };

  let form: HTMLFormElement;

  const handleSubmit = async (): Promise<void> => {
    isLoading.set(true);

    if (!form.reportValidity()) {
      isLoading.set(false);
      return;
    }

    const usernameRegex = /^[A-Za-z]+$/;

    if (!usernameRegex.test($username)) {
      errorMessage.set("Username must contain only letters.");
      isLoading.set(false);
      return;
    }

    try {
      const body: Body = {
        username: $username,
        password: $password
      };

      const response = await fetch("/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(body)
      });

      const result: ResponseLogin = await response.json();

      if (!response.ok || !result.ok) {
        errorMessage.set("Invalid username or password");
        isLoading.set(false);
        return;
      }

      if (result.token) {
        setCookie("token", result.token, 7);
        setCookie("username", $username, 7);
        errorMessage.set("");
        goto("/");
      }
    } catch (err) {
      if (err instanceof Error) {
        console.error("Error during login:", err.message);
        errorMessage.set("Invalid username or password");
      } else {
        console.error("Unexpected error", err);
        errorMessage.set("An unexpected error occurred.");
      }
    } finally {
      await new Promise(resolve => setTimeout(resolve, 300));
      isLoading.set(false);
    }
  };
</script>

<div class="flex min-h-screen items-center justify-center">
  <Card.Root class="w-[400px] h-[450px]">
    <Card.Header>
      <Card.Title>Login</Card.Title>
      <Card.Description>Please enter your login credentials.</Card.Description>
    </Card.Header>
    <Card.Content>
      <form bind:this={form} on:submit|preventDefault={handleSubmit}>
        <div class="grid w-full items-center gap-4">
          <div class="flex flex-col space-y-1.5">
            <Label for="username">Username</Label>
            <Input
              id="username"
              type="text"
              required
              placeholder="Enter your username"
              bind:value={$username}
            />
          </div>

          <div class="flex flex-col space-y-1.5 relative">
            <Label for="password">Password</Label>
            <Input
              id="password"
              type={$showPassword ? "text" : "password"}
              required
              placeholder="Enter your password"
              bind:value={$password}
            />
            <button
              type="button"
              class="absolute right-3 top-[50%] transform -translate-y-[50%]"
              on:click={togglePasswordVisibility}
            >
              {#if $showPassword}
                <Eye />
              {:else}
                <EyeClosed />
              {/if}
            </button>
          </div>
        </div>
      </form>
    </Card.Content>

    {#if $errorMessage}
      <p class="text-red-500 text-sm flex items-center justify-center">
        {$errorMessage}
      </p>
    {/if}
    <Card.Footer class="mt-6 flex justify-center">
      <div class="flex flex-col w-full items-center">
        <Button
          on:click={handleSubmit}
          class="transition transform active:scale-95 active:bg-blue-700"
          disabled={$isLoading}
        >
          {#if $isLoading}
            <LoaderCircle class="animate-spin mr-2" />
            Loading...
          {:else}
            Login
          {/if}
        </Button>
      </div>
    </Card.Footer>
  </Card.Root>
</div>
