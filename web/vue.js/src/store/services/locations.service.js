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
    return response.data
  })
}
