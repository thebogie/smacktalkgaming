<script context="module">
  import * as api from "api";
  import axios from "axios";

  let inProgress = false;
  let error = null;

  export async function preload(page, session) {
    if (!session.user) {
      return this.redirect(302, "login");
    }

    console.log("VIEW SESSION:" + JSON.stringify(session.user));

    try {
      //let user = await api.users.getUser({ userId: params.id });
      inProgress = true;
      const response = await axios.post("auth/getuser", session.user);

      console.log("FUCKINGUSER:" + JSON.stringify(response));
      //userstats = await api.users.getUserStats({
      // userId: $session.user.userid,
      //daterange: "01011970"
      //});
      session.user = response.data;
      inProgress = false;
      error = null;
    } catch (e) {
      console.log("ERRROR: %s" + e);
      error = e.response.data.message;
      inProgress = false;
    }
  }
</script>

<svelte:head>
  <title>View Profile</title>
</svelte:head>
here
