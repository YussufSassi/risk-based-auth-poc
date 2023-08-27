<script lang="ts">
  import CaptchaForm from "./lib/CaptchaForm.svelte";
  import LoginForm from "./lib/LoginForm.svelte";
  import { captchaMode } from "./store";
</script>

{#if $captchaMode}
  <CaptchaForm />
{:else}
  <LoginForm />
{/if}
