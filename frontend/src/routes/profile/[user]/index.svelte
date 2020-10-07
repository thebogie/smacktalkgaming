<script context="module">
  import * as api from "api.js";

  export async function preload({ params }, { user }) {


    if (user == undefined) {
      this.redirect(302, `/login`);
    }

        console.log("+++++++++++" + JSON.stringify(user));


    //const username = params.user.slice(1);

    const profile = await api.get(
      `users/${user._id}/stats/01011970`,
      user && user.token
    );

    console.log("+++++++++++" + JSON.stringify(profile));

    return { profile, favorites: params.view === "favorites" };
  }
</script>

<script>
  import { stores } from "@sapper/app";
  import Profile from "./_Profile.svelte";

  export let profile;
  export let favorites;

  const { session } = stores();
</script>

<svelte:head>
  <title>Smacktalk Profile</title>
</svelte:head>

<Profile {profile} {favorites} user={$session.user} />
