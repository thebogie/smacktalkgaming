//export function get(req, res) {
//	res.setHeader('Content-Type', 'application/json');

//	res.end(JSON.stringify({ user: req.session.user || null }));
//}

import * as api from "api.js";

export function post(req, res) {
  const user = req.body.user;
  const datetouse = req.body.datetouse;
  console.log(
    "^^^^^^^^^^^^^^^^^^^^^^^^^ POST INSIDED OF AUTH/USER with REQ",
    user
  );
  console.log(
    "^^^^^^^^^^^^^^^^^^^^^^^^^ POST INSIDED OF AUTH/USER with res",
    datetouse
  );

  api.get("users/${user._id}/stats/01011970", user.token).then((response) => {
    if (response.user) req.session.user = response.user;
    res.setHeader("Content-Type", "application/json");

    res.end(JSON.stringify(response));
  });
}
