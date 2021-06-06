#!/usr/bin/env node

import got from 'got'

const apiUrl = 'https://api.pinata.cloud/data/pinList?status=pinned&pageLimit=1000'

const headers = {
  pinata_api_key: process.env.PINATA_API_KEY,
  pinata_secret_api_key: process.env.PINATA_SECRET_API_KEY
}

const { body } = await got(apiUrl, { headers, responseType: 'json' })

const items = body.rows
  .filter(r => !r.date_unpinned)
  .filter(r => !r.metadata.name.startsWith('_dnslink.'))
  .map(r => ({
    tag: r.metadata.name,
    cid: r.ipfs_pin_hash
  }))


console.log(JSON.stringify(items, null, 2))
