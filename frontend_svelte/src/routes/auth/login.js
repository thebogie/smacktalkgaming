import * as api from "api.js";

export function post(req, res) {
  const user = req.body;

  console.log("^^^^^^^^^^^^^^^^^^^^^^^^^ POST INSIDED OF AUTH/LOGIN");

  api
    .post("login", { email: user.email, password: user.password })
    .then((response) => {
      if (response.user) req.session.user = response.user;
      res.setHeader("Content-Type", "application/json");

      res.end(JSON.stringify(response));
    });
}
