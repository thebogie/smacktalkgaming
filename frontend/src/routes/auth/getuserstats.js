import * as api from 'api'

export async function post (req, res) {
  console.log("auth/getuserstats: %v",req.body); 

  try {
    
   
    const response = await api.users.getUserStats( req.body )
    req.session.userstats = response
    
    res.end(JSON.stringify(response))
  } catch (err) {
    res.statusCode = err.status
    res.end(JSON.stringify(err))
  }
}
