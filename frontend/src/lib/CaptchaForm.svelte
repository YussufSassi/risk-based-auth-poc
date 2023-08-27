<script lang="ts">
  import { getCurrentBrowserFingerPrint } from "@rajesh896/broprint.js";
  import { captchaMode, username } from "../store";
  import type { IResponse } from "../intefaces/Response";

  let captcha: string;
  let fingerPrintToken: string;
  let response: IResponse;

  getCurrentBrowserFingerPrint().then((value) => (fingerPrintToken = value));

  async function verifyCaptcha() {
    const form = new FormData();
    form.append("user", $username);
    form.append("token", fingerPrintToken);
    form.append("code", captcha);

    const res = await fetch("/captcha", {
      method: "POST",
      body: form,
    });

    response = await res.json();

    if (res.status == 200) {
      captchaMode.set(false);
    }
  }
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
  <div class="sm:mx-auto sm:w-full sm:max-w-sm">
    <h2
      class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900"
    >
      Enter the following captcha below
    </h2>
    <img src="/captcha.png" alt="captcha" class="bg-black" />
  </div>

  <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
    <form
      class="space-y-6"
      on:submit|preventDefault={verifyCaptcha}
      method="POST"
    >
      <div>
        <div class="flex items-center justify-between">
          <label
            for="captcha"
            class="block text-sm font-medium leading-6 text-gray-900"
            >Captcha</label
          >
        </div>
        <div class="mt-2">
          <input
            id="captcha"
            name="captcha"
            bind:value={captcha}
            type="text"
            required
            class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
          />
        </div>
      </div>

      <div>
        <button
          type="submit"
          class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >Verify captcha</button
        >
      </div>
    </form>
  </div>
  {#if response}
    <p class="text-center mt-9 font-bold">{response.message}</p>
  {/if}
</div>
