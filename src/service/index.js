import axios from 'axios'


const baseUrl = 'http://localhost:8080'
export const getDataUser = async function () {
   
    let user = await axios.get(`/users`)
    console.log("data user", user)
    return user
}
