#!/usr/bin/env node
'use strict'

const home = require('os').homedir()
require('dotenv').config({ path: home + '/scripts/.env' })
const got = require('got')

const keys = {
  pinata_api_key: process.env.PINATA_API_KEY,
  pinata_secret_api_key: process.env.PINATA_SECRET_API_KEY
}

;(async () => {
  const { body } = await got.post('https://api.pinata.cloud/pinning/pinByHash', {
    headers: keys,
    responseType: 'json',
    json: {
      hashToPin: 'QmbA8Uw6LXEMtGiQr2X1RPRLT1hJMQnMHb9pd1HUnGftaK',
      pinataOptions: {
        hostNodes: [
          '/ip4/85.247.73.76/tcp/26264/ipfs/QmfULYxjC3qFW3k5vPmwFvjaeGz2CPosVya4DrBHz3xayp'
          // '/ip4/169.254.60.115/tcp/4001/ipfs/QmfULYxjC3qFW3k5vPmwFvjaeGz2CPosVya4DrBHz3xayp'
        ]
      }
    }
  })

  console.log(body)
})()
