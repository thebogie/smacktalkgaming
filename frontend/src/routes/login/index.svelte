<script>
  import { goto, stores } from '@sapper/app'
  import axios from 'axios'

  const { session } = stores()

  let user = { email: '', password: '' }
  let inProgress = false
  let error = null

  async function submit () {
    try {
      inProgress = true
          console.log("HERE:" )
      const response = await axios.post('/auth/login', user)
 
      $session.user = response.data
      inProgress = false
      error = null
      user = { email: '', password: '' }
      goto('/')
    } catch (err) {
      error = err.response.data.message
      inProgress = false
    }
  }
</script>

<svelte:head>
  <title>Login</title>
</svelte:head>

<form
  class="login-form"
  on:submit|preventDefault="{submit}"
>
  {#if error}
    <span class="error-message">
      {error}
    </span>
  {/if}
  <input
    class="text-input email-input"
    bind:value="{user.email}"
    type="text"
    placeholder="email"
    required
  >
  <input
    class="text-input password-input"
    bind:value="{user.password}"
    type="password"
    placeholder="password"
    required
  >
  <button
    class="login-button primary-button"
    disabled="{inProgress}"
  >
    LOG IN
  </button>
</form>
