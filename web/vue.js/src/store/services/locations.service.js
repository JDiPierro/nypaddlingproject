import axios from 'axios'
import config from 'config'

export const locationService = {
  load
};

async function load () {
  return await axios.create({
    baseURL: config.apiUrl
  }).get(`/locations`, {
    timeout: 10000
  }).then((response) => {
    console.log("API Response:")
    console.log(response.data)
    return response.data
  })
}
