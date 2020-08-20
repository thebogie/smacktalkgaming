import * as api from 'api'

export async function post (req, res) {
  console.log("auth/getuser: %s",req.body);
  //console.log("getuser get? %s", res);


  try {
    
   
    const response = await api.users.getUser( req.body )
    req.session.user = response
    
    res.end(JSON.stringify(response))
  } catch (err) {
    res.statusCode = err.status
    res.end(JSON.stringify(err))
  }
}
