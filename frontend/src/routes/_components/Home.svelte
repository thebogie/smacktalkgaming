<script>
  //import Stories from './Stories/index.svelte'
  import Contests from './Contests/index.svelte'

  export let currentPage = 1
</script>

<svelte:head>
  <title>Smack Talk Gaming</title>
</svelte:head>

<div class="homepage page">
  <Contests {currentPage}/>
</div>
