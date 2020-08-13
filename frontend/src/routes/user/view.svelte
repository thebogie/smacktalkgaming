<script context="module">
  export function preload (page, session) {
    if (!session.user) {
      return this.redirect(302, 'login')
    }
  }
</script>

<svelte:head>
  <title>View Profile</title>
</svelte:head>
