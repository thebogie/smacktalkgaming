<script context="module">
	export function preload({ params }, { user }) {
		if (!user) {
			this.redirect(302, `/login`);
		}
	}
</script>



<script>
	import DatePicker from '../_components/DatePicker.svelte';
	
</script>

<svelte:head>
	<title>New Event</title>
</svelte:head>

<div>
HERE
</div>