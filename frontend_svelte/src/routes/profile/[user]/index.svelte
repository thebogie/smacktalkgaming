<script context="module">
  import { post } from "utils.js";
  export async function preload({ params }, { user }) {
    var datetouse = "01011970";

    if (user == undefined) {
      this.redirect(302, `/login`);
    }

    console.log("%%%%%%%%%%%%%%%% inside preload with " + JSON.stringify(user));

    //const profile = await get(
    //  `auth/users/, ${user._id}/stats/01011970`, user.token
    //);

    const profile = await post(`auth/user`, { user, datetouse });

    console.log("%%%%%%%%%%%%%%%% " + JSON.stringify(profile));

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
