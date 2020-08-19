<script context="module">
  import * as api from "api";

  let error = null;

  export async function preload(page, session) {
    if (!session.user) {
      return this.redirect(302, "login");
    }

    console.log("VIEW SESSION:" + JSON.stringify(session.user));

    try {
      //let user = await api.users.getUser({ userId: params.id });

      const response = await api.users.getUser({ userId: session.user.userid });

      console.log("FUCKINGUSER:" + JSON.stringify(response));
      //userstats = await api.users.getUserStats({
      // userId: $session.user.userid,
      //daterange: "01011970"
      //});
    } catch (e) {
      console.log("ERRROR: %s" + e);
      error = e;
    }
  }
</script>

<svelte:head>
  <title>View Profile</title>
</svelte:head>
here
