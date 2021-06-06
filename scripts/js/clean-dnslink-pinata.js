#!/usr/bin/env node

import got from 'got'

const headers = {
  pinata_api_key: process.env.PINATA_API_KEY,
  pinata_secret_api_key: process.env.PINATA_SECRET_API_KEY
}

console.log('📄 Fetching all the pins...')

const { body } = await got('https://api.pinata.cloud/data/pinList?status=pinned&pageLimit=1000', {
  headers,
  responseType: 'json'
})

const items = body.rows
  .filter(r => !r.date_unpinned)
  .filter(r => r.metadata.name.startsWith('_dnslink.'))
  .sort((a, b) => new Date(a.date_pinned) - new Date(b.date_pinned))

const names = items.reduce((curr, item) => {
  if (!curr.includes(item.metadata.name)) {
    curr.push(item.metadata.name)
  }

  return curr
}, [])

console.log('💼 Pins fetched and parsed!')

for (const name of names) {
  const filtered = items.filter(r => r.metadata.name === name)
  filtered.pop()

  if (filtered.length < 1) {
    continue
  }

  for (const { ipfs_pin_hash: hash } of filtered) {
    console.log(`✅ Unpinning ${hash} for ${name}`)
    await got.delete(`https://api.pinata.cloud/pinning/unpin/${hash}`, { headers })
  }
}

