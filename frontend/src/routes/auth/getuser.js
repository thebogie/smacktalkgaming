import * as api from 'api'

export async function post (req, res) {
  console.log("getuser get? %s",req.body.userid);
  console.log("getuser get? %s", res);

  let user = req.body.userid;
  try {
    
   
    const response = await api.users.getUser( { userId: req.body.userid} )
    req.session.user = response
    
    res.end(JSON.stringify(response))
  } catch (err) {
    res.statusCode = err.status
    res.end(JSON.stringify(err))
  }
}
