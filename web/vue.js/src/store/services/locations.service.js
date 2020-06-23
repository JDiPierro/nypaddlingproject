import axios from 'axios'
import config from 'config'

export const locationService = {
  load,
  claim,
  details
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

async function claim(location_id) {
  return await axios.create({
    baseURL: config.apiUrl
  }).post(`/locations/${location_id}/claim`, {}, {
    timeout: 10000
  }).then((response) => {
    return response.data
  })
}

async function details(location_id) {
  return await axios.create({
    baseURL: config.apiUrl
  }).get(`/locations/${location_id}`, {}, {
    timeout: 10000
  }).then((response) => {
    return response.data
  })
}
