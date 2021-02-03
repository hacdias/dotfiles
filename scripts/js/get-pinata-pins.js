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
  const { body } = await got('https://api.pinata.cloud/data/pinList?status=pinned&pageLimit=1000', {
    headers: keys,
    responseType: 'json'
  })

  const items = body.rows
    .filter(r => !r.date_unpinned)
    .filter(r => !r.metadata.name.startsWith('_dnslink.'))
    .map(r => ({
      tag: r.metadata.name,
      cid: r.ipfs_pin_hash
    }))

  console.log(JSON.stringify(items, null, 2))
})()
