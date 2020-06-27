import axios from 'axios'
import config from 'config'

export const locationService = {
  load,
  claim,
  release,
  submit,
  loadClaims,
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

async function release(location_id) {
  return await axios.create({
    baseURL: config.apiUrl
  }).post(`/locations/${location_id}/claim/release`, {}, {
    timeout: 10000
  }).then((response) => {
    return response.data
  })
}

async function submit(location_id, update_info) {
  console.log(`Requesting submission of ${location_id} with info:`, update_info)
  return await axios.create({
    baseURL: config.apiUrl
  }).post(`/locations/${location_id}/claim/submit`, update_info, {
    timeout: 10000
  }).then((response) => {
    return response.data
  })
}


async function loadClaims () {
  return await axios.create({
    baseURL: config.apiUrl
  }).get(`/me/claims`, {
    timeout: 4000
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
