import axios from 'axios'
import config from 'config'

export const userService = {
  login
};

async function login () {
  return await axios.create({
    baseURL: config.apiUrl
  }).get(`/login`, {
    timeout: 4000
  }).then((response) => {
    return response.data
  })
}
