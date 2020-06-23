import axios from 'axios'
import config from 'config'

export const userService = {
  login,
  loadUser,
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

async function loadUser() {
  return await axios.create({
    baseURL: config.apiUrl
  }).get(`/me`, {
    timeout: 4000
  }).then((response) => {
    console.log(response)
    return response.data
  })
}
