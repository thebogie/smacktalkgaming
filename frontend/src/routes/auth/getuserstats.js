import * as api from 'api'

export async function post (req, res) {
  console.log("*****getuserstats: %v",req.body);
  //console.log("*****getuserstats: %v", res);


  try {
    
   
    const response = await api.users.getUserStats( 
      { userId: req.body.user._id, 
        daterange: req.body.daterange} )
    req.session.userstats = response
    
    res.end(JSON.stringify(response))
  } catch (err) {
    res.statusCode = err.status
    res.end(JSON.stringify(err))
  }
}
