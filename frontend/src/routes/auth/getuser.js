import * as api from 'api'

export async function post (req, res) {
  console.log("getuser get? %s",req.body.user);
  //console.log("getuser get? %s", res);


  try {
    
   
    const response = await api.users.getUser( { userId: req.body._id} )
    req.session.user = response
    
    res.end(JSON.stringify(response))
  } catch (err) {
    res.statusCode = err.status
    res.end(JSON.stringify(err))
  }
}
